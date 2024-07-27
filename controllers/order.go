package controllers

import (
	"log"
	"math/rand"
	"time"
	"errors"
	"fmt"
	"github.com/InfamousFreak/Tech-Task-24/models"
	"github.com/InfamousFreak/Tech-Task-24/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PlaceOrder(c *fiber.Ctx) error {
    var orderRequest struct {
        UserID uint `json:"user_id"`
		PaymentToken string `json:"payment_token"`
    }

    if err := c.BodyParser(&orderRequest); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    // Start a transaction
    tx := database.Db.Begin()

    // Fetch cart items
    var cartItems []models.CartItem
    if err := tx.Where("user_id = ?", orderRequest.UserID).Find(&cartItems).Error; err != nil {
        tx.Rollback()
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch cart items"})
    }

    if len(cartItems) == 0 {
        tx.Rollback()
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cart is empty"})
    }

    // Calculate total amount
    var totalAmount float64
    for _, item := range cartItems {
        var menuItem models.MenuItem
        if err := tx.Where("menu_item_id = ?", item.MenuItemID).First(&menuItem).Error; err != nil {
            tx.Rollback()
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch menu item"})
        }
        totalAmount += menuItem.Price * float64(item.Quantity)
    }

    // Process payment (this is a mock function, replace with actual payment gateway)
    paymentResult, err := processPayment(totalAmount, orderRequest.PaymentToken)
    if err != nil {
        tx.Rollback()
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Payment failed: " + err.Error()})
    }

    // Create order
    order := models.Order{
        UserID: orderRequest.UserID,
        Status: "Placed",
        TotalAmount: totalAmount,
		TransactionID: paymentResult.TransactionID,
    }

    if err := tx.Create(&order).Error; err != nil {
        tx.Rollback()
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create order"})
    }

    // Create order items
    for _, item := range cartItems {
        orderItem := models.OrderItem{
            OrderID:    order.ID,
            MenuItemID: item.MenuItemID,
            Quantity:   item.Quantity,
        }
        if err := tx.Create(&orderItem).Error; err != nil {
            tx.Rollback()
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create order items"})
        }
    }

    // Clear cart
    if err := tx.Where("user_id = ?", orderRequest.UserID).Delete(&models.CartItem{}).Error; err != nil {
        tx.Rollback()
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to clear cart"})
    }

    // Commit transaction
    if err := tx.Commit().Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to commit transaction"})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "success": true, 
        "message": "Order placed successfully", 
        "order_id": order.ID,
        "total_amount": order.TotalAmount,
    })
}

func GetOrderHistory(c *fiber.Ctx) error {
    userID := c.Params("userID")

    var orders []models.Order
    result := database.Db.Preload("OrderItems").Where("user_id = ?", userID).Find(&orders)
    
    if result.Error != nil {
        log.Printf("Error fetching order history for user %s: %v", userID, result.Error)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to fetch order history",
            "details": result.Error.Error(),
        })
    }

    log.Printf("Successfully fetched %d orders for user %s", len(orders), userID)

    if len(orders) == 0 {
        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "message": "No orders found for this user",
            "orders": []models.Order{},
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Order history fetched successfully",
        "orders": orders,
    })
}

func CancelOrder(c *fiber.Ctx) error {
    var cancelRequest struct {
        UserID  uint `json:"user_id"`
        OrderID uint `json:"order_id"`
    }

    if err := c.BodyParser(&cancelRequest); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

        // Start a transaction
        tx := database.Db.Begin()

    // Fetch the order
    var order models.Order
    if err := tx.Where("id = ? AND user_id = ?", cancelRequest.OrderID, cancelRequest.UserID).First(&order).Error; err != nil {
        tx.Rollback()
        if err == gorm.ErrRecordNotFound {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Order not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch order"})
    }

    // Check if the order can be cancelled (e.g., only if it's in "Placed" status)
    if order.Status != "Placed" {
        tx.Rollback()
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Order cannot be cancelled in its current status"})
    }

    // Update order status to "Cancelled"
    order.Status = "Cancelled"
    if err := tx.Save(&order).Error; err != nil {
        tx.Rollback()
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update order status"})
    }

    // Optionally: Restore items to cart
    var orderItems []models.OrderItem
    if err := tx.Where("order_id = ?", order.ID).Find(&orderItems).Error; err != nil {
        tx.Rollback()
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch order items"})
    }

    for _, item := range orderItems {
        cartItem := models.CartItem{
            UserID:     order.UserID,
            MenuItemID: item.MenuItemID,
            Quantity:   item.Quantity,
        }
        if err := tx.Create(&cartItem).Error; err != nil {
            tx.Rollback()
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to restore items to cart"})
        }
    }

    // Commit transaction
    if err := tx.Commit().Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to commit transaction"})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "success": true,
        "message": "Order cancelled successfully",
        "order_id": order.ID,
    })
}


func processPayment(amount float64, token string) (*PaymentResult, error) {
    // Simulate payment processing
    time.Sleep(time.Second * 2)
    
    // Simulate success/failure (90% success rate)
    if rand.Float32() < 0.9 {
        return &PaymentResult{
            Success:       true,
            TransactionID: fmt.Sprintf("TXN%d", time.Now().UnixNano()),
        }, nil
    }
    
    return nil, errors.New("payment failed")
}

type PaymentResult struct {
    Success       bool
    TransactionID string
}
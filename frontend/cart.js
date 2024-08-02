document.addEventListener('DOMContentLoaded', function() {
    fetchCartItems();
});

function fetchCartItems() {
    const userId = localStorage.getItem('userId');

    if (!userId) {
        console.error('User ID not found. Please log in.');
        return;
    }

    fetch(`https://tech-task-24-latest.onrender.com/cart/${userId}`)
        .then(response => response.json())
        .then(data => {
            displayCartItems(data);
        })
        .catch(error => {
            console.error('Error fetching cart items:', error);
        });
}

function displayCartItems(cartItems) {
    const cartItemsContainer = document.getElementById('cart-items');
    cartItemsContainer.innerHTML = ''; // Clear previous items
    let totalAmount = 0;

    cartItems.forEach(item => {
        const itemName = item.item_name || 'Unknown Item';
        const itemPrice = item.item_price !== undefined ? item.item_price : 0;
        const quantity = item.quantity !== undefined ? item.quantity : 1;
        const itemId = item.item_id;

        const itemDiv = document.createElement('div');
        itemDiv.className = 'cart-item';
        itemDiv.innerHTML = `
            <h3>${itemName}</h3>
            <p>Price: $${itemPrice.toFixed(2)}</p>
            <p>Quantity: ${quantity}</p>
            <p>Subtotal: $${(itemPrice * quantity).toFixed(2)}</p>
            <div class="delete-btn">
            <button class="delete-button" data-item-id="${itemId}">X</button>
            </div>
        `;

        const deleteBtn = itemDiv.querySelector('.delete-button');
        deleteBtn.addEventListener('click', () => deleteCartItem(itemId));
        cartItemsContainer.appendChild(itemDiv);

        totalAmount += itemPrice * quantity;
    });

    document.getElementById('cart-total-amount').textContent = `$${totalAmount.toFixed(2)}`;
}

function deleteCartItem(itemId) {
    const userId = localStorage.getItem('userId');
    if (!userId) {
        console.error('User ID not found. Please log in.');
        return;
    }

    fetch('https://tech-task-24-latest.onrender.com/cart/delete', {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            user_id: parseInt(userId),
            item_id: parseInt(itemId)
        })
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to delete cart item');
        }
        return response.json();
    })
    .then(data => {
        console.log('Item deleted successfully:', data);
        fetchCartItems(); // Refresh the cart after deletion
    })
    .catch(error => {
        console.error('Error deleting cart item:', error);
    });
}

document.addEventListener('DOMContentLoaded', function() {
    const userId = localStorage.getItem('userId'); // Replace with actual user ID, perhaps from localStorage or a user context

    // Add event listener for "View Order History" button
    document.getElementById('view-history-btn').addEventListener('click', function() {
        getOrderHistory(userId);
    });

    // Add event listener for "Place Order" button
    document.getElementById('place-order-btn').addEventListener('click', function() {
        placeOrder(userId);
    });
});


function getOrderHistory(userId) {
    fetch(`https://tech-task-24-latest.onrender.com/order/history/${userId}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        }
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to fetch order history');
        }
        return response.json();
    })
    .then(data => {
        displayOrderHistory(data.orders);
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Failed to fetch order history: ' + error.message);
    });
}

function displayOrderHistory(orders) {
    const orderHistoryContainer = document.getElementById('order-history');
    orderHistoryContainer.innerHTML = ''; // Clear previous content

    if (orders.length === 0) {
        orderHistoryContainer.innerHTML = '<p>No orders found.</p>';
        return;
    }

    orders.forEach(order => {
        const orderElement = document.createElement('div');
        orderElement.className = 'order';
        orderElement.innerHTML = `
            <h3>Order ID: ${order.ID}</h3><br>
            <p>Status: ${order.status || 'N/A'}</p><br>
            <p>Total Amount: ${order.total_amount ? '$' + order.total_amount.toFixed(2) : 'N/A'}</p><br>
            <button class="cancelorder" onclick="cancelOrder(${order.user_id}, ${order.ID})">Cancel Order</button>
        `;
        orderHistoryContainer.appendChild(orderElement);
    });
}

async function placeOrder(userId) {
    try {
        const response = await fetch('https://tech-task-24-latest.onrender.com/order', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ user_id: parseInt(userId) }),
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || 'Failed to place order');
        }

        const data = await response.json();
        console.log('Order placed successfully:', data);
        // Handle successful order placement (e.g., clear cart, show confirmation)
    } catch (error) {
        console.error('Error placing order:', error); 
    }
    window.location.reload();  
}

async function cancelOrder(userId, orderId) {
    try {
        const response = await fetch('https://tech-task-24-latest.onrender.com/cancel', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ user_id: userId, order_id: orderId }),
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || 'Failed to cancel order');
        }

        const data = await response.json();
        console.log('Order cancelled successfully:', data);
        // Handle successful cancellation (e.g., update UI, show confirmation)
    } catch (error) {
        console.error('Error cancelling order:', error);
        // Handle error (e.g., show error message to user)
    }
    window.location.reload();  
}


document.addEventListener("DOMContentLoaded", () => {
    const logoutLink = document.querySelector('a[href="/logout"]');
    if (logoutLink) {
        logoutLink.addEventListener("click", (event) => {
            event.preventDefault(); // Prevent default navigation behavior
            // Perform logout actions
            logoutUser();
        });
    }

    function logoutUser() {
        localStorage.removeItem('token'); // Remove JWT token from localStorage
        // You can also clear sessionStorage if needed: sessionStorage.removeItem('token');
        // Redirect to homepage
        window.location.href = 'homepage.html'; // Replace with your homepage URL
    }
});
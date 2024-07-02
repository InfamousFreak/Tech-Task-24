document.addEventListener('DOMContentLoaded', function() {
    fetchCartItems();
});

function fetchCartItems() {
    const userId = localStorage.getItem('userId');

    if (!userId) {
        console.error('User ID not found. Please log in.');
        return;
    }

    fetch(`http://127.0.0.1:8080/cart/${userId}`)
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

    fetch('http://127.0.0.1:8080/cart/delete', {
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

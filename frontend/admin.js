document.addEventListener('DOMContentLoaded', function() {
    fetchUserProfiles();
});

async function fetchUserProfiles() {
    try {
        const response = await fetch('https://tech-task-24-latest-1.onrender.com/profile/show', {
            method: 'GET',
            headers: {
                'Authorization': 'Bearer ' + localStorage.getItem('adminToken')
            }
        });

        if (!response.ok) {
            throw new Error('Failed to fetch user profiles');
        }

        const result = await response.json();
        console.log('Received data:', result); 

        if (result.success) {
            displayUserProfiles(result.data);
        } else {
            throw new Error(result.message || 'Unknown error occurred');
        }
    } catch (error) {
        document.getElementById('errorMessage').textContent = error.message;
    }
}

function displayUserProfiles(users) {
    const userList = document.getElementById('userList');
    userList.innerHTML = '';

    users.forEach(user => {
        const userId = user.ID;
        const userElement = document.createElement('div');
        userElement.className = 'user-item';
        userElement.innerHTML = `
            <p><strong>ID:</strong> ${user.ID}</p>
            <p><strong>Name:</strong> ${user.name}</p>
            <p><strong>Email:</strong> ${user.email}</p>
            <p><strong>Role:</strong> ${user.role}</p>
            <button class="addltcart" onclick="deleteUserProfile(${user.ID})">Delete User</button>
            <span class="cart-toggle" onclick="toggleCart(${user.ID}, this)">Show Cart</span>
            <div class="cart-container" id="cart">
                <h3>User Cart</h3>
                <div id="cart-items">
                    <!-- Cart items will be dynamically inserted here -->
                </div>
                <div class="cart-total">
                    <h4>Total: <span id="cart-total-amount">$0.00</span></h4>
                </div>
            </div>        `;
        userList.appendChild(userElement);
    });
}

async function toggleCart(userId, element) {
    const cartContainer = document.getElementById(`cart`);
    
    if (cartContainer.classList.contains('show')) {
        cartContainer.classList.remove('show');
        element.textContent = 'Show Cart';
    } else {
        cartContainer.classList.add('show');
        element.textContent = 'Hide Cart';
        await fetchCartItems(userId);
    }
}


async function deleteUserProfile(userId) {
    if (!confirm('Are you sure you want to delete this user?')) {
        return;
    }

    try {
        const response = await fetch(`https://tech-task-24-latest-1.onrender.com/profile/${userId}`, {
            method: 'DELETE',
            headers: {
                'Authorization': 'Bearer ' + localStorage.getItem('adminToken')
            }
        });

        const result = await response.json();

        if (response.ok) {
            alert(result.message);
            // Refresh the user list after successful deletion
            fetchUserProfiles();
        } else {
            throw new Error(result.error || 'Failed to delete user');
        }
    } catch (error) {
        console.error('Error deleting user:', error);
        alert('Failed to delete user: ' + error.message);
    }
}

function fetchCartItems(userId) {

    if (!userId) {
        console.error('User ID not found. Please log in.');
        return;
    }

    fetch(`https://tech-task-24-latest-1.onrender.com/cart/${userId}`)
        .then(response => response.json())
        .then(data => {
            displayCartItems(data, userId);
        })
        .catch(error => {
            console.error('Error fetching cart items:', error);
        });
}

function displayCartItems(cartItems, userId) {
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
        deleteBtn.addEventListener('click', () => deleteCartItem(userId, itemId));
        cartItemsContainer.appendChild(itemDiv);

        totalAmount += itemPrice * quantity;
    });

    document.getElementById('cart-total-amount').textContent = `$${totalAmount.toFixed(2)}`;
}

function deleteCartItem(userId, itemId) {
    if (!userId || !itemId) {
        console.error('User ID or Item ID not found.');
        return;
    }

    fetch('https://tech-task-24-latest-1.onrender.com/admin/cart', {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
            // Include authentication token if required
            // 'Authorization': 'Bearer ' + yourAuthToken
        },
        body: JSON.stringify({
            user_id: userId,
            item_id: itemId
        })
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to delete cart item');
        }
        return response.json();
    })
    .then(data => {
        console.log(data.message);
        // Refresh the cart items after successful deletion
        fetchCartItems(userId);
    })
    .catch(error => {
        console.error('Error deleting cart item:', error);
    });
}







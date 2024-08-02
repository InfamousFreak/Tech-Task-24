document.addEventListener("DOMContentLoaded", () => {
    const profileButton = document.getElementById("profile-button");
    const userDetails = document.getElementById("user-details");
    const profileInfo = document.getElementById("profile-info");
    const orderHistory = document.getElementById("order-history");
    const searchButton = document.getElementById('search-button');
    const searchInput = document.getElementById('search-input');
    const resultsList = document.getElementById('results-list');

    const searchResults = document.getElementById('search-results');
if (!searchResults) {
    console.error('Search results container not found in the DOM');
    return;

}

function isLoggedIn() {
    const token = localStorage.getItem('token');
    return !!token; // Returns true if token exists, false otherwise
}       

    profileButton.addEventListener("click", async () => {
        userDetails.classList.toggle("hidden");
        if (!userDetails.classList.contains("hidden")) {
            try {
                const token = localStorage.getItem('token');
                if (!token) {
                    throw new Error('No token found in localStorage');
                }
            

                const payload = JSON.parse(atob(token.split('.')[1]));
                const userId = payload.ID; // Assuming the ID is stored in the token payload
                const response = await fetch(`https://tech-task-24-latest.onrender.com/profile/${userId}/details`, {
                    headers: {
                        'Authorization': `Bearer ${token}`
                    }
                });
                if (!response.ok) {
                    throw new Error('Failed to fetch user details');
                }
                const userData = await response.json();
                // Populate user details with real data
                let profileText = `Name: ${userData.data.name}<br><br>Email: ${userData.data.email}`;
                if (userData.data.city) profileText += `<br><br>City: ${userData.data.city}`;
                if (userData.data.role) profileText += `<br><br>Role: ${userData.data.role}`;
                profileInfo.innerHTML = profileText;// Update this as needed
            } catch (error) {
                console.error('Failed to fetch user details:', error);
                profileInfo.textContent = "Failed to load user details.";
                orderHistory.textContent = "";
            }
        }
    });

    // New code for search functionality
    searchButton.addEventListener('click', performSearch);

    searchInput.addEventListener('keypress', function(e) {
        if (e.key === 'Enter') {
            performSearch();
        }
    });

    function performSearch() {
        const searchTerm = searchInput.value.trim();
        if (searchTerm) {
            const token = localStorage.getItem('token');
            fetch(`https://tech-task-24-latest.onrender.com/menu/search?tag=${encodeURIComponent(searchTerm)}`, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            })
                .then(response => response.json())
                .then(data => {
                    displaySearchResults(data);
                })
                .catch(error => {
                    console.error('Error:', error);
                    alert('An error occurred while searching. Please try again.');
                });
        }
    }
    const closeIcon = document.getElementById('close-icon');
    
    closeIcon.addEventListener('click', () => {
      searchInput.value = '';
    });
    
    searchInput.addEventListener('input', () => {
      if (searchInput.value !== '') {
        closeIcon.style.display = 'inline-block';
      } else {
        closeIcon.style.display = 'none';
      }
    });

    function displaySearchResults(items) {
        console.log('Received items:', items); // Log the received items
        resultsList.innerHTML = ''; // Clear previous results
        if (!Array.isArray(items) || items.length === 0) {
            console.log('No items found or invalid data structure');
            resultsList.innerHTML = '<li>No items found</li>';
        } else {
            items.forEach(item => {
                console.log('Processing item:', item); // Log each item
                const li = document.createElement('li');
                const a = document.createElement('a');
                a.href = `/frontend/usermenu.html`;
                a.textContent = item.name;
                li.appendChild(a);
                resultsList.appendChild(li);
            });
        }
        searchResults.classList.remove('hidden');
        console.log('Search results visibility:', searchResults.style.display); 
    }
});

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

document.addEventListener('DOMContentLoaded', function() {
    const userId = localStorage.getItem('userId'); // Replace with actual user ID, perhaps from localStorage or a user context

    // Add event listener for "View Order History" button
    document.getElementById('view-history-butn').addEventListener('click', function() {
        getOrderHistory(userId);
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
            <h3>Order ID: ${order.ID}</h3>
            <p>Status: ${order.status || 'N/A'}</p>
            <p>Total Amount: ${order.total_amount ? '$' + order.total_amount.toFixed(2) : 'N/A'}</p>
            <h4>Items:</h4>
            <ul>
                ${(order.OrderItems || []).map(item => `
                    <li>Item ID: ${item.item_id}, Quantity: ${item.quantity}</li>
                `).join('')}
            </ul>
            <button class="cancelorder" onclick="cancelOrder(${order.user_id}, ${order.ID})">Cancel Order</button>
        `;
        orderHistoryContainer.appendChild(orderElement);
    });
}


document.addEventListener("DOMContentLoaded", () => {

    
    const slider = document.querySelector('.slider');
    let isDragging = false;
    let startPos = 0;
    let currentTranslate = 0;
    let prevTranslate = 0;
    let animationID = 0;

    // Sample menu items (replace with your actual data)
    const menuItems = [
        { name: "Chicken Burrito", image: "https://img.freepik.com/free-photo/side-view-shawarma-with-fried-potatoes-board-cookware_176474-3215.jpg?ga=GA1.1.1216645911.1718196968&semt=sph" },
        { name: "CheeseBurger", image: "https://img.freepik.com/free-photo/sandwich-with-chicken-burger-tomatoes-lettuce_2829-16577.jpg?semt=sph" },
        { name: "Caesar Salad", image: "https://img.freepik.com/free-photo/delicious-salad-red-plate-with-oil-high-angle-view-wooden-background_176474-3635.jpg?semt=ais_user" },
        { name: "Margherita Pizza", image: "https://img.freepik.com/free-photo/top-view-pizza-wooden-stand-with-tablecloth_176474-2554.jpg?semt=ais_user" },
        { name: "Butter Chicken", image: "https://img.freepik.com/premium-photo/chiken-dish_884653-13429.jpg?semt=ais_user" },
        { name: "Litti Chokha", image: "https://qph.cf2.quoracdn.net/main-qimg-3f43722a6f4b1e0b4f6eda2f2037c4bb-lq" },
        { name: "Sushi", image: "https://img.freepik.com/free-photo/fresh-sushi-with-red-caviar_140725-1264.jpg?t=st=1721542002~exp=1721545602~hmac=712f018378b91c2e670530312cc386521fe848f172f3ad84ce4718fe6aaf3554&w=740" },
        { name: "Tacos", image: "https://img.freepik.com/free-photo/top-view-fresh-mexican-food-with-lime_23-2148614359.jpg?ga=GA1.1.1216645911.1718196968&semt=sph" },
        { name: "Masala Dosa", image: "https://img.freepik.com/premium-photo/tasty-asian-food-plates-table_1037615-1882.jpg?ga=GA1.1.1216645911.1718196968&semt=ais_user" },
        { name: "Samosa", image: "https://img.freepik.com/premium-photo/vegetarian-samosa-indian-special-traditional-street-food-punjabi-snack-generated-by-ai_1038983-13204.jpg?ga=GA1.1.1216645911.1718196968&semt=sph" }
    ];

    function populateSlider() {
        menuItems.forEach((item, index) => {
            const div = document.createElement('div');
            div.className = 'slider-item';
            div.innerHTML = `
                <img src="${item.image}" alt="${item.name}">
                <p>${item.name}</p>
            `;
            slider.appendChild(div);
        });
    }

    function setSliderPosition() {
        slider.style.transform = `translateX(${currentTranslate}px)`;
    }

    function animation() {
        setSliderPosition();
        if (isDragging) requestAnimationFrame(animation);
    }


    function touchEnd() {
        isDragging = false;
        cancelAnimationFrame(animationID);
        slider.style.cursor = 'grab';
        
        const movedBy = currentTranslate - prevTranslate;
        if (movedBy < -100) {
            currentTranslate -= 220; // Move to next item
        } else if (movedBy > 100) {
            currentTranslate += 220; // Move to previous item
        }

        // Ensure we don't scroll past the end or beginning
        const maxTranslate = -(slider.scrollWidth - slider.clientWidth);
        currentTranslate = Math.max(Math.min(currentTranslate, 0), maxTranslate);

        setSliderPosition();
        prevTranslate = currentTranslate;
    }

    // Populate the slider
    populateSlider();

    // Infinite scroll effect
    setInterval(() => {
        currentTranslate -= 1;
        const maxTranslate = -(slider.scrollWidth - slider.clientWidth);
        if (currentTranslate < maxTranslate) {
            currentTranslate = 0;
        }
        setSliderPosition();
        prevTranslate = currentTranslate;
    }, 50);
});


document.getElementById('alert').style.display = 'block';
setTimeout(function() {
  document.getElementById('alert').style.display = 'none';
}, 2000);   
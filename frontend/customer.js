document.addEventListener("DOMContentLoaded", () => {
    const profileButton = document.getElementById("profile-button");
    const userDetails = document.getElementById("user-details");
    const profileInfo = document.getElementById("profile-info");
    const orderHistory = document.getElementById("order-history");
    const settings = document.getElementById("settings");
    const searchButton = document.getElementById('search-button');
    const searchInput = document.getElementById('search-input');
    const resultsList = document.getElementById('results-list');

    const searchResults = document.getElementById('search-results');
if (!searchResults) {
    console.error('Search results container not found in the DOM');
    return;
}

    profileButton.addEventListener("click", async () => {
        userDetails.classList.toggle("hidden");
        if (!userDetails.classList.contains("hidden")) {
            try {
                const token = localStorage.getItem('token');
                // Decode the JWT token to get the user ID
                const payload = JSON.parse(atob(token.split('.')[1]));
                const userId = payload.ID; // Assuming the ID is stored in the token payload
                const response = await fetch(`http://127.0.0.1:8080/profile/${userId}/details`, {
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
                if (userData.data.businessLicense) profileText += `<br><br>Business License: ${userData.data.businessLicense}`;
                profileInfo.innerHTML = profileText;
                orderHistory.textContent = "Order history data not implemented"; // Update this as needed
                settings.textContent = "Settings data not implemented"; // Update this as needed
            } catch (error) {
                console.error('Failed to fetch user details:', error);
                profileInfo.textContent = "Failed to load user details.";
                orderHistory.textContent = "";
                settings.textContent = "";
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
            fetch(`http://127.0.0.1:8080/menu/search?tag=${encodeURIComponent(searchTerm)}`, {
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
                a.href = `/menu-item/${item.item_id}`;
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


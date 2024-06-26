document.addEventListener("DOMContentLoaded", () => {
    const profileButton = document.getElementById("profile-button");
    const userDetails = document.getElementById("user-details");
    const profileInfo = document.getElementById("profile-info");
    const orderHistory = document.getElementById("order-history");
    const settings = document.getElementById("settings");

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

                profileInfo.innerHTML =  profileText;
            
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




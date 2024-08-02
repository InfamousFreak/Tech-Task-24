/*document.addEventListener("DOMContentLoaded", async () => {
    try {
        const response = await fetch("https://tech-task-24-latest.onrender.com/menu/get");
        if (!response.ok) {
            throw new Error("Failed to fetch menu items");
        }
        const menuItems = await response.json();
        
        const mostPopularItemsContainer = document.getElementById("most-popular-items");
        if (!mostPopularItemsContainer) {
            throw new Error("Most Popular Items container not found");
        }
        
        menuItems.forEach(item => {
            const itemElement = document.createElement("div");
            itemElement.classList.add("menu-item");
            itemElement.innerHTML = `
                <h2>${item.name}</h2>
                <p>${item.description}</p>
                <br><p>Price: $${item.price.toFixed(2)}</p>
            `;
            mostPopularItemsContainer.appendChild(itemElement);
        });
    } catch (error) {
        console.error("Error fetching menu items:", error);
        // Optionally, display an error message to the user
        const errorContainer = document.getElementById("error-message");
        if (errorContainer) {
            errorContainer.textContent = "Failed to load menu items. Please try again later.";
        }
    }
});*/


document.addEventListener("DOMContentLoaded", async () => {
    try {
        const response = await fetch("https://tech-task-24-latest.onrender.com/menu/get");
        if (!response.ok) {
            throw new Error("Failed to fetch menu items");

        }
        const menuItems = await response.json();

        
        const mostPopularItemsContainer = document.getElementById("most-popular-items");
        menuItems.forEach(item => {
            const itemElement = createMenuItemElement(item);
            mostPopularItemsContainer.appendChild(itemElement);
        });

    } catch (error) {
        console.error("Error fetching menu items:", error);
        // Handle error, show message, etc.
    }
});


  function createMenuItemElement(item) {
    const itemElement = document.createElement("div");
    itemElement.classList.add("menu-item");



    const imageElement = document.createElement("img");
    imageElement.src = item.imageUrl; // Assuming imageUrl is a property of your MenuItem model
    imageElement.alt = item.name; // Provide appropriate alt text for accessibility
    itemElement.appendChild(imageElement);
    imageElement.addEventListener("click", () => showItemDetails(item));


    const textContainer = document.createElement("div");
    textContainer.classList.add("menu-item-details");

    const itemName = document.createElement("h2");
    itemName.textContent = item.name;
    textContainer.appendChild(itemName);

    const itemDescription = document.createElement("p");
    itemDescription.textContent = item.description;
    textContainer.appendChild(itemDescription);

    const itemPrice = document.createElement("p");
    itemPrice.textContent = `Price: ₹${item.price.toFixed(2)}`;
    textContainer.appendChild(itemPrice);

    itemElement.appendChild(textContainer);
    

    return itemElement;
}

/*function showItemDetails(item) {
    // Create a modal or navigate to a new page
    const detailsElement = document.createElement("div");
    detailsElement.classList.add("item-details-modal");

    detailsElement.innerHTML = `
        <h2>${item.name}</h2>
        <img src="${item.imageUrl}" alt="${item.name}">
        <p>${item.description}</p>
        <p>Price: $${item.price.toFixed(2)}</p>
        <p>Tags: ${item.tags}</p>
        <div class="quantity-control">
            <button class="decrease">-</button>
            <input type="number" value="1" min="1" class="quantity-input">
            <button class="increase">+</button>
        </div>
        <button class="add-to-cart">Add to Cart</button>
    `;

    // Add event listeners for quantity control and add to cart
    // ...

    document.body.appendChild(detailsElement);
}*/

        function showItemDetails(item) {
            const overlay = document.createElement("div");
            overlay.classList.add("modal-overlay");

            const detailsElement = document.createElement("div");
            detailsElement.classList.add("item-details-modal");

            detailsElement.innerHTML = `
                <button class="close-modal">&times;</button>
                <h2>${item.name}</h2>
                <img src="${item.imageUrl}" alt="${item.name}">
                <p>${item.description}</p>
                <p>Price: ₹${item.price.toFixed(2)}</p>
                <p>Tags: ${item.tags}</p>
                <div class="quantity-control">
                    <button class="decrease">-</button>
                    <input type="number" value="1" min="1" class="quantity-input">
                    <button class="increase">+</button>
                </div>
                <button class="add-to-cart">Add to Cart</button>
            `;

            overlay.appendChild(detailsElement);
            document.body.appendChild(overlay);

            // Close modal when clicking outside or on close button
            overlay.addEventListener("click", (e) => {
                if (e.target === overlay || e.target.classList.contains("close-modal")) {
                    document.body.removeChild(overlay);
                }
            });

        // Add event listeners for quantity control and add to cart
        const decreaseBtn = detailsElement.querySelector(".decrease");
        const increaseBtn = detailsElement.querySelector(".increase");
        const quantityInput = detailsElement.querySelector(".quantity-input");
        const addToCartBtn = detailsElement.querySelector(".add-to-cart");

        decreaseBtn.addEventListener("click", () => {
            if (quantityInput.value > 1) {
                quantityInput.value = parseInt(quantityInput.value) - 1;
            }
        });

        increaseBtn.addEventListener("click", () => {
            quantityInput.value = parseInt(quantityInput.value) + 1;
        });

        addToCartBtn.addEventListener("click", async () => {
            const quantity = parseInt(quantityInput.value);
            try {
                await upsertCartItem(item.item_id, quantity);
                // Success handling (e.g., close modal, update UI)
                document.body.removeChild(overlay);
            } catch (error) {
                console.error("Failed to add to cart:", error);
                // Error handling (e.g., show error message to user)
            }
        });
    }


async function upsertCartItem(itemId, quantity) {
    try {
        const userId = localStorage.getItem("userId");
        if (!userId) {
            throw new Error("User not logged in");
        }

        const response = await fetch("https://tech-task-24-latest.onrender.com/cart/upsert", {
            method: "PUT",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                user_id: parseInt(userId),
                item_id: parseInt(itemId),
                quantity: parseInt(quantity)
            })
        });
        const requestBody = {
            user_id: parseInt(userId),
            item_id: parseInt(itemId),
            quantity: parseInt(quantity)
        };
        console.log("Request body:", requestBody);
        
        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(`Failed to add/update cart item: ${errorText}`);
        }

        const result = await response.json();
        console.log("Cart item added/updated successfully:", result);
        return result;
    } catch (error) {
        console.error("Error adding/updating cart item:", error);
        throw error;
    }
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




async function fetchMenuItemsByTag(partialTag) {
    try {
        const response = await fetch(`https://tech-task-24-latest.onrender.com/menu/search?tag=${encodeURIComponent(partialTag)}`);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return await response.json();
    } catch (error) {
        console.error('Error fetching menu items:', error);
        return [];
    }
}

function populateCategory(categoryId, items) {
    const container = document.getElementById(categoryId);
    container.innerHTML = ''; // Clear existing items
    items.forEach(item => {
        const itemElement = createMenuItemElement(item);
        container.appendChild(itemElement);
    });
}



async function initializeMenu() {
    // Fetch and populate starters (assuming 'starter' is a tag)
    const starters = await fetchMenuItemsByTag('starter');
    populateCategory('starters', starters);

    // Fetch and populate desserts (assuming 'dessert' is a tag)
    const desserts = await fetchMenuItemsByTag('dessert');
    populateCategory('desserts', desserts);     
}

// Call this function when the page loads
document.addEventListener('DOMContentLoaded', initializeMenu);



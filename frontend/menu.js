/*document.addEventListener("DOMContentLoaded", async () => {
    try {
        const response = await fetch("http://127.0.0.1:8080/menu/get");
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
        const response = await fetch("http://127.0.0.1:8080/menu/get");
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

    const textContainer = document.createElement("div");
    textContainer.classList.add("menu-item-details");

    const itemName = document.createElement("h2");
    itemName.textContent = item.name;
    textContainer.appendChild(itemName);

    const itemDescription = document.createElement("p");
    itemDescription.textContent = item.description;
    textContainer.appendChild(itemDescription);

    const itemPrice = document.createElement("p");
    itemPrice.textContent = `Price: $${item.price.toFixed(2)}`;
    textContainer.appendChild(itemPrice);

    itemElement.appendChild(textContainer);

    return itemElement;
}


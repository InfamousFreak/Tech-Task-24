document.addEventListener('DOMContentLoaded', function() {
    fetchUserProfiles();
});

async function fetchUserProfiles() {
    try {
        const response = await fetch('http://localhost:8080/profile/show', {
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
            <button onclick="deleteUserProfile(${user.ID})">Delete User</button>
        `;
        userList.appendChild(userElement);
    });
}


async function deleteUserProfile(userId) {
    if (!confirm('Are you sure you want to delete this user?')) {
        return;
    }

    try {
        const response = await fetch(`http://localhost:8080/profile/${userId}`, {
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
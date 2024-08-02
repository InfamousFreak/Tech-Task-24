$(window).on("hashchange", function () {
    if (location.hash.slice(1) == "signup") {
        $(".page").addClass("extend");
        $("#login").removeClass("active");
        $("#signup").addClass("active");
    } else {
        $(".page").removeClass("extend");
        $("#login").addClass("active");
        $("#signup").removeClass("active");
    }
});
$(window).trigger("hashchange");

async function handleAdminLoginSubmit(event) {
    event.preventDefault();

    const email = document.getElementById('adminEmail').value;
    const password = document.getElementById('adminPassword').value;

    if (!email || !password) {
        document.getElementById("errorMsg").innerHTML = "Please fill the required fields";
        return false;
    }

    try {
        const response = await fetch('https://tech-task-24-latest-1.onrender.com/admin/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email, password }),
        });

        if (response.ok) {
            const { token, adminId } = await response.json();
            localStorage.setItem('adminToken', token);
            localStorage.setItem('adminId', adminId);
            window.location.href = 'restaurateur.html';
        } else {
            const error = await response.json();
            document.getElementById('errorMsg').textContent = error.error;
        }
    } catch (error) {
        console.error('Network error:', error);
        document.getElementById('errorMsg').textContent = 'An error occurred during login';
    }
}

async function handleAdminSignupSubmit(event) {
    event.preventDefault();

    const email = document.getElementById('adminSignEmail').value;
    const name = document.getElementById('adminSignName').value;
    const password = document.getElementById('adminSignPassword').value;

    if (!email || !name || !password) {
        document.getElementById("errorMsg").innerHTML = "Please fill the required fields";
        return false;
    }

    if (password.length < 8) {
        document.getElementById("errorMsg").innerHTML = "Your password must include at least 8 characters";
        return false;
    }

    try {
        const response = await fetch('https://tech-task-24-latest-1.onrender.com/admin/signup', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ email, name, password })
        });

        if (response.ok) {
            console.log('Admin registration successful');
            alert("Successfully signed up as admin");
            window.location.hash = "#login";
        } else {
            console.error('Admin registration failed');
            const error = await response.json();
            document.getElementById('errorMsg').textContent = error.error;
        }
    } catch (error) {
        console.error('Network error:', error);
        document.getElementById('errorMsg').textContent = 'An error occurred during signup';
    }
}
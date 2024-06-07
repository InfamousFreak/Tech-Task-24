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

/*function validateLoginForm() {
	var name = document.getElementById("logName").value;
	var password = document.getElementById("logPassword").value;

	if (name == "" || password == "") {
		document.getElementById("errorMsg").innerHTML = "Please fill the required fields"
		return false;
	}

	else if (password.length < 8) {
		document.getElementById("errorMsg").innerHTML = "Your password must include atleast 8 characters"
		return false;
	}
	else {
		alert("Successfully logged in");
		return true;
	}
}*/
function validateSignupForm() {
	var mail = document.getElementById("signEmail").value;
	var name = document.getElementById("signName").value;
	var password = document.getElementById("signPassword").value;

	if (mail == "" || name == "" || password == "") {
		document.getElementById("errorMsg").innerHTML = "Please fill the required fields"
		return false;
	}

	else if (password.length < 8) {
		document.getElementById("errorMsg").innerHTML = "Your password must include atleast 8 characters"
		return false;
	}
	else {
		alert("Successfully signed up");
		return true;
	}
}

document.getElementById('signup-buttom').addEventListener('click', async function(event) {
	event.preventDefault();
	const email = document.getElementById('signEmail').value;
	const password = document.getElementById('signPassword').value;
	const name = document.getElementById('signName').value;

	try {
		const response = await fetch('http://localhost:8080/profile/create', {
			method: 'POST',
			headers: {
				'Content-type': 'application/json'
			},
			body: JSON.stringify({ email, name, password })
		});

		
		if (response.ok) {
			// Handle successful registration
			console.log('Registration successful');
			// Optionally, you can redirect the user or display a success message
		  } else {
			// Handle registration failure
			console.error('Registration failed');
		  }
		} catch (error) {
			console.error('Network error:', error);
		}
	
	});

	async function GoogleAuthURL() {
		try {
			const response = await fetch('http://localhost:8080/google_login');
			if (response.ok) {
				const data = await response.json();
				return data.authURL;
			} else {
				console.error ('Failed to get Google Authentication URL');
			}
		} catch (error) {
			console.error ('Network error:', error);
		}
	}

	document.getElementById('google-authBtn').onclick = async function () {
		const authURL = await GoogleAuthURL();
		if (authURL) {
			window.location.href = authURL;
		} else {
			console.error('Failed to get Google authentication URL');
		}
	};
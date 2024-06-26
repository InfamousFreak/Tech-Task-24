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

function validateLoginForm() {
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
}

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

const signupbtn = document.getElementById('signup-buttom');



signupbtn.addEventListener('click', async function(event) {
	event.preventDefault();
	const email = document.getElementById('signEmail').value;
	const password = document.getElementById('signPassword').value;
	const name = document.getElementById('signName').value;

	if (!name || !email || !password) {
		alert('Please fill in all fields');
		return;
	}

	try {
		const response = await fetch('http://127.0.0.1:8080/profile/create', {
			method: 'POST',
			headers: {
				'Content-type': 'application/json'
			},
			body: JSON.stringify({ email, name, password })
		});

		
		if (response.ok) {
			
			console.log('Registration successful');
			
		  } else {
			
			console.error('Registration failed');
		  }
		} catch (error) {
			console.error('Network error:', error);
		}
	
	});

/*async function GoogleAuthURL() {
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
			window.location.href = 'http://localhost:8080/google_login';
		} else {
			console.error('Failed to get Google authentication URL');
		}
	};*/

	async function handleLoginSubmit(event) {
		event.preventDefault(); // Prevent the default form submission
	  
		// Get the email and password values from the form
		const email = document.getElementById('logEmail').value;
		const password = document.getElementById('logPassword').value;
	  
		try {
		  const response = await fetch('http://localhost:8080/login', {
			method: 'POST',
			headers: {
			  'Content-Type': 'application/json',
			},
			body: JSON.stringify({ email, password }),
		  });
	  
		  if (response.ok) {
			// Login successful
			const { token, userId } = await response.json();
			// Store the token in localStorage or handle it as needed
			localStorage.setItem('token', token);
			localStorage.setItem('userId', userId);
			// Redirect the user to the "Hello User" page
			window.location.href = 'hello.html';
		  } else {
			// Login failed
			const error = await response.json();
			document.getElementById('errorMsg').textContent = error.error;
		  }
		} catch (error) {
		  console.error('Network error:', error);
		  document.getElementById('errorMsg').textContent = 'An error occurred during login';
		}
	  }

	  document.getElementById('google-authBtn').addEventListener('click', async function() {
		window.location.href = 'http://127.0.0.1:8080/google_login'; // Update this URL to match your backend's Google OAuth URL
		const params = new URLSearchParams(window.location.search);
				const code = params.get('code');
				const state = params.get('state');
	
				if (code && state) {
					try {
						const response = await fetch(`http://127.0.0.1:8081/google_callback?code=${code}&state=${state}`);
						const data = await response.json();
	
						if (data.token) {
							// Store the token in session storage
							sessionStorage.setItem('authToken', data.token);
	
							// Redirect to protected page
							window.location.href = 'http://127.0.0.1:8080/hello.html';
						} else {
							console.error('Token not found in response:', data);
						}
					} catch (error) {
						console.error('Failed to fetch token:', error);
					}
				}
			});


	



 
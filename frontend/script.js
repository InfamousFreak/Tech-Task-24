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
signupbtn.addEventListener('click', validateSignupForm);			

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
		const response = await fetch('https://tech-task-24-latest-1.onrender.com/profile/create', {
			method: 'POST',
			headers: {
				'Content-type': 'application/json'
			},
			body: JSON.stringify({ email, name, password })
		});

		
		if (response.ok) {
			
			const { token, userId, } = await response.json();

			localStorage.setItem('token', token);
			localStorage.setItem('userId', userId);
			console.log('Registration successful');
			alert("Successfully signed up.");
			window.location.href = "customer.html";
			
		  } else {
			
			console.error('Registration failed');
		  }
		} catch (error) {
			console.error('Network error:', error);
		}
	
	});


	async function handleLoginSubmit(event) {
		event.preventDefault(); 
	  

		const email = document.getElementById('logEmail').value;
		const password = document.getElementById('logPassword').value;
	  
		try {
		  const response = await fetch('https://tech-task-24-latest-1.onrender.com/login', {
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
			window.location.href = 'customer.html';
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
		
	// Update this URL to match your backend's Google OAuth URL
	/*const params = new URLSearchParams(window.location.search);
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
						window.location.href = 'https://tech-task-24-latest-1.onrender.com/customer.html';
					} else {
						console.error('Token not found in response:', data);
					}
				} catch (error) {
					console.error('Failed to fetch token:', error);
				}						
			}*/


			document.getElementById('google-authBtn').addEventListener('click', async function(event) {
				event.preventDefault();
				window.location.href = 'https://tech-task-24-latest-1.onrender.com/google_login'; 
				const params = new URLSearchParams(window.location.search);
				const code = params.get('code');
				const state = params.get('state');
							try {
								const response = await fetch('https://tech-task-24-latest-1.onrender.com/google_callback?code=' + code + '&state=' + state);
										//     method: 'GET',
										//     headers: {
										//         'Content-Type': 'application/json'
										//     },
										// });
								const data = await response.json();
								console.log(data.token);
			
								if (data.token) {
									
									localStorage.setItem('token', data.token);
			
									//Redirect to protected page
								   window.location.href = 'http://127.0.0.1:5501/frontend/customer.html';
								} else {
									console.error('Token not found in response:', data);
								}
							} catch (error) {
								console.error('Failed to fetch token:', error);
							}
						// }
					});
//f

function togglePasswordVisibility(passwordFieldId) {
	const passwordField = document.getElementById('logPassword');
	const eyeIcon = document.querySelector(`#${passwordFieldId} ~ .toggle-password > svg`);

	if (passwordField.type === 'password') {
		passwordField.type = 'text';
		eyeIcon.classList.remove('bi-eye');
		eyeIcon.classList.add('bi-eye-slash');
	} else {
		passwordField.type = 'password';
		eyeIcon.classList.remove('bi-eye-slash');
		eyeIcon.classList.add('bi-eye');
	}
}

			







 
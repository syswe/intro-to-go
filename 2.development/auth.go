package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/dashboard", dashboardHandler)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method == "POST" {
		// Get the username and password from the request body
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Validate the username and password (e.g., check against a database)
		if username == "admin" && password == "password" {
			// Set a session cookie to indicate successful login
			http.SetCookie(w, &http.Cookie{
				Name:  "session",
				Value: "authenticated",
			})

			// Redirect the user to the dashboard
			http.Redirect(w, r, "/dashboard", http.StatusFound)
			return
		}
	}

	// Display the login form
	fmt.Fprintf(w, `
		<h1>Login</h1>
		<form method="POST" action="/login">
			<label for="username">Username:</label>
			<input type="text" id="username" name="username"><br>
			<label for="password">Password:</label>
			<input type="password" id="password" name="password"><br>
			<input type="submit" value="Login">
		</form>
	`)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the user is authenticated
	cookie, err := r.Cookie("session")
	if err != nil || cookie.Value != "authenticated" {
		// Redirect the user to the login page
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	// Display the dashboard
	fmt.Fprintf(w, "<h1>Welcome to the Dashboard!</h1>")
}

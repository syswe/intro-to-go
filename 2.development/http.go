package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
func cartHandler(w http.ResponseWriter, r *http.Request) {
	// Handle cart requests here
	fmt.Fprintf(w, "This is your shopping cart!")
}

func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	// Handle checkout requests here
	fmt.Fprintf(w, "Checkout successful!")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/cart", cartHandler)         // Add cart route
	http.HandleFunc("/checkout", checkoutHandler) // Add checkout route

	log.Fatal(http.ListenAndServe(":8080", nil))
}
func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Handle API requests here
	switch r.URL.Path {
	case "/api/products":
		// Handle GET request for products
		// Return a list of products in JSON format
		products := []string{"Product 1", "Product 2", "Product 3"}
		json.NewEncoder(w).Encode(products)
	case "/api/cart":
		// Handle GET request for cart
		// Return the contents of the shopping cart in JSON format
		cart := []string{"Product 1", "Product 2"}
		json.NewEncoder(w).Encode(cart)
	case "/api/checkout":
		// Handle POST request for checkout
		// Process the checkout and return a success message in JSON format
		response := map[string]string{"message": "Checkout successful!"}
		json.NewEncoder(w).Encode(response)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/cart", cartHandler)
	http.HandleFunc("/checkout", checkoutHandler)
	http.HandleFunc("/api", apiHandler) // Add API route

	log.Fatal(http.ListenAndServe(":8080", nil))
}

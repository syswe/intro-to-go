package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read data from the client
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	// Process the data
	data := buffer[:n]
	fmt.Println("Received data:", string(data))

	// Send a response back to the client
	response := []byte("Hello from the server!")
	_, err = conn.Write(response)
	if err != nil {
		fmt.Println("Error sending response:", err)
		return
	}
}

func main() {
	// Start a TCP server on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started, listening on port 8080")

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle each connection in a separate goroutine
		go handleConnection(conn)
	}
}
// Track network statistics
go trackNetwork()

// Function to track network statistics
func trackNetwork() {
	// Create a ticker to track statistics every second
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	// Variables to store network statistics
	var totalConnections int
	var totalBytesRead int64
	var totalBytesWritten int64

	// Start tracking statistics
	for range ticker.C {
		// Get current statistics
		connections := listener.(net.StatConn).Stats().CurrEstab
		bytesRead := listener.(net.StatConn).Stats().ReadBytes
		bytesWritten := listener.(net.StatConn).Stats().WriteBytes

		// Calculate the difference in statistics since the last check
		connectionsDiff := connections - totalConnections
		bytesReadDiff := bytesRead - totalBytesRead
		bytesWrittenDiff := bytesWritten - totalBytesWritten

		// Update the total statistics
		totalConnections = connections
		totalBytesRead = bytesRead
		totalBytesWritten = bytesWritten

		// Print the network statistics
		fmt.Printf("Connections: %d, Bytes Read: %d, Bytes Written: %d\n", connectionsDiff, bytesReadDiff, bytesWrittenDiff)
	}
}
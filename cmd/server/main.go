package main

import (
	"fmt"
	"http-server-GO/internal/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/load", handlers.LoadHandler)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}

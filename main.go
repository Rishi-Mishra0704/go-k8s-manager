package main

import (
	"fmt"
	"log"
)

func main() {
	server := NewServer(":8080")
	if err := server.Start(); err != nil {
		log.Fatal(fmt.Sprintf("Failed to start server: %v", err))
	}
}

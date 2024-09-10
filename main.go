package main

import (
	"flag"
	"fmt"
)

func main() {

	action := flag.String("action", "serve", "Action to perform: serve | dockerize | k8s-deploy")
	flag.Parse()

	switch *action {
	case "serve":
		server := NewServer(":8080")
		if err := server.Start(); err != nil {
			panic(err)
		}
	case "k8s-deploy":
		deployToK8s()
	default:
		fmt.Println("Unknown action. Use 'serve' or 'k8s-deploy'.")
	}
}

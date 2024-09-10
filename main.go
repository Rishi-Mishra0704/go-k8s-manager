package main

func main() {
	server := NewServer(":8080")
	if err := server.Start(); err != nil {
		panic(err)
	}
}

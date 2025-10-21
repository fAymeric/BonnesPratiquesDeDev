package main

import (
	"fmt"
	"hangman_web/internal/server"
)

func main() {
	fmt.Println("Server started at http://localhost:8080")
	server.StartServer()
}

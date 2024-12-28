package main

import (
	"fmt"
	"os"
)

func main() {
	// Load `.env file
	data, err := os.ReadFile(".env")
	if err != nil {
		fmt.Println("Error reading .env file: ", err)
		return
	}

	// Print the content of the file
	fmt.Println("=== Content of .env file ===")
	fmt.Println(string(data))
	fmt.Println("=============================")
}

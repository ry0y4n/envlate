package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Load `.env file
	file, err := os.Open(".env")
	if err != nil {
		fmt.Println("Can't open .env file ", err)
		return
	}
	defer file.Close()

	// Parse `.env` file
	fmt.Println("=====Parse `.env` file=====")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "#") {
			continue
		}

		fmt.Println("Each line: ", trimmedLine)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error loading .env file: ", err)
	}
	fmt.Println("=========================")
}

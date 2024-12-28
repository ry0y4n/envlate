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

	// key-value map for env variables
	envMap := make(map[string]string)

	// Parse `.env` file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "#") {
			continue
		}
		
		// Split key and value
		splitIndex := strings.Index(trimmedLine, "=")
		if (splitIndex == -1) {
			fmt.Println("Warning: invalid line (no '='):", trimmedLine)
			continue
		}
		key := trimmedLine[:splitIndex]
		value := trimmedLine[splitIndex+1:]

		// Remove spaces
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		// Add to map
		envMap[key] = value
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading .env file: ", err)
	}

	fmt.Println("====== Key/Value Pairs ======")
	for k, v := range envMap {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}
	fmt.Println("=============================")
}

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func loadEnvFile(filename string) (map[string]string, error) {
	// Check if file exists
	info, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("%s does not exist", filename)
	} else if err != nil {
		return nil, fmt.Errorf("could not stat %s: %w", filename, err)
	}

	// check if file is a directory
	if info.IsDir() {
		return nil, fmt.Errorf("%s is a directory, not a file", filename)
	}

	// Load `.env file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %w", filename, err)
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
		if splitIndex == -1 {
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
		return nil, fmt.Errorf("error while reading %s: %w", filename, err)
	}

	return envMap, nil
}

func main() {
	fileFlag := flag.String("file", ".env", "path to .env file")
	flag.StringVar(fileFlag, "f", ".env", "short alias for --file")
	flag.Parse()

	// Load env file
	envMap, err := loadEnvFile(*fileFlag)
	if err != nil {
		fmt.Println("[ERROR] ", err)
		fmt.Println("Please make sure the file exists or specify the correct file name.")
		os.Exit(1)
	}

	// Create .env.template file
	templateFilename := fmt.Sprintf("%s.template", *fileFlag)
	templateFile, err := os.Create(templateFilename)
	if err != nil {
		fmt.Println("Cloud not create .env.template file: ", err)
		os.Exit(1)
	}
	defer templateFile.Close()

	// Write key-value pairs to .env.template file
	writer := bufio.NewWriter(templateFile)
	for k := range envMap {
		line := fmt.Sprintf("%s=\n", k)
		_, err := writer.WriteString(line)
		if err != nil {
			fmt.Println("Error while writing to template file: ", err)
			os.Exit(1)
		}
	}
	// Flush the buffer
	if err := writer.Flush(); err != nil {
		fmt.Println("Error while flushing write: ", err)
		os.Exit(1)
	}

	fmt.Println("Successfully created .env.template")
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var file string

func loadEnvFile(filename string) ([]string, map[string]string, error) {
	// Check if file exists
	info, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil, fmt.Errorf("%s does not exist", filename)
	} else if err != nil {
		return nil, nil, fmt.Errorf("could not stat %s: %w", filename, err)
	}

	// check if file is a directory
	if info.IsDir() {
		return nil, nil, fmt.Errorf("%s is a directory, not a file", filename)
	}

	// Load `.env file
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("could not open %s: %w", filename, err)
	}
	defer file.Close()

	// key-value map for env variables
	keys := make([]string, 0)
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
		keys = append(keys, key)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error while reading %s: %w", filename, err)
	}

	return keys, envMap, nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "envlate",
	Short: "Create a template file from an existing .env file",
	Long: `This CLI tool generates a template file from an existing environment variables file.

By default, it looks for a .env file in the current directory. However, you can specify any environment variables file, such as .env.local, using the --file option.

This makes it easier to manage and share environment variable templates across different environments or projects.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// この中に main() の「本体部分」をコピペ＆調整
		keys, envMap, err := loadEnvFile(file)
		if err != nil {
			fmt.Println("[ERROR] ", err)
			fmt.Println("Please make sure the file exists or specify the correct file name.")
			os.Exit(1)
		}

		// Create .env.template file
		templateFilename := fmt.Sprintf("%s.template", file)
		templateFile, err := os.Create(templateFilename)
		if err != nil {
			fmt.Println("Could not create .env.template file: ", err)
			os.Exit(1)
		}
		defer templateFile.Close()

		// Write key-value pairs to .env.template file
		writer := bufio.NewWriter(templateFile)

		// Write in the order of the saved "keys"
		for _, k := range keys {
			if k == "" {
				continue
			}

			_ = envMap[k]

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
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.envlate.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle").S
	rootCmd.Flags().StringVarP(&file, "file", "f", ".env", "path to .env file")
}

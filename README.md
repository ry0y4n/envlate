# Envlate

Envlate is a CLI tool, written in Go, that helps you generate template files (e.g., `.env.template` or `.env.local.template`) from existing `.env` or `.env.local` files. This makes it easier to manage and share environment variable configurations across different environments or with other developers.

> **Thank you for using Envlate!** If you find it useful, consider giving the repository a star and sharing it with others who might benefit.

## Features

- **Automatic template generation**  
  Quickly create a template version of your `.env` or `.env.local` file to share or track.
- **Flexible file targeting**  
  Use the `--file` option to specify any file path, and Envlate will generate a template in the same directory.
- **Simple command**  
  Just run `envlate` (with no options) to automatically look for a `.env` file in the current directory and create a `.env.template`.

## Installation

### macOS / Linux

```bash
curl -fsSL https://raw.githubusercontent.com/ry0y4n/envlate/main/install.sh | sh
```

### Windows

```powershell
powershell -Command "Invoke-WebRequest -Uri https://raw.githubusercontent.com/ry0y4n/envlate/main/install.ps1 -OutFile install.ps1; .\install.ps1"
```

Once the script completes, you can use the `envlate` command in your terminal.

## Usage

```bash
# Generate a template from the .env file in the current directory
envlate
```

Envlate will look for a `.env` file in the current directory. If found, it creates a `.env.template` in the same location.

### Specify a File

```bash
envlate --file path/to/envfile
```

Replace `path/to/envfile` with the location of your environment file (e.g., `.env.production`). Envlate will then generate a corresponding template file in the same directory.

## Development

If you want to build Envlate from source or run it locally:

```bash
# Build and run
go run -o envlate

# Execute
./envlate --file path/to/envfile
```

### Using Goreleaser

```bash
# Create a local snapshot release
goreleaser release --snapshot --clean
```

## Contributing

Contributions are welcome! If you encounter any issues or have suggestions for improvements, please [open an issue](https://github.com/ry0y4n/envlate/issues) or submit a Pull Request.

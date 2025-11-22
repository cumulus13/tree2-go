# Tree2

A beautiful and feature-rich directory tree visualization tool written in Go with colors, emojis, and gitignore support.

![Tree2 Go Example](https://via.placeholder.com/800x400.png?text=Tree2+Go+Directory+Listing+Example)

## Features

- 🎨 **Colorful Output**: Beautiful colored output with emojis for better visualization
- 📊 **File Sizes**: Human-readable file sizes with color-coded values
- 🔒 **Permission Handling**: Gracefully handles permission denied errors
- 📋 **Exclusion Support**: Supports `.gitignore` files and custom exclude patterns
- 🚀 **High Performance**: Compiled Go binary for fast execution
- 📦 **Single Binary**: No dependencies, easy to distribute
- 💻 **Cross-Platform**: Works on Windows, macOS, and Linux

## Color Scheme

- **Folders**: Yellow (#FFFF00) with 📁 emoji
- **Files**: Cyan (#00FFFF) with 📄 emoji  
- **Size Values**: Light magenta (red if size is 0)
- **Size Units**: Orange suffix
- **Permission Denied**: Red with 🔒 emoji

## Installation

### Method 1: Install from source
```bash
go install github.com/cumulus13/tree2-go@latest
```

### Method 2: Clone and build
```bash
git clone https://github.com/cumulus13/tree2-go
cd tree2-go
go build -o tree2 tree2.go
```

### Method 3: Download pre-built binary
Check the [Releases page](https://github.com/cumulus13/tree2-go/releases) for pre-built binaries for your platform.

## Usage

### Basic Usage

```bash
# Show current directory tree
tree2

# Show specific directory
tree2 /path/to/directory
```

### With Exclusions

```bash
# Exclude patterns (comma-separated)
tree2 -e node_modules,.git,target,dist

# Using long form
tree2 --exclude=__pycache__,*.tmp,temp
```

### Examples

```bash
# Typical project directory
tree2 -e node_modules,.git,target,dist,build

# System directory (with common exclusions)
tree2 /etc -e *.bak,*.tmp,backup

# Multiple exclude patterns
tree2 -e "node_modules,.git,__pycache__,target,dist,build,*.log"
```

## Command Line Options

```
Usage: tree2 [path] [options]

Print directory tree with file sizes, exclusions, and .gitignore support.

Options:
  -e, --exclude string    Exclude patterns (comma-separated)
  -h, --help             Show help message

Examples:
  tree2                          # Current directory
  tree2 /path/to/directory       # Specific directory  
  tree2 -e node_modules,.git     # Exclude patterns
```

## Output Example

```
📂 /home/user/project/
├── 📁 src/
│   ├── 📄 main.go (12.45 KB)
│   └── 📄 utils.go (0.00 B)
├── 📁 pkg/
│   └── 📄 module.go (2.10 KB)
├── 📄 go.mod (1.20 KB)
├── 📄 README.md (4.50 KB)
└── 🔒 [Permission Denied]
```

## Project Structure

```
tree2-go/
├── tree2.go          # Main source code
├── go.mod            # Go module file
├── LICENSE           # MIT License
└── README.md         # This file
```

## Building from Source

### Prerequisites
- Go 1.19 or higher

### Build Steps
```bash
git clone https://github.com/cumulus13/tree2-go
cd tree2-go

# Build for current platform
go build -o tree2 tree2.go

# Build for specific platforms
GOOS=linux GOARCH=amd64 go build -o tree2-linux tree2.go
GOOS=windows GOARCH=amd64 go build -o tree2-windows.exe tree2.go
GOOS=darwin GOARCH=amd64 go build -o tree2-macos tree2.go
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Author

**Hadi Cahyadi**
- Email: cumulus13@gmail.com
- GitHub: [cumulus13](https://github.com/cumulus13)

## Support

If you encounter any issues or have questions:

1. Check the [Issues](https://github.com/cumulus13/tree2-go/issues) page
2. Create a new issue with detailed description
3. Contact: cumulus13@gmail.com

[![Buy Me a Coffee](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/cumulus13)

[![Donate via Ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/cumulus13)
 
[Support me on Patreon](https://www.patreon.com/cumulus13)

---

**Enjoy visualizing your directory structures with Tree2 Go!** 🎉

---


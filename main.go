// File: main.go
// Author: Hadi Cahyadi <cumulus13@gmail.com>
// Date: 2025-11-22
// Description: A beautiful and feature-rich directory tree visualization tool written in Go with colors, emojis, and gitignore support.
// License: MIT

package main

import (
	"flag"
	"fmt"
	//"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// ANSI Color Codes dengan True Color (24-bit)
const (
	ColorReset        = "\033[0m"
	// ColorRed          = "\033[91m"
	ColorRed          = "\033[1;97;41m" 
	ColorYellow       = "\033[93m" // Fallback yellow
	ColorCyan         = "\033[96m" // Fallback cyan
	ColorOrange       = "\033[38;5;214m"
	ColorLightMagenta = "\033[38;5;213m"
	
	// True Color (24-bit) ANSI codes
	ColorBrightYellow = "\033[38;2;255;255;0m"    // #FFFF00
	ColorBrightCyan   = "\033[38;2;0;255;255m"    // #00FFFF
	ColorLightMagentaTrue = "\033[38;2;255;128;255m" // Light magenta
)

type Config struct {
	Excludes     []string
	RootExcludes []string
}

func humanSize(size int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB"}
	floatSize := float64(size)

	for _, unit := range units {
		if floatSize < 1024 {
			return fmt.Sprintf("%.2f %s", floatSize, unit)
		}
		floatSize /= 1024
	}
	return fmt.Sprintf("%.2f PB", floatSize)
}

func loadGitignore(path string) []string {
	gitignorePath := filepath.Join(path, ".gitignore")
	if _, err := os.Stat(gitignorePath); os.IsNotExist(err) {
		return []string{}
	}

	content, err := os.ReadFile(gitignorePath)
	if err != nil {
		return []string{}
	}

	var ignores []string
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		ignores = append(ignores, strings.TrimSuffix(line, "/"))
	}
	return ignores
}

func shouldExclude(entry string, excludes []string, rootExcludes []string) bool {
	for _, ex := range excludes {
		if entry == ex || strings.HasPrefix(entry, ex) {
			return true
		}
	}

	for _, ex := range rootExcludes {
		if entry == ex || strings.HasPrefix(entry, ex) {
			return true
		}
	}
	return false
}

func printTree(path string, prefix string, config *Config) {
	entries, err := os.ReadDir(path)
	if err != nil {
		permissionText := prefix + "└── 🔒 [Permission Denied]"
		fmt.Println(ColorRed + permissionText + ColorReset)
		return
	}

	// Sort entries
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	for idx, entry := range entries {
		if shouldExclude(entry.Name(), config.Excludes, config.RootExcludes) {
			continue
		}

		fullPath := filepath.Join(path, entry.Name())
		connector := "└── "
		if idx < len(entries)-1 {
			connector = "├── "
		}

		if entry.IsDir() {
			// Folder dengan warna kuning terang (#FFFF00)
			folderText := fmt.Sprintf("%s%s📁 %s/", prefix, connector, entry.Name())
			fmt.Println(ColorBrightYellow + folderText + ColorReset)

			newPrefix := prefix
			if idx == len(entries)-1 {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}
			printTree(fullPath, newPrefix, config)
		} else {
			info, err := entry.Info()
			if err != nil {
				continue
			}

			size := info.Size()
			sizeStr := humanSize(size)
			parts := strings.Split(sizeStr, " ")
			sizeValue, sizeUnit := parts[0], parts[1]

			// File dengan warna cyan terang (#00FFFF)
			filePart := fmt.Sprintf("%s%s📄 %s (", prefix, connector, entry.Name())
			fmt.Print(ColorBrightCyan + filePart + ColorReset)

			// Size value
			if size == 0 {
				fmt.Print(ColorRed + sizeValue + ColorReset)
			} else {
				fmt.Print(ColorLightMagentaTrue + sizeValue + ColorReset)
			}

			fmt.Print(" ")

			// Size unit dengan warna orange
			fmt.Print(ColorOrange + sizeUnit + ColorReset)
			fmt.Println(")")
		}
	}
}

func main() {
	var (
		excludeList string
		helpFlag    bool
	)

	flag.StringVar(&excludeList, "e", "", "Exclude patterns (comma-separated)")
	flag.StringVar(&excludeList, "exclude", "", "Exclude patterns (comma-separated)")
	flag.BoolVar(&helpFlag, "h", false, "Show help")
	flag.BoolVar(&helpFlag, "help", false, "Show help")

	flag.Usage = func() {
		fmt.Printf("Usage: %s [path] [options]\n\n", os.Args[0])
		fmt.Println("Print directory tree with file sizes, exclusions, and .gitignore support.")
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
		fmt.Println("\nExamples:")
		fmt.Printf("  %s                    # Current directory\n", os.Args[0])
		fmt.Printf("  %s /path/to/dir       # Specific directory\n", os.Args[0])
		fmt.Printf("  %s -e node_modules,.git  # Exclude patterns\n", os.Args[0])
	}

	flag.Parse()

	if helpFlag {
		flag.Usage()
		return
	}

	path := "."
	if flag.NArg() > 0 {
		path = flag.Arg(0)
	}

	var excludes []string
	if excludeList != "" {
		excludes = strings.Split(excludeList, ",")
		for i, ex := range excludes {
			excludes[i] = strings.TrimSpace(ex)
		}
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	gitignoreExcludes := loadGitignore(absPath)

	config := &Config{
		Excludes:     excludes,
		RootExcludes: gitignoreExcludes,
	}

	// Print root directory dengan warna kuning terang
	rootText := fmt.Sprintf("📂 %s/", absPath)
	fmt.Println(ColorBrightYellow + rootText + ColorReset)

	printTree(absPath, "", config)
}

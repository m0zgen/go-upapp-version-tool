package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	appVersion  = "v1.1.1"
	versionFile = "version.txt"
)

func main() {
	greenBold := "\033[1;32m"
	reset := "\033[0m"

	fmt.Println("──────────────────────────────")
	fmt.Printf(" %s%s%s %s\n", greenBold, "Update Releaser Tool", reset, appVersion)
	fmt.Println("──────────────────────────────")

	data, err := os.ReadFile(versionFile)
	if err != nil {
		log.Fatalf("Failed to read version file: %v", err)
	}

	version := strings.TrimSpace(string(data))
	fmt.Printf("Current version: %s\n", version)

	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		log.Fatalf("Invalid version format: %s", version)
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		log.Fatalf("Failed to parse patch version: %v", err)
	}

	patch++
	parts[2] = strconv.Itoa(patch)

	newVersion := strings.Join(parts, ".")

	if err := os.WriteFile(versionFile, []byte(newVersion+"\n"), 0644); err != nil {
		log.Fatalf("Failed to write new version: %v", err)
	}

	fmt.Printf("New version: %s\n", newVersion)
}
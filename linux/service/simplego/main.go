package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	fmt.Println("This will go to stdout")
	fmt.Fprintln(os.Stderr, "This will go to stderr")

	// Get path of the running binary
	execPath, err := os.Executable()
	if err != nil {
		log.Fatal("Failed to get executable path:", err)
	}

	// Use the directory of the executable for log placement
	execDir := filepath.Dir(execPath)
	logFilePath := filepath.Join(execDir, "simplego.log")

	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println("=== Starting homecheck service ===")

	// Get user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		log.Println("os.UserHomeDir() failed:", err)
	} else {
		log.Println("User home directory:", home)
	}
}

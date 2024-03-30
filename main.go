package main

import "github.com/nitintf/navi/cmd"

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Error getting current exectuable path: ", err)
	}
	exeDir := filepath.Dir(exePath)
	envPath := filepath.Join(exeDir, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	cmd.Execute()
}

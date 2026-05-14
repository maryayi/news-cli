package main

import (
	"github.com/joho/godotenv"
	"github.com/maryayi/news-cli/cmd"
)

func main() {
	// Load .env if present; ignore error so real env vars still work without a file
	godotenv.Load()
	cmd.Execute()
}

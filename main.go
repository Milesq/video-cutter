package main

import (
	_ "embed"
	"fmt"
	"os"
)

const configPath = "config.json"

//go:embed configTemplate.json
var configTemplate []byte

func main() {
	if !fileExists(configPath) {
		createDefaultConfig()
		return
	}

	fmt.Println("cut...")
	config := getConfigFromFile()
	cutter(config)
}

func createDefaultConfig() {
	os.WriteFile(configPath, configTemplate, 0666)
}

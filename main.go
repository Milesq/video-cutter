package main

import "fmt"

const configPath = "config.json"

func main() {
	fmt.Println("cut...")
	config := getConfigFromFile()

	cutter(config)
}

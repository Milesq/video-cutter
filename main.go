package main

import (
	"fmt"
)

func main() {
	fmt.Println("cut...")
	config := getConfigFromFile()

	cutter(config)
}

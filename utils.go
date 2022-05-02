package main

import (
	"io/fs"
	"os"
	"time"
)

func parseDuration(part string) time.Duration {
	dur, err := time.ParseDuration(part)

	if err != nil {
		panic(part)
	}

	return dur
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func safeMkdir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, fs.ModeDir)
	}
}

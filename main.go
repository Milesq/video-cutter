package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/jtguibas/cinema"
)

type Config struct {
	Name       string      `json:"name"`
	Filename   string      `json:"filename"`
	Timestamps [][2]string `json:"timestamps"`
}

func getConfig() (config []Config) {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return
}

func parseDuration(part string) time.Duration {
	dur, err := time.ParseDuration(part)

	if err != nil {
		panic(part)
	}

	return dur
}

func mkdir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, fs.ModeDir)
	}
}

func main() {
	fmt.Println("cut...")
	config := getConfig()

	for _, group := range config {
		for i, part := range group.Timestamps {
			video, _ := cinema.Load(group.Filename)
			beg := parseDuration(part[0])
			end := parseDuration(part[1])

			mkdir(group.Name)

			video.Trim(beg, end)
			video.Render(group.Name + "/" + strconv.Itoa(i) + ".mp4")
		}
	}
}

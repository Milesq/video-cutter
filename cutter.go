package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/jtguibas/cinema"
)

type Config struct {
	Name       string      `json:"name"`
	Filename   string      `json:"filename"`
	Timestamps [][2]string `json:"timestamps"`
}

func getConfigFromFile() (config []Config) {
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

func cutter(config []Config) {
	for _, group := range config {
		for i, part := range group.Timestamps {
			video, _ := cinema.Load(group.Filename)
			beg := parseDuration(part[0])
			end := parseDuration(part[1])

			safeMkdir(group.Name)

			video.Trim(beg, end)
			video.Render(group.Name + "/" + strconv.Itoa(i) + ".mp4")
		}
	}
}

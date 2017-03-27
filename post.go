package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli"
	"log"
)

const configFile = "config.toml"
const doing = "🔁"
const done = "☑"
const NEWLINE = "\n"

type SlackType struct {
	Text string `json:"text"`
}

type Config struct {
	API SlackConfig
}

type SlackConfig struct {
	Url string
}

func post(c *cli.Context) error {
	var config Config
	_, err := toml.DecodeFile(configFile, &config)
	if err != nil {
		return err
	}

	tasks, err := readTasks(file)
	if err != nil {
		return err
	}

	var bufferDo bytes.Buffer
	var bufferDone bytes.Buffer
	for _, v := range tasks.Tasks {
		if v.Completed {
			bufferDo.WriteString(doing)
			bufferDo.WriteString(" ")
			bufferDo.WriteString(v.Title)
			bufferDo.WriteString(NEWLINE)
		} else {
			bufferDone.WriteString(done)
			bufferDone.WriteString(" ")
			bufferDone.WriteString(v.Title)
			bufferDone.WriteString(NEWLINE)
		}
	}
	log.Println(bufferDo.String())
	log.Println(bufferDone.String())
	slack := SlackType{
		Text: "*Doing*" +
			NEWLINE +
			bufferDo.String() +
			NEWLINE +
			NEWLINE +
			"*Completed*" +
			NEWLINE +
			bufferDone.String(),
	}
	body, err := json.Marshal(slack)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", config.API.Url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return err
}

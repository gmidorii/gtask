package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"log"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli"
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
			bufferDone = slackText(bufferDone, done, v)
		} else {
			bufferDo = slackText(bufferDo, doing, v)
		}
	}
	log.Println(bufferDo.String())
	log.Println(bufferDone.String())
	slack := SlackType{
		Text: "*Task List*" +
			NEWLINE +
			"---------------------------------" +
			NEWLINE +
			"`Doing`" +
			NEWLINE +
			bufferDo.String() +
			NEWLINE +
			NEWLINE +
			"---------------------------------" +
			NEWLINE +
			"`Completed`" +
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

func slackText(buf bytes.Buffer, mark string, v Task) bytes.Buffer {
	buf.WriteString(mark)
	buf.WriteString(" ")
	buf.WriteString(v.Title)
	buf.WriteString("     `")
	buf.WriteString(v.DeadLine)
	buf.WriteString("`")
	buf.WriteString(NEWLINE)
	return buf
}
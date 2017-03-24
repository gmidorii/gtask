package main

import (
	"net/http"
	"github.com/urfave/cli"
	"encoding/json"
	"bytes"
	"github.com/BurntSushi/toml"
	"fmt"
)

const configFile = "config.toml"
const notCompleted = "🔁"
const completed = "☑"

type SlackType struct{
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

	var textC string
	var textNC string
	for _, v := range tasks.Tasks {
		if v.Completed {
			textC += notCompleted + " " + v.Title + "\n"
			continue
		}
		textNC += completed + " " + v.Title + "\n"
	}
	fmt.Println(textC)
	fmt.Println(textNC)
	slack := SlackType{Text: textC + "\n\n" + textNC}
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

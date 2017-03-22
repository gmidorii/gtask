package main

import (
	"net/http"
	"github.com/urfave/cli"
	"encoding/json"
	"bytes"
	"github.com/BurntSushi/toml"
)

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
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		return err
	}


	slack := SlackType{Text: "Test Text"}
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

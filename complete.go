package main

import (
	"log"

	"github.com/urfave/cli"
)

func complete(c *cli.Context) error {
	fId := c.Int("i")
	tasks, err := ReadTask(taskfile)
	if err != nil {
		log.Fatal(err)
	}
	if err := completeTask(fId, tasks.Tasks); err != nil {
		return err
	}
	return nil
}

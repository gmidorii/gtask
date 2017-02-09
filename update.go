package main

import (
	"errors"

	"github.com/urfave/cli"
)

func update(c *cli.Context) error {
	fId := c.Int("i")
	if fId == -1 {
		return errors.New("error flag")
	}
	fTask := c.String("t")
	fDays := c.Int("d")
	tasks, err := readTasks(file)
	if err != nil {
		return err
	}

	new := make([]Task, 0, 0)
	for _, old := range tasks.Tasks {
		if old.Id == fId {
			if fTask != "" {
				old.Title = fTask
			}
			if fDays != -1 {
				old.DeadLine = generateDate(fDays, layout)
			}
		}
		new = append(new, old)
	}

	newTasks := Tasks{new}
	return writeTasks(newTasks)
}

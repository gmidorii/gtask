package main

import (
	"log"

	"errors"
	"fmt"
	"strconv"

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

func completeTask(id int, tasks []Task) error {
	newTasks := make([]Task, 0, 0)
	for _, v := range tasks {
		if id == v.Id {
			v.Completed = true
			fmt.Println(colorString(red, "-- Completed --"))
			PrintOneTask(v.Id, v.Title, v.DeadLine)
		}
		newTasks = append(newTasks, v)
	}
	if len(newTasks) == 0 {
		return errors.New("Not found id: " + strconv.Itoa(id))
	}
	return WriteTask(Tasks{newTasks})
}

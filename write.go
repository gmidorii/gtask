package main

import (
	"log"

	"fmt"
	"time"

	"github.com/urfave/cli"
)

func Write(c *cli.Context) error {
	fTask := c.String("t")
	fDays := c.Int("d")
	tasks, err := ReadTask(taskfile)
	if err != nil {
		log.Fatal(err)
	}

	if fDays == -1 {
		// default day is (plus) 3
		fDays = 3
	}
	date := generateDate(fDays, layout)
	appendTask(&tasks, fTask, date)

	WriteTask(tasks)
	if err != nil {
		return err
	}
	return nil
}

func generateDate(plusDays int, layout string) string {
	now := time.Now().AddDate(0, 0, plusDays)
	return now.Format(layout)
}

func appendTask(tasks *Tasks, title string, deadline string) {
	var id int
	if taskSlice := tasks.Tasks; len(taskSlice) != 0 {
		id = tasks.Tasks[len(tasks.Tasks)-1].Id + 1
	} else {
		id = 1
	}
	task := Task{
		Id:        id,
		Title:     title,
		DeadLine:  deadline,
		Completed: false,
	}

	tasks.Tasks = append(tasks.Tasks, task)
	fmt.Println(colorString(blue, "-- Append --"))
	PrintOneTask(id, title, deadline)
}

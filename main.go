package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"time"

	"github.com/urfave/cli"
)

// console color
const cyan = "\u001b[36m"
const red = "\u001b[31m"
const blue = "\u001b[34m"
const reset = "\u001b[0m"

const file = "./tasks/task.json"

var now = time.Now()

// Date Layout
const layout = "2006/01/02"

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	DeadLine  string `json:"dead_line"`
	Completed bool   `json:"completed"`
}

func main() {
	// Flag
	var fComp bool

	app := cli.NewApp()

	// App info
	app.Name = "gtask"
	app.Usage = "Task management"
	app.Version = "0.1"

	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "add Task",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "t",
					Usage: "Task Title",
					Value: "Task",
				},
				cli.IntFlag{
					Name:  "d",
					Usage: "Deadline num from today",
					Value: -1,
				},
			},
			Action: write,
		},
		{
			Name:  "finish",
			Usage: "finished task",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "i",
					Usage: "id",
				},
			},
			Action: complete,
		},
		{
			Name:  "print",
			Usage: "Print Task",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "c",
					Usage:       "Completed flag",
					Destination: &fComp,
				},
			},
			Action: print,
		},
		{
			Name:  "update",
			Usage: "Update Task",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "t",
					Usage: "Task Title",
				},
				cli.IntFlag{
					Name:  "d",
					Usage: "Deadline num from today",
					Value: -1,
				},
				cli.IntFlag{
					Name:  "i",
					Usage: "Task id",
					Value: -1,
				},
			},
			Action: update,
		},
		{
			Name:   "post",
			Usage:  "Slack post",
			Action: post,
		},
	}
	app.Run(os.Args)
}

// readTask return struct Tasks
// read 'json' file
func readTasks(file string) (Tasks, error) {
	fp, err := os.Open(file)
	if err != nil {
		return Tasks{}, err
	}
	defer fp.Close()

	taskjson, err := ioutil.ReadAll(fp)
	if err != nil {
		return Tasks{}, err
	}

	var tasks Tasks
	err = json.Unmarshal(taskjson, &tasks)
	return tasks, err
}

// writeTask
// write task to 'json' file
func writeTasks(tasks Tasks) error {
	fp, err := os.Create(file)
	if err != nil {
		return err
	}
	defer fp.Close()

	taskjson, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(fp)
	_, err = writer.Write(taskjson)
	if err != nil {
		return err
	}
	return writer.Flush()
}

// printOneTask
// print a task with status
func printOneTask(id int, title string, deadline string) {
	fmt.Println(colorString(cyan, "id:    ") + strconv.Itoa(id))
	fmt.Println(colorString(cyan, "title: ") + title)
	fmt.Println(colorString(cyan, "date:  ") + deadline)
}

// colorString retur color string
// add color code
func colorString(color string, v string) string {
	return color + v + reset
}

func generateDate(now time.Time, plusDays int, layout string) string {
	day := now.AddDate(0, 0, plusDays)
	return day.Format(layout)
}

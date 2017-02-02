package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"time"

	"github.com/mattn/go-runewidth"
	"github.com/urfave/cli"
)

var cyan = "\u001b[36m"
var red = "\u001b[31m"
var blue = "\u001b[34m"
var reset = "\u001b[0m"

var taskfile = "./tasks/task.json"

var (
	idNum   = 3
	taskNum = 35
	dateNum = 15
)

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	DeadLine string `json:"dead_line"`
}

func main() {
	// Flag
	var fTask string
	var fDate string
	var fId int

	app := cli.NewApp()

	// App info
	app.Name = "gtask"
	app.Usage = "Task management"
	app.Version = "0.1"

	tasks, err := readTask(taskfile)
	if err != nil {
		log.Fatal(err)
	}
	app.Commands = []cli.Command{
		{
			Name:  "in",
			Usage: "add Task",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "t",
					Value:       "Task",
					Destination: &fTask,
				},
				cli.StringFlag{
					Name:        "d",
					Value:       genareteDate(3, "2006/01/02"),
					Destination: &fDate,
				},
			},
			Action: func(c *cli.Context) error {
				appendTask(&tasks, fTask, fDate)

				writeTask(tasks)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "del",
			Usage: "Delete Task",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "i",
					Usage:       "id",
					Destination: &fId,
				},
			},
			Action: func(c *cli.Context) error {
				if err = rmTask(fId, tasks.Tasks); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "print",
			Usage: "Print Task",
			Action: func(c *cli.Context) error {
				printTasks(tasks)
				return nil
			},
		},
	}
	app.Run(os.Args)
}

func readTask(file string) (Tasks, error) {
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

func appendTask(tasks *Tasks, title string, deadline string) {
	var id int
	if taskSlice := tasks.Tasks; len(taskSlice) != 0 {
		id = tasks.Tasks[len(tasks.Tasks)-1].Id + 1
	} else {
		id = 1
	}
	task := Task{
		Id:       id,
		Title:    title,
		DeadLine: deadline,
	}

	tasks.Tasks = append(tasks.Tasks, task)
	fmt.Println(colorString(blue, "-- Append --"))
	printOneTask(id, title, deadline)
}

func writeTask(tasks Tasks) error {
	fp, err := os.Create(taskfile)
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

func printTasks(tasks Tasks) {
	printLine(idNum + taskNum + dateNum + 4)
	fmt.Print("|" + cyan + truncateFillRight("id", idNum) + reset + "|")
	fmt.Print(cyan + truncateFillRight("task", taskNum) + reset + "|")
	fmt.Print(cyan + truncateFillRight("date", dateNum) + reset)
	fmt.Println("|")
	printLine(idNum + taskNum + dateNum + 4)
	for _, v := range tasks.Tasks {
		fmt.Print("|" + truncateFillRight(strconv.Itoa(v.Id), idNum))
		fmt.Print("|" + truncateFillRight(v.Title, taskNum))
		fmt.Print("|" + truncateFillRight(v.DeadLine, dateNum))
		fmt.Println("|")
	}
	printLine(idNum + taskNum + dateNum + 4)
}

func printLine(num int) {
	for i := 0; i < num; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func colorString(color string, v string) string {
	return color + v + reset
}

func rmTask(id int, tasks []Task) error {
	newTasks := make([]Task, 0, 0)
	for _, v := range tasks {
		if id != v.Id {
			newTasks = append(newTasks, v)
			continue
		}
		fmt.Println(colorString(red, "-- REMOVE --"))
		printOneTask(v.Id, v.Title, v.DeadLine)
	}
	if len(tasks) == len(newTasks) {
		return errors.New("Not found id: " + strconv.Itoa(id))
	}
	return writeTask(Tasks{newTasks})
}

func printOneTask(id int, title string, deadline string) {
	fmt.Println(colorString(cyan, "id:    ") + strconv.Itoa(id))
	fmt.Println(colorString(cyan, "title: ") + title)
	fmt.Println(colorString(cyan, "date:  ") + deadline)
}

func truncateFillRight(s string, w int) string {
	s = runewidth.Truncate(s, w, "..")
	return runewidth.FillRight(s, w)
}

func genareteDate(plusDays int, layout string) string {
	now := time.Now().AddDate(0, 0, plusDays)
	return now.Format(layout)
}

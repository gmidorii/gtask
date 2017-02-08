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

	"github.com/mattn/go-runewidth"
	"github.com/urfave/cli"
)

// console color
const cyan = "\u001b[36m"
const red = "\u001b[31m"
const blue = "\u001b[34m"
const reset = "\u001b[0m"

// Date Layout
const layout = "2006/01/02"

// Completed
const comp = "o"
const notComp = "-"

var taskfile = "./tasks/task.json"

var (
	idNum        = 3
	taskNum      = 35
	dateNum      = 15
	completedNum = 3
)

var lineNum = idNum + taskNum + dateNum + completedNum + 5

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

	tasks, err := ReadTask(taskfile)
	if err != nil {
		log.Fatal(err)
	}
	app.Commands = []cli.Command{
		{
			Name:  "in",
			Usage: "add Task",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "t",
					Value: "Task",
				},
				cli.IntFlag{
					Name:  "d",
					Value: -1,
				},
			},
			Action: write,
		},
		{
			Name:  "fi",
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
			Name:  "p",
			Usage: "Print Task",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "c",
					Usage:       "completed flag",
					Destination: &fComp,
				},
			},
			Action: func(c *cli.Context) error {
				printTasks(tasks, fComp)
				return nil
			},
		},
	}
	app.Run(os.Args)
}

// ReadTask return struct Tasks
// read 'json' file
func ReadTask(file string) (Tasks, error) {
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

// WriteTask
// write task to 'json' file
func WriteTask(tasks Tasks) error {
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

// PrintOneTask
// print a task with status
func PrintOneTask(id int, title string, deadline string) {
	fmt.Println(colorString(cyan, "id:    ") + strconv.Itoa(id))
	fmt.Println(colorString(cyan, "title: ") + title)
	fmt.Println(colorString(cyan, "date:  ") + deadline)
}

func printTasks(tasks Tasks, completed bool) {
	printLine(lineNum)
	fmt.Print("|" + cyan + truncateFillRight("id", idNum) + reset + "|")
	fmt.Print(cyan + truncateFillRight("title", taskNum) + reset + "|")
	fmt.Print(cyan + truncateFillRight("date", dateNum) + reset + "|")
	fmt.Print(cyan + truncateFillRight("c", completedNum) + reset)
	fmt.Println("|")
	printLine(lineNum)
	for _, v := range tasks.Tasks {
		if completed == false && v.Completed == true {
			continue
		}
		fmt.Print("|" + truncateFillRight(strconv.Itoa(v.Id), idNum))
		fmt.Print("|" + truncateFillRight(v.Title, taskNum))
		fmt.Print("|" + truncateFillRight(v.DeadLine, dateNum))
		if v.Completed {
			fmt.Print("|" + truncateFillRight(comp, completedNum))
		} else {
			fmt.Print("|" + truncateFillRight(notComp, completedNum))
		}
		fmt.Println("|")
	}
	printLine(lineNum)
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

func truncateFillRight(s string, w int) string {
	s = runewidth.Truncate(s, w, "..")
	return runewidth.FillRight(s, w)
}

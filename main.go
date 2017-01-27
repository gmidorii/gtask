package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"errors"
)

var cyan = "\u001b[36m"
var red = "\u001b[31m"
var blue = "\u001b[34m"
var reset = "\u001b[0m"

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	DeadLine string `json:"dead_line"`
}

func main() {
	// set flag
	var (
		task string
		date string
		id   int
		i    bool
		d    bool
	)
	flag.StringVar(&task, "task", "Task", "task title")
	flag.StringVar(&date, "date", "2017-01-24", "deadline")
	flag.IntVar(&id, "id", 1, "id")
	flag.BoolVar(&i, "i", false, "insert")
	flag.BoolVar(&d, "d", false, "delete")
	flag.Parse()

	tasks, err := readTask()
	if err != nil {
		log.Fatal(err)
	}

	if i {
		appendTask(&tasks, task, date)

		writeTask(tasks)
		if err != nil {
			log.Fatal(err)
		}
	} else if d {
		if err = rmTask(id, tasks.Tasks); err != nil {
			log.Fatal(err)
		}
	} else {
		printTasks(tasks)
	}
}

func readTask() (Tasks, error) {
	fp, err := os.Open("./tasks/task.json")
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
	fp, err := os.Create("./tasks/task.json")
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
	fmt.Println("-------------")
	fmt.Print("|")
	fmt.Print(cyan + "id" + reset)
	fmt.Print("|")
	fmt.Print(cyan + "task" + reset)
	fmt.Print("|")
	fmt.Print(cyan + "date" + reset)
	fmt.Println("|")
	fmt.Println("-------------")
	for _, v := range tasks.Tasks {
		fmt.Println("|" + strconv.Itoa(v.Id) + "|" + v.Title + " | " + v.DeadLine + "|")
	}
	fmt.Println("-------------")
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

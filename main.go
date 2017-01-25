package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var cyan = "\u001b[36m"
var reset = "\u001b[0m"

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Title    string `json:"title"`
	DeadLine string `json:"dead_line"`
}

func main() {
	// set flag
	var t string
	var d string
	var i bool
	flag.StringVar(&t, "t", "Task", "task title")
	flag.StringVar(&d, "d", "2017-01-24", "deadline")
	flag.BoolVar(&i, "i", false, "insert")
	flag.Parse()

	tasks, err := readTask()
	if err != nil {
		log.Fatal(err)
	}

	if i {
		appendTask(&tasks, t, d)

		writeTask(tasks)
		if err != nil {
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
	task := Task{
		Title:    title,
		DeadLine: deadline,
	}

	tasks.Tasks = append(tasks.Tasks, task)
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
	fmt.Print(cyan + "task" + reset)
	fmt.Print("|")
	fmt.Print(cyan + "date" + reset)
	fmt.Println("|")
	fmt.Println("-------------")
	for _, v := range tasks.Tasks {
		fmt.Println("|" + v.Title + " | " + v.DeadLine + "|")
	}
	fmt.Println("-------------")
}

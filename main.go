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
)

var cyan = "\u001b[36m"
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
		t string
		d string
		i bool
	)
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
	fmt.Println(colorString(cyan, "id:    ") + strconv.Itoa(id))
	fmt.Println(colorString(cyan, "title: ") + title)
	fmt.Println(colorString(cyan, "date:  ") + deadline)
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

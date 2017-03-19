package main

import (
	"testing"
	"time"
	"strconv"
)

func Test_generateDate(t *testing.T) {
	now := time.Date(2017, 3, 18, 0,0, 0, 0, time.UTC)
	layout := "2006/01/02"
	date := generateDate(now, 3, layout)
	if date != "2017/03/21" {
		t.Error("Fail to craate Date")
		t.Log(date)
	}
}

func Test_colorString(t *testing.T) {
	cyan := "\u001b[36m"
	reset := "\u001b[0m"

	v := "string"
	str := colorString(cyan, v)

	if str != cyan + v + reset {
		t.Error("Fail to create string")
		t.Log(v)
	}
}

func Test_readTask(t *testing.T) {
	file := "./test/read.json"
	tasks, err := readTasks(file)
	if err != nil {
		t.Error("Fail to run func")
		t.Log(err)
	}
	if len(tasks.Tasks) != 1 {
		t.Error("Unexpected read number of task: " + strconv.Itoa(len(tasks.Tasks)))
	}
	task := tasks.Tasks[0]
	if task.Id != 1 || task.Completed != false || task.DeadLine != "2017/03/19" || task.Title != "Hoge" {
		t.Error("Unexpected task value")
		t.Log(task)
	}
}
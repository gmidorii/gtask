package main

import (
	"fmt"
	"strconv"

	"github.com/mattn/go-runewidth"
	"github.com/urfave/cli"
)

func print(c *cli.Context) error {
	fComp := c.Bool("c")
	tasks, err := ReadTask(taskfile)
	if err != nil {
		return err
	}
	printTasks(tasks, fComp)
	return nil
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

func truncateFillRight(s string, w int) string {
	s = runewidth.Truncate(s, w, "..")
	return runewidth.FillRight(s, w)
}

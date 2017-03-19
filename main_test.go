package main

import (
	"testing"
	"time"
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
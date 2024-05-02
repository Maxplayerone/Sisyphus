package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func IntToString(input int) string {
	i := 1
	for input >= i {
		i *= 10
	}
	i /= 10

	var b strings.Builder

	for i >= 1 {
		num := input / i % 10
		b.WriteRune(rune(num + 48))
		i /= 10
	}

	return b.String()
}

func ShowTasksForToday() {
	year, month, day := time.Now().Date()
	fmt.Println("====================================")
	fmt.Println("Hello, today is " + IntToString(day) + " of " + month.String() + " " + IntToString(year))
	fmt.Println("====================================")
}

type Date struct {
	day   int
	month string
	year  int
}

func (date Date) GetMonthNumericString() string {
	switch date.month {
	case "January":
		return "01"
	case "February":
		return "02"
	case "March":
		return "03"
	case "April":
		return "04"
	case "May":
		return "05"
	case "June":
		return "06"
	case "July":
		return "07"
	case "August":
		return "08"
	case "September":
		return "09"
	case "October":
		return "10"
	case "November":
		return "11"
	case "December":
		return "12"
	}
	return "error"
}

func (date Date) ToString() string {
	var b strings.Builder
	b.WriteString(IntToString((date.day)))
	b.WriteRune('.')
	b.WriteString(date.GetMonthNumericString())
	b.WriteRune('.')
	b.WriteString(IntToString((date.year)))
	b.WriteRune('\n')
	return b.String()
}

func NewDate(year int, month time.Month, day int) Date {
	return Date{
		day:   day,
		month: month.String(),
		year:  year,
	}
}

func Clear() {
	//copying data to the new file
	data, err := ioutil.ReadFile("tasks.sf")
	if err != nil {
		panic(err)
	}
	if len(data) > 0 {
		err = ioutil.WriteFile("tasks_backup.sf", data, 0644)
		if err != nil {
			panic(err)
		}
	}

	//clearing the old file

	err = os.Remove("tasks.sf")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("tasks.sf")
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func Write(task string, date Date) {
	f, err := os.OpenFile("tasks.sf", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(date.ToString()); err != nil {
		panic(err)
	}
	if _, err = f.WriteString(task); err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args[1:]
	date := NewDate(time.Now().Date())
	if len(args) > 0 && args[0] == "today" {
		ShowTasksForToday()
	} else if len(args) > 1 && args[0] == "write" {
		var new_task string
		for _, arg := range args[1:] {
			new_task += arg + " "
		}
		new_task = new_task[0 : len(new_task)-1]

		Write(new_task, date)
	} else if len(args) > 0 && args[0] == "clear" {
		Clear()
	}
}

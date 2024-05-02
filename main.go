package main

import (
	"fmt"
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
	return b.String()
}

func NewDate(year int, month time.Month, day int) Date {
	return Date{
		day:   day,
		month: month.String(),
		year:  year,
	}
}

func main() {
	args := os.Args[1:]
	date := NewDate(time.Now().Date())
	if len(args) > 0 && args[0] == "today" {
		ShowTasksForToday()
	}

	file, err := os.Create("test.sf")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.Write([]byte(date.ToString()))
	if err != nil {
		fmt.Println(err)
	}
	_, _ = file.Write([]byte("\n"))
}

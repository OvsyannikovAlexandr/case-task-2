package main

import (
	"fmt"
	"time"
)

type BirthDate struct {
	Day   int
	Month int
	Year  int
}

func main() {
	date := inputBirthDate()

	t := time.Date(date.Year, time.Month(date.Month), date.Day, 0, 0, 0, 0, time.Local)
	weekday := getWeekDay(t)
	fmt.Printf("Weekday of your birthday: %s\n", weekday)

	if isLeapYear(t.Year()) {
		fmt.Printf("Your birthday year is leap\n")
	} else {
		fmt.Printf("Your birthday year is not leap\n")
	}

	age := calculateAge(t)
	fmt.Printf("Your age: %d\n", age)
}

func inputBirthDate() BirthDate {
	var d, m, y int
	fmt.Println("Enter day (1-31): ")
	fmt.Scan(&d)
	fmt.Println("Enter month (1-12): ")
	fmt.Scan(&m)
	fmt.Println("Enter year (example 1998): ")
	fmt.Scan(&y)
	return BirthDate{d, m, y}
}

func getWeekDay(t time.Time) string {
	days := map[time.Weekday]string{
		time.Monday:    "Monday",
		time.Tuesday:   "Tuesday",
		time.Wednesday: "Wednesday",
		time.Thursday:  "Thursday",
		time.Friday:    "Friday",
		time.Saturday:  "Saturday",
		time.Sunday:    "Sunday",
	}
	return days[t.Weekday()]
}

func isLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

func calculateAge(birthDate time.Time) int {
	now := time.Now()
	years := now.Year() - birthDate.Year()

	if now.Month() < birthDate.Month() || now.Month() == birthDate.Month() && now.Day() < birthDate.Day() {
		years--
	}
	return years
}

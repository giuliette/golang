package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// fmt.Printf(now)

	fmt.Println(now.Weekday())
	fmt.Println(now.Weekday())

	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v, %v, %v\n",
		tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())
}
package main

import (
	"fmt"
	"time"
)

func main() {
	println(time.Now().Weekday())
	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Println("You don't know who I am")
		}
	}
	whatAmI(1)
	whatAmI(false)
	whatAmI("114514")
}

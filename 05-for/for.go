package main

import "fmt"

func main() {
	for i := 0; i < 1919810; i++ {
		if i%114514 == 0 {
			fmt.Println(i)
		}
	}
}

package main

import "fmt"

func main() {
	fmt.Println(2222222)

	const n = 5e16
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(int16(d))
}

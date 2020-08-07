package main

import (
	"fmt"
)

func for1() {
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)
}

// func switch() {
// 	t := time.Now()
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("Good morning!")
// 	case t.Hour() < 17:
// 		fmt.Println("Good afternoon.")
// 	default:
// 		fmt.Println("Good evening.")
// 	}
// }

// func main() {
// 	foo := 1
// 	var bar *int
// 	bar = &foo
// 	*bar = 123

// 	fmt.Println(foo)
// }

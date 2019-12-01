package main

import "fmt"

func main() {
	// x := 100
	// y := 10 if x >= 10 else x
	// fmt.Printf("x: %d y: %d\n", x, y)
	x := 10
	s := "A" if x >= 10 else "B"
	fmt.Printf("x: %d s: %s\n", x, s)
}

// s := "A" if x >= 10 else "B"
// ----------------------------
// if x >= 10 {
// 	s := "A"
// } else {
// 	s := "B"
// }
package main

import "fmt"

func main() {
	y := 0
	for x := 1; x <= 10; x++ {
		let y = 1 if x <= 5 else 2
		fmt.Printf("(x,y)=(%d,%d)\n", x, y)
	}
}
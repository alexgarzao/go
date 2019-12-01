package main

import "fmt"

func main() {
	for x := 1; x <= 10; x++ {
		y := 0
		let y = 1 if x <= 5 else 2
		fmt.Printf("(x,y)=(%d,%d)\n", x, y)
	}
}

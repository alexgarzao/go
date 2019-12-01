package main

import "fmt"

func main() {
	fmt.Printf("\nTest 1\n")
	production := false
	port := 80 if production else 8080
	fmt.Printf("Production: %s port: %d\n", production, port)

	fmt.Printf("\nTest 2\n")
	production = true
	port = 80 if production else 8080
	fmt.Printf("Production: %s port: %d\n", production, port)

	fmt.Printf("\nTest 3\n")
	y := 0
	for x := 1; x <= 10; x++ {
		y = 1 if x <= 5 else 2
		fmt.Printf("(x,y)=(%d,%d)\n", x, y)
	}

	fmt.Printf("\nTest 3\n")
	y = 0
	for x := 1; x <= 10; x++ {
		fmt.Printf("(x,y)=(%d,%d)\n", x, 5 if x <= 5 else 10)
	}
}

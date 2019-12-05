package main

import "fmt"

func main() {
	production := false
	port := 0
	let port = 80 if production else 8080
	fmt.Printf("Production: %v port: %d\n", production, port)

	production = true
	let port = 80 if production else 8080
	fmt.Printf("Production: %v port: %d\n", production, port)
}
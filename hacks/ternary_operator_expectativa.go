package main

import "fmt"

func main() {
	production := false
	port := 80 if production else 8080 
	fmt.Printf("Production: %v port: %d\n", production, port)

	production = true
	port = 80 if production else 8080
	fmt.Printf("Production: %v port: %d\n", production, port)
}
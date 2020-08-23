package main

import (
	"fmt"
)

func main() {
	fmt.Println(hellowworld("Juan"))
}

func hellowworld(name string) string {
	response := "Hello " + name
	return response
}

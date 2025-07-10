package main

import "fmt"

// Struct - basically creates a new custom data type
type Person struct {
	Name string
	Age  int
}

func main() {
	var abhinav Person = Person{
		Name: "Abhinav",
		Age:  32,
	}

	fmt.Println("Person", abhinav)
}

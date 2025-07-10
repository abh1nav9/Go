package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// while loop
	j := 0
	for j < 5 {
		fmt.Println("While loop iteration:", j)
		j++
	}
}

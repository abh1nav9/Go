package main

import (
	"errors"
	"fmt"
)

func main() {
	var numerator int = 11
	var denominator int = 2

	var result, remainder, err = calculation(numerator, denominator)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Please provide a valid denominator.")
		return
	}

	fmt.Printf("The result is %v and the remainder is %v\n", result, remainder)

}

func calculation(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("denominator cannot be zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, nil
}

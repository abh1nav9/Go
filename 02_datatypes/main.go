package main
import "fmt"

func main(){

    // we all know different between var and const, thanks to JavaScript.

	var number int = 01; // int, int8 int16, int32, int64 : // stores both positive and negative values
	fmt.Println("The number is:", number)

	var unumber uint = 01; // uint, uint8, uint16, uint32, uint64 : stores only positive values
	fmt.Println("The unsigned number is:", unumber)

	var floatNumber float32 = 1.5; // float32, float64 : stores decimal values
	fmt.Println("The float number is:", floatNumber)

	var str string = "Hello, World!"; // string : stores text
	fmt.Println("The string is:", str)

	var boolean bool = true; // bool : stores true or false
	fmt.Println("The boolean value is:", boolean)

	var rune rune = 'A'; // rune : represents a Unicode code point, equivalent to int32
	fmt.Println("The rune is:", rune)

	var complexNumber complex64 = 1 + 2i; // complex64, complex128 : stores complex numbers
	fmt.Println("The complex number is:", complexNumber)

	var byteValue byte = 65; // byte : alias for uint8, used to represent ASCII values
	fmt.Println("The byte value is:", byteValue)

	var array [3]int = [3]int{1, 2, 3}; // array : fixed-size sequence of elements of the same type
	fmt.Println("The array is:", array)

	var slice []int = []int{1, 2, 3}; // slice : dynamically-sized sequence of elements of the same type
	fmt.Println("The slice is:", slice)

	var mapExample map[string]int = map[string]int{"one": 1, "two": 2}; // map : collection of key-value pairs
	fmt.Println("The map is:", mapExample)

	var channelExample chan int = make(chan int); // channel : used for communication between goroutines
	fmt.Println("The channel is created:", channelExample)
}
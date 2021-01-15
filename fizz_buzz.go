package main

// importing fmt package
import "fmt"

func main() {
	// := variable type is int and sets it to 1
	// i is less or equal to 100
	// increment i by 1
	for i := 1; i <= 100; i++ {
		// setting result as empty string
		result := ""
		// if i is divisable by 3
		if i%3 == 0 {
			result += "Fizz"
		}
		// if i is divisable
		if i%5 == 0 {
			result += "Buzz"
		}
		// if result is not an empty string adding both Fizz + Buzz
		if result != "" {
			fmt.Println(result)
			continue
		}
		// print all non Fizz Buzz numbers
		fmt.Println(i)
	}
}

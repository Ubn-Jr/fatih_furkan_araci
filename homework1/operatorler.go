package main

import "fmt"

func main() {

	fibonacci(0, 1)
}

func fibonacci(a int, b int) {
	var sum int

	for a := 0; a <= 10; a++ {
		sum = a + b
		sum = sum + a

		fmt.Println(a, " + ", a+1, " = ", sum)
	}

}

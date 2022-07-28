package main

import "fmt"

func main() {

	fibonacci(0, 1)
}

func fibonacci(a int, b int) {

	var sum int

	for i := 0; i <= 20; i++ {
		sum = a + b
		a = b
		b = sum

		fmt.Println(i, " = ", "Fibonacci Sayisi : ", sum)
	}

}

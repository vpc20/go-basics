package main

import "fmt"

func fibo(n int) int {
	if n < 2 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}
func main() {
	fmt.Println(fibo(0))
	fmt.Println(fibo(1))
	fmt.Println(fibo(2))
	fmt.Println(fibo(3))
	fmt.Println(fibo(4))
	fmt.Println(fibo(5))
}

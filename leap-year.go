package main

import "fmt"

func isLeapYear(yr int) bool {
	return yr%4 == 0 && yr%100 != 0 || yr%400 == 0
}

func main() {
	fmt.Println(isLeapYear(2000))
	fmt.Println(isLeapYear(2004))
	fmt.Println(isLeapYear(2100))
	fmt.Println(isLeapYear(2200))
	fmt.Println(isLeapYear(2300))
	fmt.Println(isLeapYear(2400))

	var x int
	x = 5 + 10 + 15
	fmt.Println(x)

	y := 10
	fmt.Println(y)

}

package main

import (
	"fmt"
)

// func isLeapYear(yr int) bool {
// 	return yr%4 == 0 && yr%100 != 0 || yr%400 == 0
// }

func main() {
	// board := [][]string{
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// }
	// fmt.Println(board)

	row := 3
	col := 5
	data2d := make([][]int, row)
	// for i := 0; i < row; i++ {
	// 	data2d[i] = make([]int, col)
	// }
	for i := range data2d {
		data2d[i] = make([]int, col)
	}
	fmt.Println(data2d)

	data2dx := make([][]int, row)
	// var data2dx [row][col]int
	rdata := make([]int, col)
	fmt.Println(rdata)
	// for i := 0; i < row; i++ {
	// 	append(data2dx, rdata)
	// }
	fmt.Println(data2dx)

}

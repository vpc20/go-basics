package main

import (
	"fmt"
)

func main() {
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
}

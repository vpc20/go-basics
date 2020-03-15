package main

import (
	"fmt"
	// "math"
)

// Sqrt comment
func Sqrt(x float64) float64 {
	var z, prevz float64 = 1.0, 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(prevz, z, prevz-z)
		// if math.Abs(z) == math.Abs(prevz) {
		// 	break
		// } else {
		// 	prevz = z
		// }
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

package main

import "fmt"

// Vertex structure
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var m2 = map[int]string{1: "one", 2: "two", 3: "three"}

func main() {
	fmt.Println(m)
	fmt.Println(m2)
}

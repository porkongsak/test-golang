package main

import "fmt"

func main() {

	var a = []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	  }
	  
	min, max := findMinAndMax(a)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)

}

func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
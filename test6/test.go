package main

import "fmt"

func main(){
	PrintStarTriangle(6)
}

func PrintStarTriangle(number int) {
	
	for i := 1; i <= number; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
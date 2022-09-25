package main 

import "fmt"

func main () {
	var myWorld string = "abcdefghijklmnopqrstuvwxyz"
	for i,v := range myWorld {
		fmt.Println("Index = ", i ,",  value =", string(v))
	}

}
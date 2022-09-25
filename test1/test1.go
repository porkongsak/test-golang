package main

import "fmt"



func main(){

    var s []int
	fmt.Println("Example -: if condition .")
	
	for i :=1 ; i <= 100; i++ {
		if i%3 ==0 {
			s = append(s, i)	
		} 
	}
	fmt.Printf("%d \n",s)
	fmt.Println("len:", len(s))
	
}


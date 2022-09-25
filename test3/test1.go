package main 

import (
	"fmt"
	"strconv"
)


func main(){
	counter(1000)
}

func counter(t int){
	var s []int
	for i := 1; i < t ; i++ {
		//fmt.Printf("%d\n",i)
		str:= strconv.Itoa(i)
  		//fmt.Println( str)
		
		for j := 0 ; j < len(str); j++ {

			//fmt.Println(str[j]) 
			if str[j] == '9' {
				s = append(s, i)
			}
		}	
	}
	fmt.Println(s)
	fmt.Println(len(s))
}


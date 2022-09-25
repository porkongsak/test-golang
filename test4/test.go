package main

import "fmt"

func BubbleSort(array[] int) ([]int) {

   for i:=0; i< len(array)-1; i++ {     
	  
	fmt.Println(i)

      for j:=0; j < len(array)-i-1; j++ {   

         if (array[j] > array[j+1]) { 
			

            array[j], array[j+1] = array[j+1], array[j] 

         }
      }
   }
   return array
}

func main() {
   array:= []int{
	48,96,86,68,
	57,82,63,70,
	37,34,83,27,
	19,97, 9,17,
	49,199,289,2,
	15,67,125,112,
  }
  fmt.Println(array[0])
  
   fmt.Println(BubbleSort(array))
}
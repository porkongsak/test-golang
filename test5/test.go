package main


import "fmt"

type employee struct {
	name     string
	salary   int
	position string
}

func main(){
	var name int

	 people := []employee{
		{name: "John", salary: 100 , position : "Back-end"},
		{name: "Jo", salary: 100 , position : "Back-end2"},
		{name: "Jokkkk", salary: 200 , position : "front-end"},
	 }


	 fmt.Scanf("%d", &name)


	 switch name {

	 case 1 :
		for _, item := range people{
			fmt.Println(item)
		 }

	case 2:
			fmt.Println(people[1])
	default:
		fmt.Println("ไม่เจอข้อมูล")
	 }

	

		 

	
}


	

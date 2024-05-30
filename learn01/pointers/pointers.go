package main

import "fmt"

func main(){
	age := 29
	fmt.Println(age)

	var agePointer *int
	agePointer = &age
	fmt.Println(agePointer)
	fmt.Println(*agePointer)


	getAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int){
	// return *age - 18
	*age = *age - 18
}
package main

import "fmt"

func main(){
	var age int = 10
	var fage float64 = float64(age)
	fmt.Printf("Age: %f \n",fage)
	
	//Float to integer

	var experience float64 = 2.5
	var experience_int int = int(experience)
	fmt.Printf("%v",experience_int)

	//constant
	const PI float64 = 2.02
	fmt.Println(PI)
}
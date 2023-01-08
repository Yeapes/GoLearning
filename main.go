package main

import "fmt"

func main()  {
	// fmt.Println("Hello world")

	// name:="Yeapes"
	// fmt.Println(name)

	// var name string
	// var b bool

	// fmt.Println("Enter your input")

	// fmt.Scanf("%s %d",&name, &b);

	// fmt.Println("Your name is ",name)
	// fmt.Println("You are  ",b)

	// count, error:=fmt.Scanf("%s %t",&name, &b)

	// fmt.Println("Count: ", count)
	// fmt.Println("error", error)
	// fmt.Println("Your name is ", name)
	// fmt.Println("You are  ", b)


	var name string = "Yeapes"
	var age int = 23
	var experience float64 = 2.5
	var is_engineer = true 

	fmt.Printf("The coder name is: %v type of %T \n",name,name)
	fmt.Printf("The coder Age is: %v type of %T \n",age,age)
	fmt.Printf("The coder experience is: %v type of %T \n",experience,experience)
	fmt.Printf("The coder is_engineer is: %v type of %T \n",is_engineer,is_engineer)
}
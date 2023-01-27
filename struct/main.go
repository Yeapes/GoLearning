package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var person1 Person
	person1.name = "Yeapes"
	person1.age = 23
	person1.job = "Software Engineer"
	person1.salary = 0001
	fmt.Println("Person 1 Name :", person1.name)
	fmt.Println("Person 1 Age :", person1.age)
}

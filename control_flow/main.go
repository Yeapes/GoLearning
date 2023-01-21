package main

import "fmt"

func main() {
	// age := 23;

	// if age >= 18{
	// 	fmt.Println("You have right access to the website")
	// }else if age <= 10{
	// 	fmt.Println("You need to parental access")
	// }else {
	// 	fmt.Println("It's not a correct value to check condition")
	// }

	day := "wednesday"
	switch day {
	case "monday":
		fmt.Println("monday")
	case "tuesday":
		fmt.Println("tuesday")
	case "wednesday":
		fmt.Println("wednesday")
		fallthrough
	case "thursday":
		fmt.Println("thursday")
		fallthrough
	case "friday":
		fmt.Println("friday")
	case "saturday", "sunday":
		fmt.Println("weekend")
	default:
		fmt.Println("default")
	}

}

package main

import "fmt"

func main(){


	//Bitwise and
	var x, y int = 12, 25
	z := x & y;

	fmt.Println(z)

	//Bitwise or example 12 in binary = 000001100 and 25 in binary 00011001
	//So it will be binary or operation with the help of bitwise or operation provide the result
	//29
	var a, b int = 12, 25
	c := a | b 
	fmt.Println(c)

	//Bitwise Xor example 12 in binary = 000001100 and 25 in binary 00011001
	//So it will be binary Xor operation with the help of bitwise or operation provide the result
	//21
	// result true if there is opostion of digit
	var m, n int = 12, 25
	o := m ^ n 
	fmt.Println(o)


	//Bitwise left shift operator  
	// 424
	// 
	// result true if there is opostion of digit
	//  
	var ls = 212 
	lsr := ls << 1 
	fmt.Println(lsr)

	//Bitwise left shift operator  
	// 424
	// 
	// result true if there is opostion of digit
	//  adding 2 bit right
	var rs = 212 
	rsr := rs >> 2
	fmt.Println(rsr)

}
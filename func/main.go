package main
import ("fmt")

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func myFunction2(x int, y int) int {
  return x + y
}

func main() {

	//we can omit using _
  a, b := myFunction(5, "Hello")
  fmt.Println(a, b)
fmt.Println(myFunction2(1,2));
}
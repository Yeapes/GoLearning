package main
import ("fmt")

func main() {
  var adj = [2]string{"big", "tasty"}
  var fruits = [3]string{"apple", "orange", "banana"}
  for i:=0; i < len(adj); i++ {
    for j:=0; j < len(fruits); j++ {
      fmt.Println(adj[i],fruits[j])
    }
  }


  fruits2 := [3]string{"apple", "orange", "banana"}

  //you can omit idx and value using _
  for idx, val := range fruits2 {
     fmt.Printf("%v\t%v\n", idx, val)
  }
}
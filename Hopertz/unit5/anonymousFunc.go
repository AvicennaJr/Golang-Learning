package main
import ("fmt");


func main() {
 var i func() int
 i = func() int {
 return 5
 }
 fmt.Println(i()) // 5
}

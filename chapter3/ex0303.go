package main

import "fmt"

func main() {
	var data []int
	x := []int{}
	fmt.Println(data == nil) //true
	fmt.Println(x == nil)    //false
	calc()
}

func calc() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, "i", "j", "k")
	x = append(x, "y")
	z = append(z, "y")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

}

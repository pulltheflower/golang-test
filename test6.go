package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	n := arr[:3]
	fmt.Println(n)
	fmt.Println(n.len())
}
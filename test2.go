package main

import "fmt"

func main() {
	{
		defer fmt.Println("defer runs")
		fmt.Println("block ends")
	}
	fmt.Println("main ends")
}


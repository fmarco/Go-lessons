package main

import "fmt"


//A function is declared by 'func' keyword

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(15, 27))
}

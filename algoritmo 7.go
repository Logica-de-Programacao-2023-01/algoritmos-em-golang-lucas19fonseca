package main

import "fmt"

func main() {
	var x int
	fmt.Print("qual o número")
	fmt.Scan(&x)
	antecessor := x - 1
	sucessor := x + 1
	fmt.Println("seu antecessor é", antecessor, "seu sucessor é", sucessor)
	
}

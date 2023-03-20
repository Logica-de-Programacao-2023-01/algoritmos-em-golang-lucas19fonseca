package main

import "fmt"

func main() {

	var salario float64
	fmt.Print("qual é o seu salário")
	fmt.Scan(&salario)

	aumento := salario * 0.15
	total := salario + aumento
	fmt.Println("seu salário com aumento é", total)
}

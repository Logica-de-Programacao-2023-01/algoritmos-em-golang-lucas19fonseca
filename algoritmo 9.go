package main

import "fmt"

func main() {
	var valorProduto float64
	fmt.Print("qual o valor do produto?")
	fmt.Scan(&valorProduto)
	desconto := valorProduto * 0.1
	resultado := valorProduto - desconto
	fmt.Println("o valo com descontos é de", resultado)

}

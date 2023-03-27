package main

import "fmt"

// um Array de inteiros com 5 elementos e imprima cada elemento em uma linha separada.
func main() {
	nome := [5]string{"joão", "maria", "josé", "Lucas", "Gabriel"}
	nome1 := nome[0]
	nome2 := nome[1]
	nome3 := nome[2]
	nome4 := nome[3]
	nome5 := nome[4]
	fmt.Println(nome1)
	fmt.Println(nome2)
	fmt.Println(nome3)
	fmt.Println(nome4)
	fmt.Println(nome5)
}

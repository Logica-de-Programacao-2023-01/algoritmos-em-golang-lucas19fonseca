package main

import "fmt"

func main() {
	var numeroDias int
	fmt.Print("quantos dias os funcinários trabalhou?")
	fmt.Scan(&numeroDias)
	var valorDiaria float64
	fmt.Print("qual o valor da diária do funcionário")
	fmt.Scan(&valorDiaria)
	salario := float64(numeroDias)
	fmt.Print("qual o valor da diária do funcionario?")
	fmt.Scan(&valorDiaria)
	salario := numeroDias * valorDiaria

}

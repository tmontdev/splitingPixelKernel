package main

import "fmt"

type Pessoa struct{
	nome string,
	sobrenome string,
	nomeCompleto string,
	idade int
}

func soma(numA, numB int)int{
	return numA + numB
}

func makeSobrenome(pessoa Pessoa)Pessoa{
	pessoa.nomeCompleto = nome+sobrenome
}

func main() {
	nomes := [3]string{"jose", "vanessa", "isis"}
	fmt.Println(soma(2, 2))
	for cont := 0; cont < 3; cont++ {
		fmt.Println(nomes[cont])
	}
}
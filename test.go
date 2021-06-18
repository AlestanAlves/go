package main

import "fmt"

func main() {
	nome := "Ale"
	fmt.Println("Eae", nome)
	
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair")

	var comand int 
	fmt.Scan(&comand)
	// fmt.Scanf("%d", &comand)

	fmt.Println("O endereço da var comand", &comand)
	fmt.Println("Comand:", comand)
}
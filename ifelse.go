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

	// if comand == 1 {
	// 	fmt.Println("Monitorando...")
	// }else if comand == 2{
	// 	fmt.Println("Exibindo logs")
	// }else if comand == 0{
	// 	fmt.Println("Eae")
	// }else {
	// 	fmt.Println("comand não registrado")
	// }

	switch comand {
		case 1:
			fmt.Println("Eae")
		case 2:
			fmt.Println("eae 2")
		case 0:
			fmt.Println("eae 0")
		default:
			fmt.Println("nao conheço esse comando")
	}
}


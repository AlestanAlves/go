package main

import "fmt"
import "os"
import "net/http"

func main() {

	intro()
	for {
		menu()
		comand := comand()

		switch comand {
			case 1:
				iniciarMon()
			case 2:
				fmt.Println("eae 2")
			case 0:
				fmt.Println("eae 0")
				os.Exit(0)
			default:
				fmt.Println("nao conheço esse comando")
				os.Exit(-1)
		}
	}

}

func nome() (string,int) {
	nome := "Ale"
	idade := 25
	return nome, idade
}

func intro(){
	nome := "Ale"
	fmt.Println("Eae", nome)
}

func comand() int {
	var comand int 
	fmt.Scan(&comand)
	// fmt.Scanf("%d", &comand)

	fmt.Println("O endereço da var comand", &comand)
	fmt.Println("Comand:", comand)

	return comand
}

func menu(){
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair")
}

func iniciarMon(){
	fmt.Println("Monitorando...")
	site := "http://a1f2f4249336240809598ff4288ee577-1174557885.us-east-1.elb.amazonaws.com"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "carregado com sucesso 200")
	}else {
		fmt.Println("Site:", site, "está com problem, Status code:", resp.StatusCode)
	}
}
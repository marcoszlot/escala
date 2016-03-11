package main

import "fmt"

func main() {
	fmt.Println("\n\nESCALA 1.O     by: Marcos Zlotnik\n\nDIGITE 1 PARA CONVERTER DE cm PARA m OU DIGITE 2 PARA CONVERTER DE cm PARA km.\nENTER:")
	var selecionando int
	fmt.Scanln(&selecionando)
	if selecionando == 1 {
		fmt.Println("\nMETROS\n\nDIGITE A MEDIDA EM cm:")
		var m int
		fmt.Scanln(&m)
		fmt.Println("ESCALA 1:", m / 100)
	} else if selecionando == 2 {
		fmt.Println("\n\nDIGITE A MEDIDA EM cm:")
		var km int
		fmt.Scanln(&km)
		fmt.Println("ESCALA 1:", km / 100000)

	} else {
		fmt.Println("ERROR")
	}
}
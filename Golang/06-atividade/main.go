package main

import "fmt"

func main() {
	herois := make(map[string]int)

	heroi := ""

	herois["Batman"] = 10
	herois["Flash"] = 7
	herois["Thor"] = 10
	herois["Homem de Ferro"] = 10
	herois["Homem Aranha"] = 3
	herois["Aquaman"] = 3
	herois["Capitão America"] = 2
	herois["Namor"] = 6
	herois["Doutor Estranho"] = 8
	herois["Mulher Maravilha"] = 10
	herois["Coringa"] = 6
	herois["Robin"] = 3
	herois["Capitã Marvel"] = 1
	herois["Loki"] = 2
	herois["Homem Formiga"] = 3
	herois["Doutor Destino"] = 2
	herois["Superman"] = 8

	fmt.Println("Digite um heroi:")
	fmt.Scanln(&heroi)
	// heroi := "Batman"

	if herois[heroi] >= 7 {
		fmt.Println("Fortissimo")
	} else if herois[heroi] >= 5 {
		fmt.Println("Apenas melhore")
	} else {
		fmt.Println("Nem Tente")
	}
}

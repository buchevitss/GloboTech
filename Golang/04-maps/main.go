package main

import "fmt"

func main() {
	languagelevel := make(map[string]int)

	languagelevel["Java"] = 10
	languagelevel["Python"] = 5

	fmt.Println(languagelevel)

	delete(languagelevel,"Java")

	//---//
	consoleAndGames := make(map[string][]string)
	consoleAndGames["Sonic"] = append(consoleAndGames["Sonic"],"Donkey Kong")

	//um slice dentro de um slice
}
package main

import "fmt"

func main() {
	favoriteFoods := [3]string{"Lasanha", "Pizza", "Carbonara"}

	fmt.Println("Minhas Comidas favoritas são")

	for index,food := range favoriteFoods{ fmt.Println(index, "-",food)
}
	for i:=0;i<len(favoriteFoods);i++{
		fmt.Println(i,"-", favoriteFoods[i])
	//while é um for com 1 argumento só

	count:=1
	for count  <= 5 {
		fmt.Println("O contador é",  count)
		count++
	}

	//for - controlar ele por uma regra que está dentro dele
	for{
		
	}
	}
}
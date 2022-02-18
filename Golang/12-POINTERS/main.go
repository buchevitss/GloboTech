package main

import "fmt"

type game struct{
	name string
	year uint
}

func (g game)changeName(newName string){
	g.name=newName
}

func main(){
	mario:=game{
		name:"Super Mario"
		year:1990,
	}
	mario.changeName("Super Mario World")
	fmt.Println(mario)
}
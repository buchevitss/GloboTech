package main

func main(){
	sendMessages("teste1","teste2")
}

func sendMessages(msgs...string)bool{
	for_,message := range msgs{
		fmt.Println("Enviando",message)

	}return true
}
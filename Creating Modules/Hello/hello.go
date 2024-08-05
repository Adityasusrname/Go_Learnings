package main

import(
	"fmt"

	"example.com/greetings"
)

func main(){
   message:=greetings.Hello("Aditya")
   fmt.Println(message)
}
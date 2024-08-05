package main

import(
	"fmt"
     "log"
	"example.com/greetings"
)

func main(){

   log.SetPrefix("greetings: ")
   log.SetFlags(0)

  // message,err:=greetings.Hello("Aditya")

  names:=[]string{"Aditya","Vidushi","Samridhi"}

  messages,err:=greetings.Hellos(names)

   if err != nil{
	log.Fatal(err)
   }

   fmt.Println(messages)

}
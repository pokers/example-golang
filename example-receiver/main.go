package main

import (
	"fmt"
	"example-receiver/receiver"
)

func main(){
	fmt.Printf("%s\n", receiver.GetUser())
	receiver.GetUsers()
	receiver.GetUsers2()

	val1 := receiver.User{
		Id: 10,
		Name: "JoongSu, Seo",
		Age: 22,
	}

	fmt.Println(val1)
}
package main

import (
	"fmt"
	"example-package/controller"
)

func main(){
	fmt.Printf("here is main\n")
	controller.ExController2()
	controller.Ex2()
	controller.Ex3()
	var result string = ""
	result = controller.PrintLn("Seo Jung Su")
	fmt.Printf("%s\n", result)
}
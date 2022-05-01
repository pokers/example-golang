package controller

import (
	"fmt"
)

func exController1(){
	fmt.Printf("Here is in exController1!\n")
}

func ExController2(){
	fmt.Printf("Here is in exController2!\n")
}

func printLine(message string) string {
	fmt.Printf("printLine > %s\n", message)
	return "[ " + message + " ]... printed"
}

func PrintLn(msg string) string {
	return printLine(msg)
}
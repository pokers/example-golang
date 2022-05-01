package main

import (
	"fmt"
)

func getPrintFunc() func(string) {
	var count uint32 = 0
	return func(msg string){
		count++
		fmt.Printf("[%d]> %s\n", count, msg)
	}
}
func main (){
	// value := map[string]string{"1":"a"}
	var value map[uint32]string = map[uint32]string{1:"a"}

	fmt.Println(value)
	fmt.Println(len(value))

	fnPrint := getPrintFunc()

	fnPrint("Print Function1")
	fnPrint("Print Function2")
	fnPrint("Print Function3")
}
package main

import (
	"fmt"
	"math"
	"rsc.io/quote"
)

func main(){
	fmt.Println("hello, golang!")
	bData := byte('B')
	var rData byte = '~'

	fmt.Printf("%c=%d, %c = %U\n", bData, bData, rData, rData)
	fmt.Printf("%f\n", math.Floor(1.8))
	fmt.Println(quote.Go())
}


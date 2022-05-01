package controller

import (
	"fmt"
)

// 같은 package내에서만 자유롭게 호출가능
func ex1(){
	fmt.Printf("Here is in controller-example ex1\n")
}

// 다른 package에서 호출 가능. 일종의 package의 interface 역활
func Ex2(){
	fmt.Printf("Here is in controller-example ex2\n")
}

func Ex3(){
	exController1()
	ex1()
}
package main

import (
	"fmt"
	"log"
	"net/http"
	"example-webApp/handler"
	"example-webApp/route"
)


func main(){
	path := []string{"view", "edit", "save"}
	
	handler.InitHandler(path)
	route.RegisterView()
	route.RegisterEdit()
	route.RegisterSave()
	fmt.Println("Starting server......")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
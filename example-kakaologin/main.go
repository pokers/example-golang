package main

import (
	"example-kakaologin/src/route"
	handler "example-kakaologin/src/viewHandlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	path := []string{"login", "auth"}

	handler.InitHandler(path)
	route.RegisterKakaoLogin()
	route.RegisterAsserts()
	route.RegisterKakaoAuth()
	fmt.Println("Starting server......")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

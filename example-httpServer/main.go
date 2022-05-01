package main

import (
	"fmt"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request){
	fmt.Println("index > request : ", req)
	fmt.Fprintf(res, "<h1>Example-http : index</h1>")
}

func healthCheck(res http.ResponseWriter, req *http.Request){
	fmt.Println("healthCheck > request : ", req)
	fmt.Fprintf(res, "<h1>Example-http : healthCheck</h1>")
}

func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/health-check", healthCheck)
	fmt.Println("Example-http : server starting...")
	http.ListenAndServe(":8080", nil)
}
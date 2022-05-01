package route

import (
	"net/http"
	"example-webApp/handler"
)

func RegisterView(){
	http.HandleFunc("/view/", handler.ViewHandler)
}
func RegisterEdit(){
	http.HandleFunc("/edit/", handler.EditHandler)
}
func RegisterSave(){
	http.HandleFunc("/save/", handler.SaveHandler)
}
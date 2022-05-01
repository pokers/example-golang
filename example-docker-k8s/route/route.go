package route

import (
	"net/http"
	"example-webApp/handler"
)

func RegisterView(){
	http.HandleFunc("/view/", handler.BuildHandler(2, handler.ViewHandler))
}
func RegisterEdit(){
	http.HandleFunc("/edit/", handler.BuildHandler(2, handler.EditHandler))
}
func RegisterSave(){
	http.HandleFunc("/save/", handler.BuildHandler(2, handler.SaveHandler))
}
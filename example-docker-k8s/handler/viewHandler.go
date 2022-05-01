package handler

import (
	"fmt"
	"net/http"
)

func ViewHandler(res http.ResponseWriter, req *http.Request, pageName string){
	page, err := loadPage(pageName)
	if err != nil {
		fmt.Println("Not Found : ", pageName, "end redirect to edit path")
		http.Redirect(res, req, "/edit/"+pageName, http.StatusFound)
		return
	}
	renderTemplate(res, "view", page)
}
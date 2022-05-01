package handler

import (
	"fmt"
	"net/http"
)

func ViewHandler(res http.ResponseWriter, req *http.Request){
	if pageName, err := getValidPath(2, res, req); err == nil {
		page, err := loadPage(pageName)
		if err != nil {
			fmt.Println("Not Found : ", pageName, "end redirect to edit path")
			http.Redirect(res, req, "/edit/"+pageName, http.StatusFound)
			return
		}
		renderTemplate(res, "view", page)	
	}
}
package handler

import (
	"net/http"
)

func EditHandler(res http.ResponseWriter, req *http.Request, pageName string){
	page, err := loadPage(pageName)
	if err != nil {
		page = &Page{Name: pageName}
	}
	renderTemplate(res, "edit", page)
}
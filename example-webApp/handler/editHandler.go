package handler

import (
	"net/http"
)

func EditHandler(res http.ResponseWriter, req *http.Request){
	if pageName, err := getValidPath(2, res, req); err == nil {
		page, err := loadPage(pageName)
		if err != nil {
			page = &Page{Name: pageName}
		}
		renderTemplate(res, "edit", page)	
	}
}
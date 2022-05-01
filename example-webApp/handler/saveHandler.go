package handler

import (
	"net/http"
)

func SaveHandler(res http.ResponseWriter, req *http.Request, pageName string){
	var body string = req.FormValue("body")
	var page *Page = &Page{Name: pageName, Body: []byte(body)}
	var err error = page.store()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(res, req, "/view/"+pageName, http.StatusFound)
}
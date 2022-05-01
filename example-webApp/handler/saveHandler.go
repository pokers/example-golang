package handler

import (
	"net/http"
)

func SaveHandler(res http.ResponseWriter, req *http.Request){
	if pageName, err := getValidPath(2, res, req); err == nil {
		var body string = req.FormValue("body")
		var page *Page = &Page{Name: pageName, Body: []byte(body)}
		var err error = page.store()
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(res, req, "/view/"+pageName, http.StatusFound)
	}
}
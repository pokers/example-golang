package handler

import (
	"fmt"
	"net/http"
	"html/template"
	"io/ioutil"
	"regexp"
	"strings"
	"errors"
)

var templates *template.Template
var validPath *regexp.Regexp

func renderTemplate(res http.ResponseWriter, templateName string, page * Page){
	fmt.Println("render Template : ", templateName)
	err := templates.ExecuteTemplate(res, templateName+".html", page)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func loadPage(pageName string) (*Page, error) {
	var fileName string = "./pageData/" + pageName + "PageData.txt"
	var body []byte
	var err error
	body, err = ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{Name: pageName, Body: body}, nil
}

func getValidPath(index int, res http.ResponseWriter, req *http.Request) (string, error){
	fmt.Println("path : ", req.URL.Path)
	var path []string = validPath.FindStringSubmatch(req.URL.Path)
	if path == nil || len(path) < index {
		http.NotFound(res, req)
		return "", errors.New("Invalid Path")
	}
	return path[index], nil
}

func initTemplates(){
	templates = template.Must(template.ParseFiles("./html/edit.html", "./html/view.html"))
}

func initValidPath(path []string){
	validPath = regexp.MustCompile("^/("+ strings.Join(path, "|") + ")/([a-zA-Z0-9]+)$")
}

func InitHandler(path []string){
	initTemplates()
	initValidPath(path)
}
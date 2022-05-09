package viewHandlers

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// Page Struct
type Page struct {
	Name string
	Body []byte
	ClientId string
	RedirectURL string
	Token string
}

func (page *Page) store() error {
	var fileName string = "./pageData/" + page.Name + "PageData.txt"
	return ioutil.WriteFile(fileName, page.Body, 0600)
}

//
var templates *template.Template
var validPath *regexp.Regexp

func renderTemplate(res http.ResponseWriter, templateName string, page *Page) {
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
	return &Page{Name: pageName, Body: body, ClientId: "", RedirectURL: "", Token: ""}, nil
}

func getValidPath(index int, res http.ResponseWriter, req *http.Request) (string, error) {
	fmt.Println("path : ", req.URL.Path)
	var path []string = validPath.FindStringSubmatch(req.URL.Path)
	if path == nil || len(path) < index {
		http.NotFound(res, req)
		return "", errors.New("Invalid Path")
	}
	return path[index], nil
}

func initTemplates() {
	templates = template.Must(template.ParseFiles("./html/kakaologin.html", "./html/kakaoauth.html"))
}

func initValidPath(path []string) {
	validPath = regexp.MustCompile("^/(" + strings.Join(path, "|") + ")/([a-zA-Z0-9]+)$")
}

////////////////////////////////////////////////////
// Public
func InitHandler(path []string) {
	initTemplates()
	initValidPath(path)
}

func BuildHandler(index int, fnProc func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("path : ", req.URL.Path)
		var path []string = validPath.FindStringSubmatch(req.URL.Path)
		if path == nil || len(path) < index {
			http.NotFound(res, req)
			return
		}
		fnProc(res, req, path[index])
	}
}

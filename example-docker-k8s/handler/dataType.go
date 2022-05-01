package handler

import (
	"io/ioutil"
)


// Page Struct
type Page struct{
	Name string
	Body []byte
}
func (page *Page) store() error {
	var fileName string = "./pageData/" + page.Name + "PageData.txt"
	return ioutil.WriteFile(fileName, page.Body, 0600)
}

//
package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

const templateDir="./templates"

type templateHandler struct{
	once sync.Once
	filename string
	compiledTemplate *template.Template

}


func (t *templateHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	t.once.Do(func(){
		t.compiledTemplate=template.Must(template.ParseFiles(filepath.Join(templateDir,t.filename)))
	})
	t.compiledTemplate.Execute(writer,request)
}

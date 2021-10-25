package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type loginHandler struct{

}


func (l *loginHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	segs:= strings.Split(request.URL.Path,"/")
	action:= segs[2]
	provider:=segs[3]
	switch action{
	case "login":
		log.Println("Todo handler for ",provider)
	default:
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(writer,"Authentication action %s not supported",action)
	}
}
package main

import (
	"log"
	"net/http"
)

type authHandler struct{
	next http.Handler
}

func (h *authHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_,err:= request.Cookie("auth")
	if err==http.ErrNoCookie{
		log.Println("No Cookie found")
		writer.Header().Set("Location","/login")
		writer.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	if err!=nil{
		log.Println("Internal Server Error")
		http.Error(writer,err.Error(),http.StatusInternalServerError)
		return
	}
	log.Println("Authentication has been done")
	h.next.ServeHTTP(writer,request)
}

func NewAuthHandler(next http.Handler) *authHandler{
	return &authHandler{next}
}





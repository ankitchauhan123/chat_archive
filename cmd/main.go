package main

import (
	"flag"
	"log"
	"net/http"
)

func main(){
	port:=flag.String("port",":8080","Port number of the app")
	flag.Parse()
	r := newRoom()
	ch:=NewAuthHandler(&templateHandler{filename: "chat.html"})
	login:=&templateHandler{filename: "login.html"}
	loginHandler:=&loginHandler{}

	http.Handle("/room", r)
	http.Handle("/login",login)
	http.Handle("/auth/", loginHandler)
	http.Handle("/chat",ch)


	//chat operation in background
	go r.run()
	//main function listens to connections
	log.Println("The port being used is :",*port)
	err:=http.ListenAndServe(*port,nil)
	log.Println("Done...",*port)

	if err!=nil{
		log.Fatal("Error in Listen & Serve",err)
	}

}




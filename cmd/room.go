package main

import (
	"chat/trace"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)


const (
	socketBufferSize = 1024
	messageBufferSize= 256
)

type room struct{
	//channel that receives the incoming messages from various clients
	forward chan []byte
	// join channel
	join chan *client
	//leave channel
	leave chan *client
	//active clients
	clients map[*client]bool
	//any tracer that implements this interface
	tracer trace.Tracer
}


func (r *room) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("New connection initiated...")
	//every connection to browser lands here and starts a go routine that would
	//write to the socket connection
	//and this function would block in read to read continuously from the socket
	socket,err:= upgrader.Upgrade(writer,request,nil)
	if err!=nil{
		log.Fatal("Error", err)
		return
	}
	client := & client{
		socket: socket,
		send: make(chan []byte, messageBufferSize),
		room: r,
	}
	r.join<- client
	defer func() {r.leave<-client}()

	go client.write()
	client.read()
}

func (r *room) run(){
	for {
		select {
			case client := <- r.join:
				r.clients[client]=true
				r.tracer.Trace("New client added...")

			case client := <- r.leave:
				delete(r.clients,client)
				close(client.send)
				r.tracer.Trace("Client added...")


		case msg := <- r.forward:
			r.tracer.Trace("Frowarding message...",msg)
			for client := range r.clients{
					client.send <- msg
				}
		}

	}
}

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize,WriteBufferSize: messageBufferSize}


func newRoom() *room {
	return &room{
		forward: make (chan []byte),
		join: make (chan *client),
		leave: make (chan *client),
		clients: make(map[*client]bool),
		tracer: trace.New(os.Stdout),
	}
}




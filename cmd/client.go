package main

import (
	"github.com/gorilla/websocket"
)

type client struct{
	//connection to communicate with browser
	socket *websocket.Conn
	//message would be sent to this channel via other users
	send chan []byte
	//room to which client is connected
	room *room
}


func (c *client) read(){
	defer c.socket.Close()
	for {
		_,msg,err:= c.socket.ReadMessage()
		if err==nil{
			c.room.forward <- msg
		}else{
			return
		}
	}
}

func (c *client) write(){
	defer c.socket.Close()
	for msg:= range c.send{
		err:= c.socket.WriteMessage(websocket.TextMessage,msg)
		if err!=nil{
			return
		}
	}
}

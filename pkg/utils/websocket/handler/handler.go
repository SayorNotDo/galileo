package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade error: %v", err)
	}
	defer func(c *websocket.Conn) {
		err := c.Close()
		if err != nil {
			log.Fatalf("close error: %v", err)
		}
	}(c)
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("read message error: %v", err)
			break
		}
		log.Printf("read message: %v", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Printf("write message error: %v", err)
			break
		}
	}
}

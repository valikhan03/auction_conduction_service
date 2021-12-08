package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{}

func (h *Handler) AuctionWSEndpoint(c *gin.Context){
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil{
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	messageReader(conn)
}


func messageReader(conn *websocket.Conn){
	defer conn.Close()
	for{

	}
}
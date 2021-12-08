package delivery

import (
	"log"
	"net/http"

	"auction_conduction/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func (h *Handler) AuctionWSEndpoint(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	messageReader(conn)
}

func (h *Handler) AuctionAdminWSEndpoint(c *gin.Context){
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil{
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	messageReader(conn)
}


func messageReader(conn *websocket.Conn) {
	defer conn.Close()

	var suggestion models.Suggestion
	for {
		err := conn.ReadJSON(&suggestion)
		if err != nil {
			log.Println(err)
		}

		err = conn.WriteJSON(suggestion)
		if err != nil {
			log.Println(suggestion)
		}
	}
}



package delivery

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"html/template"
)

func (h *Handler) AuctionPage(c *gin.Context){
	tmp, err := template.ParseFiles("client/auctionPage/auction.htm")
	if err != nil{
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	//http.ServeFile(c.Writer, c.Request, "client/auctionPage/static/ws_service.js")
	if err != nil{
		log.Println(err)
	}

	tmp.Execute(c.Writer, nil)
}

func (h *Handler) JSFile(c *gin.Context){
	http.ServeFile(c.Writer, c.Request, "client/auctionPage/static/ws_service.js")
}

func (h *Handler) AuctionAdminPage(c *gin.Context){
	
}
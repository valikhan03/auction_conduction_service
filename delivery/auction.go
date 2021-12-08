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

	tmp.Execute(c.Writer, nil)
}


func (h *Handler) AuctionAdminPage(c *gin.Context){
	
}
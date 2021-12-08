package delivery

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) ConnectAuction(c *gin.Context){
	auction_id, ok := c.Params.Get("id");  
	if ok == false{
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	

	result, err := h.RemSer.CheckAccess(auction_id, "")
	if err != nil{
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if result{
		//redirect to ws-endpoint
	}
}
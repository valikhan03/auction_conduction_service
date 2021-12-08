package main

import(
	"github.com/gin-gonic/gin"

	"auction_conduction/delivery"
)

func main(){
	remoteService := delivery.NewRemoteService("")
	handler := delivery.NewHandler(remoteService)

	router := gin.New()

	auction := router.Group("auction")
	{
		auction.GET("/connect/:id", handler.ConnectAuction)
		auction.GET("/service/:id", handler.AuctionWSEndpoint)
		auction.GET("/:id", handler.AuctionPage)
	}
}


package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"

	"auction_conduction/database"
	"auction_conduction/delivery"
)

func main() {
	repository := database.NewRepository(initRedis(), initPostgres())
	handler := delivery.NewHandler(repository)
	router := gin.New()

	auction := router.Group("auction")
	{
		auction.GET("/service/:id", handler.AuctionWSEndpoint)
		auction.GET("/:id", handler.AuctionPage)
		auction.GET("/ws-file", handler.JSFile)
	}

	server := &http.Server{
		Addr:    ":8100",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}


func initPostgres() *sqlx.DB{
	db, err := sqlx.Connect("pgx", "")
	if err != nil{
		log.Fatal(err)
	}

	return db
}

func initRedis() *redis.Client{
	redis_client := redis.NewClient(&redis.Options{
		Addr: "",
		Password: "",
		DB: 0,
	})
	return redis_client
}
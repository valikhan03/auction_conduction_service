package database

import(
	"github.com/jmoiron/sqlx"
	"github.com/go-redis/redis/v8"
)
type Repository struct{
	postgres *sqlx.DB
	redis *redis.Client
}


func NewRepository(redis_client *redis.Client, postgres_db *sqlx.DB) *Repository{
	return &Repository{
		redis: redis_client,
		postgres: postgres_db,
	}
}


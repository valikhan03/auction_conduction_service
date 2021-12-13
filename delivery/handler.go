package delivery

import(
	"auction_conduction/database"
)

type Handler struct{
	repository *database.Repository
}

func NewHandler(repos *database.Repository) *Handler{
	return &Handler{
		repository: repos,
	}
}


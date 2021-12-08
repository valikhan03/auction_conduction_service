package delivery

type Handler struct{
	RemSer RemoteService
}

func NewHandler(rs *RemoteService) *Handler{
	return &Handler{
		RemSer: *rs,
	}
}


package ws

type Handler struct {
	hub *Hub
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

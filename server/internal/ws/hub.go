package ws

type Room struct {
	ID      string             `json:"ID"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Room map[string]*Room
}

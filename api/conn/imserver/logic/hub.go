package logic

var Hubs = NewHub()

type Hub struct {
	clients    map[*Client]bool
	Broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.reader)
			}
		case message := <-h.Broadcast:
			for client := range h.clients {
				select {
				case client.reader <- message:
				default:
					close(client.reader)
					delete(h.clients, client)
				}
			}
		}
	}
}

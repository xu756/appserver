package logic

import (
	"sync"
)

var Hubs = NewHub()

type Hub struct {
	clients    sync.Map
	Broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	lock       sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.lock.Lock()
			h.clients.Store(client.imId, client)
			h.lock.Unlock()
		case client := <-h.unregister:
			h.lock.Lock()
			h.clients.Delete(client.imId)
			client.close <- true // 关闭客户端
			h.lock.Unlock()
		case message := <-h.Broadcast:
			h.clients.Range(func(_, value interface{}) bool {
				client := value.(*Client)
				select {
				case client.write <- message:
				default:
					h.lock.Lock()
					client.close <- true
					h.clients.Delete(client.imId)
					h.lock.Unlock()
				}
				return true
			})
		}
	}
}

func (h *Hub) FindOneByImId(imId string) *Client {
	h.lock.RLock()
	defer h.lock.RUnlock()
	value, ok := h.clients.Load(imId)
	if !ok {
		return nil
	}
	return value.(*Client)
}

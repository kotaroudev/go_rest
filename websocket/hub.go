package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Hub struct {
	clients    []*Client
	register   chan *Client
	unregister chan *Client
	mutex      *sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make([]*Client, 0),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		mutex:      &sync.Mutex{},
	}
}

func (hub *Hub) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// upgrader.Upgrade es la funcion que hace el upgrade de protocolo
	// de htttp a websocket
	socket, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	client := NewClient(hub, socket)
	hub.register <- client

	go client.Write()
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			hub.onConnect(client)
		case client := <-hub.unregister:
			hub.onDisconnect(client)
		}
	}
}

func (hub *Hub) onConnect(client *Client) {
	log.Println("Client connected", client.socket.RemoteAddr())

	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	client.id = client.socket.RemoteAddr().String()
	hub.clients = append(hub.clients, client)
}

func (hub *Hub) onDisconnect(client *Client) {
	log.Println("Client Disconnected", client.socket.RemoteAddr())
	client.socket.Close()
	hub.mutex.Lock()
	defer hub.mutex.Unlock()

	i := -1
	for idx, c := range hub.clients {
		if c.id == client.id {
			i = idx
		}
	}

	// Recordar que hub.clients es un slice y para eliminar
	// un elemento se debe identificar el idx del elemento
	// y reemplazarlos por los elementos a partir i+1
	// [1,2,3,4] i = 2
	// [i:] = [3,4]
	// [i+1:] = [4]
	// Copia el 4 en la posicion de [i:]
	// Resultado final: [1,2,4,4]
	copy(hub.clients[i:], hub.clients[i+1:])
	// Ahora el ultimo elemento ponlo como nil
	// [1,2,4,nil]
	hub.clients[len(hub.clients)-1] = nil
	// hub.clients va ser igual a hub.clients[0:3]
	// [1,2,4]
	hub.clients = hub.clients[:len(hub.clients)-1]
}

func (hub *Hub) Broadcast(message interface{}, ignore *Client) {
	data, _ := json.Marshal(message)

	for _, client := range hub.clients {
		if client != ignore {
			client.outbound <- data
		}
	}
}

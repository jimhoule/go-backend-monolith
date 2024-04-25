package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"net/http"

	gorillaWebsocket "github.com/gorilla/websocket"
)

//type Payload = any
type Event struct{
	Name    string  `json:"name"`
	Payload json.RawMessage `json:"payload,omitempty"`
}
type EventHandler func(client *Client, payload json.RawMessage)

// Upgrades incomming http requests into a persitent websocket connection
var upgrader = gorillaWebsocket.Upgrader{
	CheckOrigin: func(request *http.Request) bool {
		return true
	},
}

// References a connected client (frontend connection)
type Client struct {
	Connection *gorillaWebsocket.Conn
	Server *Server
	// Avoids concurrent writes on websocket
	Egress chan Event
}

// Emits event
func (c *Client) Emit(eventName string, payload any) {
	marshalledPayload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("error marshalling payload: %v", err)
	}

	event := Event{
		Name: eventName,
		Payload: marshalledPayload,
	}

	c.Egress <- event
}

// Reads messages from frontend
// NOTE: Must be ran in a goroutine
func (c *Client) ReadMessages() {
	// Closes the Connection once this function is done
	defer c.Server.RemoveClient(c)

	for {
		// Reads the next message in queue in the connection
		// NOTE: If Connection is closed, we will Recieve an error here
		_, messagePayload, err := c.Connection.ReadMessage()
		if err != nil {
			// NOTE: We only want to log Strange errors, but not simple Disconnection
			isUnexpectedCloseError := gorillaWebsocket.IsUnexpectedCloseError(
				err,
				gorillaWebsocket.CloseGoingAway,
				gorillaWebsocket.CloseAbnormalClosure,
			)
			if isUnexpectedCloseError  {
				fmt.Printf("unexpected close error: %v", err)
			} else {
				fmt.Printf("expected close error: %v", err)
			}

			// Breaks the loop to close connection (removes client)
			break
		}

		//Marshals incoming data into a Event struct
		var event Event
		err = json.Unmarshal(messagePayload, &event);
		if err != nil {
			fmt.Printf("error unmarshalling message: %v", err)
		}

		// Routes the Event
		c.Server.RouteEvent(event, c);
	}
}

// Writes messages to frontend
// NOTE: Must be ran in a goroutine
func (c *Client) WriteMessages() {
	// Closes the Connection once this function is done
	defer c.Server.RemoveClient(c)

	for {
		// NOTE: Ok will be false if egress channel is closed
		event, ok := <-c.Egress
		if !ok {
			// Returns close message
			err := c.Connection.WriteMessage(gorillaWebsocket.CloseMessage, nil)
			if err != nil {
				fmt.Printf("connection is closed: %v", err)
			}

			// Breaks the loop to close connection (removes client)
			break
		}

		marshalledEvent, err := json.Marshal(event)
		if err != nil {
			fmt.Printf("error marshalling event: %v", err)
		}
		
		// Sends event to frontend
		err = c.Connection.WriteMessage(gorillaWebsocket.TextMessage, marshalledEvent)
		if err != nil {
			fmt.Printf("error writing message: %v", err)
		}
	}
}

// Holds references to all clients and handlers
type Server struct{
	// Using a syncMutex to lock state before editing clients
	sync.RWMutex
	clients map[*Client]bool
	handlers map[string]EventHandler
}

// Serves and allows connections
func (s *Server) ServeWS(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("new connection")

	connection, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Sets incoming messages size limit (messages fron frontend)
	// NOTE: Files are being sent from javascript frontend as utin8arrays
	connection.SetReadLimit(50000000 * 8)

	// Create new client
	client := &Client{
		Connection: connection,
		Server:    s,
		Egress:     make(chan Event),
	}
	// Add new client to server
	s.AddClient(client)

	// Starts read and write process
	go client.ReadMessages()
	go client.WriteMessages()
}

// Adds event handler
func (s *Server) On(eventName string, handler EventHandler) {
	s.handlers[eventName] = handler
}

// Routes event to the correct handler
func (s *Server) RouteEvent(event Event, client *Client) {
	handler, ok := s.handlers[event.Name]
	if ok {
		handler(client, event.Payload);
	}
}

// Adds client to server clients
func (s *Server) AddClient(client *Client) {
	// Locks so we can manipulate
	s.Lock()
	defer s.Unlock()

	s.clients[client] = true
}

// Removes client from server clients
func (s *Server) RemoveClient(client *Client) {
	// Locks so we can manipulate
	s.Lock()
	defer s.Unlock()

	_, ok := s.clients[client]
	if ok {
		client.Connection.Close()
		delete(s.clients, client)
	}
}

var webSocketServer *Server

func Get() *Server {
	if webSocketServer == nil {
		webSocketServer = &Server{
			clients:  map[*Client]bool{},
			handlers: map[string]EventHandler{},
		}
	}
	
	return webSocketServer
}
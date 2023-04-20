package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// We'll need to define an Upgrader this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// define a reader which will listen for new messages being sent to our WebSocket endpoint
func reader(conn *websocket.Conn) {
	for {
		// read a message
		// messageType is either TextMessage or BinaryMessage.
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		// print message
		fmt.Printf("Received: %s\n", string(p))

		// send back message
		if err := conn.WriteMessage(messageType, p); err != nil {
			fmt.Println(err)
			return
		}

	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		fmt.Println(err)
	}
	// listen indefinitely for new messages coming through on our WebSocket connection
	reader(ws)
}

func main() {
	http.HandleFunc("/ws", wsEndpoint)
	fmt.Println(http.ListenAndServe(":8080", nil))
}

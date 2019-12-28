package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// WsServer specialises on serving ws requests.
type WsServer struct {
	// The actual http server to be upgraded
	http.ServeMux
	// node is a pointer to the actual node
	node *Node
}

func newWsServer(node *Node) *WsServer {
	return &WsServer{
		node: node,
	}
}

func (ws *WsServer) setUp() {
	ws.HandleFunc("/", ws.serveWs)
}

func (ws *WsServer) serveWs(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("Http server upgrade connection: ", err)
		return
	}
	configureWebsocket(conn)
	node := ws.node
	client := newClient(conn)
	go node.Accept(client)
}

func (ws *WsServer) Serve(port string) {
	address := fmt.Sprintf(":%s", port)
	log.Printf("WS server listening on %s", address)
	err := http.ListenAndServe(address, ws)
	if err != nil {
		log.Fatal(err)
	}
}

func configureWebsocket(ws *websocket.Conn) {
	ws.SetReadLimit(maxMessageSize)
	setSocketReadDeadline(pongWait, ws)
	ws.SetPongHandler(func(string) error {
		setSocketReadDeadline(pongWait, ws)
		return nil
	})
}

func setSocketReadDeadline(waitTime time.Duration, ws *websocket.Conn) {
	if err := ws.SetReadDeadline(time.Now().Add(waitTime)); err != nil {
		log.Fatal(err)
	}
}

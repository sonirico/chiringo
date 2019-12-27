package main

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type Joiner interface {
	Join(chan *client)
}

type Leaver interface {
	Leave(chan *client)
}

var (
	parser = NewMessageParser()
)

type client struct {
	// The unique identifier for a client
	Id uuid.UUID
	// The actual connection. I would have preferred to accept it
	// on methods, matching io.ReadWriteCloser, but that would render
	// unusable the socket's low level API
	socket *websocket.Conn
	// Channel to send outbound messages
	send chan Message
}

func newClient(conn *websocket.Conn) *client {
	c := &client{
		Id:     uuid.New(),
		socket: conn,
		send:   make(chan Message),
	}
	return c
}

func (c *client) Join(ch chan *client) {
	log.Printf("client joined: %s", c.Id)
	ch <- c
}

func (c *client) Leave(ch chan *client) {
	log.Printf("client left: %s", c.Id)
	ch <- c
}

func (c *client) Read(ch chan *ClientMessage) {
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		msg = bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))
		message, err := parser.Parse(msg)
		if err != nil {
			log.Printf(err.Error())
		} else {
			ch <- NewClientMessage(c.Id, message)
		}
		log.Printf("Message received: %s\n", string(msg))
	}
}
func (c *client) Write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case msg, ok := <-c.send:
			c.socket.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.socket.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(Serialize(msg))

			// Add queued messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				msg := <-c.send
				w.Write(Serialize(msg))
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.socket.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.socket.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *client) Exit() {
	close(c.send)
	err := c.socket.Close()
	if err != nil {
		log.Printf("error on closing client connection: %v", err)
	}
}

func (c *client) Send(msg Message) {
	c.send <- msg
}

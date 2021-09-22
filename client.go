package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type wsClient struct {
	mu   sync.Mutex
	conn *websocket.Conn
}

func newClient(url string, header http.Header) (*wsClient, error) {
	client := new(wsClient)

	conn, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		return nil, err
	}

	client.conn = conn
	return client, nil
}

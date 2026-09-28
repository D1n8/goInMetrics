package server

import (
	"backend/internal/pcstat"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func checkOrigin(r *http.Request) bool {
	// FIX: Add domain check
	return true
}

var upgrader = websocket.Upgrader{ CheckOrigin: checkOrigin }

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP to websocket
	connection, err := upgrader.Upgrade(w, r , nil)
	if err != nil {
		log.Printf("%20s%v", "Websocket error:", err)
		return
	}
	// Close connection after any scenario
	defer connection.Close()

	updateTicker := time.NewTicker(1 * time.Second)
	defer updateTicker.Stop()

	for {
		select {
		case <-updateTicker.C:

			// FIX:	move jsonData, err := pcstat.SendDataToServer() into main.go
			jsonData, err := pcstat.SendDataToServer()
			if err != nil {
				log.Printf("%20s%v", "Collection error:", err)
				return
			}

			err = connection.WriteMessage(websocket.TextMessage, jsonData)
			if err != nil {
				log.Printf("%20s%v", "Write error:", err)
				return
			}
		}
	}
}

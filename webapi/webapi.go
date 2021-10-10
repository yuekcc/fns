package webapi

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{}
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/health", handleHealthCheck)
	router.HandleFunc("/api/stream", handleStreamTrans)

	return router
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func handleStreamTrans(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": false, "message": "failed to upgrade to websocket"})
		return
	}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			break
		}
	}
}

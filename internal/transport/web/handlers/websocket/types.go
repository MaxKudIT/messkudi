package websocket

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var Upgrader websocket.Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true //  !!!  Осторожно! В production - проверять origin!
	},
}

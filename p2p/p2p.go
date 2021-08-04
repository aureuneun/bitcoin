package p2p

import (
	"net/http"

	"github.com/aureuneun/bitcoin/utils"
	"github.com/gorilla/websocket"
)

var conns []*websocket.Conn
var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
	conns = append(conns, conn)
	for {
		_, payload, err := conn.ReadMessage()
		if err != nil {
			break
		}
		for _, aConn := range conns {
			if aConn != conn {
				utils.HandleErr(aConn.WriteMessage(websocket.TextMessage, []byte(payload)))

			}
		}
	}
}

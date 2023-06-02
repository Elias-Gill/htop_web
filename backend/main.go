package main

import (
	"encoding/json"
	"net/http"

	"github.com/elias-gill/htop/utils"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	// canal que contiene la lista de procesos. Se actualiza cada 200 milisegundos
	ch := make(chan *utils.ProcessesTree)
	go utils.InicializarSistema(ch)

	// crear primero una conexion http
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// convertirla en websocket
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			println(err.Error())
			return
		}

        // handler
		go func() {
			defer conn.Close()
			for {
				// leer mensajes del cliente
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					println(err.Error())
					return
				}
				println(string(msg))

				// mandar mensajes
				msg, err = json.Marshal(<-ch)
				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					println(err)
					return
				}
			}
		}()
	}))
}

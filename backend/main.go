package main

import (
	"encoding/json"
	"net/http"

	"github.com/elias-gill/htop/utils"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() { 
    /* INFO: emplementacion vieja con go routines
    INFO: fue al pedo lo de hacer con un channel al final. Los websockets son como las http pero sin cerrar conexion nomas */
	// canal que contiene la lista de procesos. Se actualiza cada 200 milisegundos
    /* ch := make(chan *utils.ProcessesTree)
	go utils.InicializarSistema(ch) */

	// Crear el server http
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Convertir la conexion a websocket
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			println(err.Error())
			return
		}

		go func() {
			defer conn.Close()
			for {
				// leer request del cliente
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					println(err.Error())
					return
				}
				println(string(msg))

                // buscar la data
                data, err:= utils.GetProcessesTree()
                if err != nil {
					println(err)
					return
                }

				// mandar la respuesta
				msg, err = json.Marshal(data)
				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					println(err)
					return
				}
			}
		}()
	}))
}

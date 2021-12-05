package restaurant

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/software-engineering-components/lab3/server/tools"
)

type ClientOrder struct {
	TableNamber int     `json:"tableNamber"`
	Items       []*Item `json:"items"`
}

type HttpHandlerFunc http.HandlerFunc

func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListMenu(rw, store)
		} else if r.Method == "POST" {
			handleCreateOrder(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleListMenu(rw http.ResponseWriter, store *Store) {
	res, err := store.ListMenu()
	if err != nil {
		log.Printf("Error(Database Query): %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}

func handleCreateOrder(r *http.Request, rw http.ResponseWriter, store *Store) {
	var clientOrder *ClientOrder
	if err := json.NewDecoder(r.Body).Decode(&clientOrder); err != nil {
		log.Printf("Error json caused an error: %s", err)
		tools.WriteJsonBadRequest(rw, "Bad JSON payload")
		return
	}

	resOrder, err := store.CreateOrder(clientOrder.TableNamber, clientOrder.Items)
	if err == nil {
		tools.WriteJsonOk(rw, resOrder)
	} else {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

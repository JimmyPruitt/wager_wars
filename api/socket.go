package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	ws "github.com/gorilla/websocket"
	wwdb "wager_wars/db"
)

type Api interface {
	Listen()
}

type WagerWarsSocketApi struct {
	host     string
	port     int
	upgrader ws.Upgrader
	db       wwdb.DB
}

type Options struct {
	Host string
	Port int
}

func New(o Options, db wwdb.DB) (WagerWarsSocketApi, error) {
	return WagerWarsSocketApi{
		host: o.Host,
		port: o.Port,
		db:   db,
		upgrader: ws.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}, nil
}

func (api WagerWarsSocketApi) respond(res http.ResponseWriter, req *http.Request) {
	conn, err := api.upgrader.Upgrade(res, req, nil)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	defer conn.Close()
	for {
		target, _, err := conn.ReadMessage() // Underscore is the message
		if err != nil {
			fmt.Println(err)
			break
		}

		err = conn.WriteMessage(target, []byte("yolo"))
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func (api WagerWarsSocketApi) Listen() {
	api.db.Subscribe(func(message []rune) {
	})

	host := strings.Join([]string{api.host, strconv.Itoa(api.port)}, ":")

	http.HandleFunc("/api", api.respond)
	fmt.Println(host)
	log.Fatal(http.ListenAndServe(host, nil))
}

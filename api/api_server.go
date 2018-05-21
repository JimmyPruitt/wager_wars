package api

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/graphql-go/handler"
	wwdb "wager_wars/db"
)

type Api interface {
	Listen()
}

type Options struct {
	Host string
	Port int
}

type wagerWarsApi struct {
	host string
	port int
	db   wwdb.DB
}

func BuildServer(o Options, db wwdb.DB) (Api, error) {
	return wagerWarsApi{
		host: o.Host,
		port: o.Port,
		db:   db,
	}, nil
}

func (api wagerWarsApi) Listen() {
	host := strings.Join([]string{api.host, strconv.Itoa(api.port)}, ":")

	schema, _ := buildSchema(api.db)

	h := handler.New(&handler.Config{
		Schema:   schema.GetSchema(),
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/api", h)
	log.Fatal(http.ListenAndServe(host, nil))
}

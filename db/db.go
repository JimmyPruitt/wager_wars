package wwdb

import (
	"strconv"
	"strings"

	r "gopkg.in/gorethink/gorethink.v4"
)

type db struct {
	connection *r.Session
}

type Options struct {
	Hosts    []string
	Port     int
	Database string
	Username string
	Password string
}

type subscribeFunc func([]rune)

type DB interface {
	Subscribe(subscribeFunc)
}

func New(co Options) (DB, error) {
	port := co.Port
	hosts := []string{}

	for _, host := range co.Hosts {
		hosts = append(hosts, strings.Join([]string{host, strconv.Itoa(port)}, ":"))
	}

	conn, err := r.Connect(r.ConnectOpts{
		Addresses: hosts,
		Database:  co.Database,
		Username:  co.Username,
		Password:  co.Password,
	})

	return db{conn}, err
}

func (conn db) Subscribe(callback subscribeFunc) {
}

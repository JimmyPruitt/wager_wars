package wwdb

import (
	"strconv"
	"strings"

	r "gopkg.in/gorethink/gorethink.v4"
)

type db struct {
	connection *r.Session
	database   string
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
	GetUser(string) (*User, error)
}

type User struct {
	Id         string
	TwitchId   string
	FacebookId string
	TwitterId  string
	GoogleId   string
	Opponents  []User
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

	return db{
		connection: conn,
		database:   co.Database,
	}, err
}

func (client db) Subscribe(callback subscribeFunc) {
}

func (client db) GetUser(id string) (*User, error) {
	res, err := r.Table("users").GetAll(id).InnerJoin(r.Table("users"), func(currentRow r.Term, opponentRow r.Term) interface{} {
		return currentRow.AtIndex("opponents").Contains(opponentRow.AtIndex("id"))
	}).Run(client.connection)

	defer res.Close()
	if err != nil {
		return nil, err
	}

	var rows []interface{}
	err = res.All(&rows)
	if err != nil {
		return nil, err
	}

	if len(rows) < 1 {
		return nil, nil
	}

	var user User = User{}

	for _, row := range rows {
		r := row.(map[string]interface{})
		left := r["left"].(map[string]interface{})
		right, _ := r["right"].(map[string]interface{})

		user.Id = left["id"].(string)
		user.TwitterId = left["twitter_id"].(string)
		user.TwitchId = left["twitch_id"].(string)
		user.FacebookId = left["facebook_id"].(string)
		user.GoogleId = left["google_id"].(string)
		user.Opponents = append(user.Opponents, User{
			Id:         right["id"].(string),
			TwitterId:  right["twitter_id"].(string),
			TwitchId:   right["twitch_id"].(string),
			FacebookId: right["facebook_id"].(string),
			GoogleId:   right["google_id"].(string),
		})
	}

	return &user, nil
}

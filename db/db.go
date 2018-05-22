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

	if err != nil {
		return nil, err
	}

	return db{
		connection: conn,
		database:   co.Database,
	}, nil
}

func (client db) Subscribe(callback subscribeFunc) {
}

type userOpponentRow struct {
	OpponentObjects r.Term // Golang r.Fold() behaves differently than fold() in the admin console, so this index must have a different name. Once a fix has been found, change this to just opponents
}

func (client db) GetUser(id string) (*User, error) {
	var defaultCarry interface{}
	res, err := r.Table("users").GetAll(id).InnerJoin(r.Table("users"), func(currentRow r.Term, opponentRow r.Term) interface{} {
		return currentRow.AtIndex("opponents").Contains(opponentRow.AtIndex("id"))
	}).Fold(defaultCarry, func(carry r.Term, current r.Term) interface{} {
		var defaultOpponents []interface{}
		opponents := []r.Term{current.AtIndex("right").Without("opponents")}
		return current.AtIndex("left").Merge(userOpponentRow{
			OpponentObjects: carry.AtIndex("OpponentObjects").Default(defaultOpponents).Add(opponents),
		})
	}).Run(client.connection)

	defer res.Close()
	if err != nil {
		return nil, err
	}

	var row interface{}
	err = res.One(&row)
	if err != nil {
		if err == r.ErrEmptyResult {
			return nil, nil
		}

		return nil, err
	}

	return client.coerceUser(row)
}

func (client db) coerceUser(row interface{}) (*User, error) {
	var user User = User{}
	if u, ok := row.(map[string]interface{}); ok {
		user.Id = u["id"].(string)
		user.TwitterId = u["twitter_id"].(string)
		user.TwitchId = u["twitch_id"].(string)
		user.FacebookId = u["facebook_id"].(string)
		user.GoogleId = u["google_id"].(string)

		if opponents, ok := u["OpponentObjects"].([]interface{}); ok {
			for _, opponentArray := range opponents {
				if o, ok := opponentArray.(map[string]interface{}); ok {
					user.Opponents = append(user.Opponents, User{
						Id:         o["id"].(string),
						TwitterId:  o["twitter_id"].(string),
						TwitchId:   o["twitch_id"].(string),
						FacebookId: o["facebook_id"].(string),
						GoogleId:   o["google_id"].(string),
					})
				}
			}
		}
	}

	return &user, nil
}

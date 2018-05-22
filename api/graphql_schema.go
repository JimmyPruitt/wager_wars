package api

import (
	"errors"
	"strings"

	"github.com/graphql-go/graphql"
	wwdb "wager_wars/db"
)

type Schema interface {
	GetSchema() *graphql.Schema
}

type wagerWarsSchema struct {
	schema graphql.Schema
	db     wwdb.DB
}

func (wws wagerWarsSchema) GetSchema() *graphql.Schema {
	return &wws.schema
}

type SocialMedia int

const (
	Twitter SocialMedia = 1 << iota
	Facebook
	Twitch
	Google
)

func buildSchema(db wwdb.DB) (Schema, error) {
	schema := wagerWarsSchema{
		db: db,
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "Query",
		Fields: schema.getRootFields(),
	}

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}

	s, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, err
	}

	schema.schema = s
	return schema, nil
}

func (wws wagerWarsSchema) getRootFields() graphql.Fields {
	return graphql.Fields{
		"user": wws.getUserRootFields(),
	}
}

func (wws wagerWarsSchema) getUserRootFields() *graphql.Field {
	return &graphql.Field{
		Type: wws.getUserQueryField(true),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "The unique user ID for the Wager Wars user",
				Type:        graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if id, ok := p.Args["id"].(string); ok {
				user, err := wws.db.GetUser(id)
				if user == nil {
					return nil, errors.New(strings.Join([]string{"User \"", id, "\" does not exist"}, ""))
				}

				return *user, err
			}

			return nil, errors.New("Failed to locate user")
		},
	}
}

func (wws wagerWarsSchema) getSocialMediaIdField(sm SocialMedia) *graphql.Field {
	var t string
	switch sm {
	case Twitter:
		t = "Twitter"
	case Facebook:
		t = "Facebook"
	case Twitch:
		t = "Twitch"
	case Google:
		t = "Google"
	}

	return &graphql.Field{
		Type:        graphql.String,
		Description: strings.Join([]string{"The ID for the user's linked", t, "account"}, " "),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(wwdb.User); ok {
				var id string
				switch sm {
				case Twitter:
					id = user.TwitterId
				case Facebook:
					id = user.FacebookId
				case Twitch:
					id = user.TwitchId
				case Google:
					id = user.GoogleId
				}

				return id, nil
			}

			return errors.New(strings.Join([]string{"Failed to resolve field for user's linked", t, "account"}, " ")), nil
		},
	}
}

func (wws wagerWarsSchema) getUserQueryField(includeOpponents bool) *graphql.Object {
	userType := wws.getBaseUserType("User")

	if includeOpponents {
		userType.AddFieldConfig("opponents", &graphql.Field{
			Type:        graphql.NewList(wws.getBaseUserType("Opponent")),
			Description: "A list of users with whom this user has agreed to join battle",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(wwdb.User); ok {
					return user.Opponents, nil
				}

				return nil, errors.New("Failed to resolve field \"opponents\" on type \"user\"")
			},
		})
	}

	return userType
}

func (wws wagerWarsSchema) getBaseUserType(name string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        name,
		Description: "A Wager Wars user and their linked social media accounts",
		Fields: graphql.Fields{
			"twitch_id":   wws.getSocialMediaIdField(Twitch),
			"facebook_id": wws.getSocialMediaIdField(Facebook),
			"twitter_id":  wws.getSocialMediaIdField(Twitter),
			"google_id":   wws.getSocialMediaIdField(Google),
			"id": &graphql.Field{
				Type:        graphql.String,
				Description: "The user's unique Wager Wars account identifier",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(wwdb.User); ok {
						return user.Id, nil
					}

					return nil, errors.New("Failed to resolve field \"id\" on type \"user\"")
				},
			},
		},
	})
}

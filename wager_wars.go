package main

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	wwapi "wager_wars/api"
	wwdb "wager_wars/db"
)

func main() {
	evn := getEnvVarNames()
	dbv, err := getDbVars(evn)

	if err != nil {
		log.Fatalln(err)
	}

	db, err := wwdb.New(wwdb.Options{
		Hosts:    []string{dbv.Host},
		Port:     dbv.Port,
		Database: dbv.Name,
		Username: dbv.User,
		Password: dbv.Pass,
	})

	if err != nil {
		log.Fatalln(err)
	}

	api, err := wwapi.New(wwapi.Options{
		Host: "0.0.0.0",
		Port: 8080,
	}, db)

	if err != nil {
		log.Fatalln(err)
	}

	api.Listen()
}

type dbVars struct {
	Host string
	Port int
	Name string
	User string
	Pass string
}

type envVarNames struct {
	DbHost string
	DbPort string
	DbName string
	DbUser string
	DbPass string
}

func getEnvVarNames() envVarNames {
	return envVarNames{
		DbHost: "WAGER_WARS_DB_HOST",
		DbPort: "WAGER_WARS_DB_PORT",
		DbName: "WAGER_WARS_DB_NAME",
		DbUser: "WAGER_WARS_DB_USER",
		DbPass: "WAGER_WARS_DB_PASS",
	}
}

func getDbVars(evn envVarNames) (dbVars, error) {
	port, err := strconv.Atoi(os.Getenv(evn.DbPort))
	if err != nil {
		return dbVars{}, errors.New(strings.Join([]string{"Environment variable", evn.DbPort, "has not been set or is not a valid base 10 number"}, " "))
	}

	vars := dbVars{
		Host: os.Getenv(evn.DbHost), // 35.230.1.175
		Port: port,                  // 28015
		Name: os.Getenv(evn.DbName), // wager_wars
		User: os.Getenv(evn.DbUser), // wager_wars_app
		Pass: os.Getenv(evn.DbPass), // w4g3r W4r5
	}

	if len([]rune(vars.Host)) < 1 {
		err = errors.New(strings.Join([]string{"Environment variable", evn.DbHost, "has not been set"}, " "))
	}

	if len([]rune(vars.Name)) < 1 {
		err = errors.New(strings.Join([]string{"Environment variable", evn.DbName, "has not been set"}, " "))
	}

	if len([]rune(vars.User)) < 1 {
		err = errors.New(strings.Join([]string{"Environment variable", evn.DbUser, "has not been set"}, " "))
	}

	if len([]rune(vars.Pass)) < 1 {
		err = errors.New(strings.Join([]string{"Environment variable", evn.DbPass, "has not been set"}, " "))
	}

	return vars, err
}

package provider

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Message struct {
	Err string `json:"msg"`
}
type User struct {
	Message
	Login    *string
	Password *string
	Nickname *string
	Avatar   *string
}

type Provider struct {
	db *sql.DB
}

func CreateDBProvider(host string, port int, user string, dbname string, password string) *Provider {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &Provider{db: conn}
}

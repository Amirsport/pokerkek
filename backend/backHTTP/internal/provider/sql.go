package provider

import "fmt"

func InsertReg(user User) error {
	return nil
}

func (db *Provider) Registration(u User) error {
	_, err := db.db.Exec("insert into users(pass,login,nickname) values ($1,$2,$3)", *u.Password, *u.Login, *u.Nickname)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (db *Provider) Auth(u User) error {
	return nil
}

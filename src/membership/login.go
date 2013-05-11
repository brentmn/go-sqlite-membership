package main

import (
	"errors"
	"fmt"
)

func (this *Db) Login(r Registration) (Registration, error) {
	var reg Registration
	stmt, err := this.db.Prepare("select id, email, password from registration where email = ?")
	if err != nil {
		fmt.Println(err)
		return reg, err
	}
	defer stmt.Close()
	var id int64
	var email, pass string
	err = stmt.QueryRow(r.Email).Scan(&id, &email, &pass)
	if err != nil {
		fmt.Println(err)
		return reg, err
	}

	if PassMatch(r.Password, pass) {
		fmt.Println("OKAY login", email)
	} else {
		fmt.Println("BAD login", email)
		err = errors.New("wrong password")
	}

	reg = Registration{id, email, ""}

	return reg, err
}

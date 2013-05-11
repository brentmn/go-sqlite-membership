package main

import (
	"fmt"
)

func (this *Db) Login(r Registration) (Registration, error) {
	var reg Registration
	stmt, err := this.db.Prepare("select id, email from registration where email = ? and password = ?")
	if err != nil {
		fmt.Println(err)
		return reg, err
	}
	defer stmt.Close()
	var id int64
	var email string
	err = stmt.QueryRow(r.Email, r.Password).Scan(&id, &email)
	if err != nil {
		fmt.Println(err)
		return reg, err
	}
	fmt.Println("okay login", id, email)

	reg = Registration{id, email, ""}

	return reg, err
}

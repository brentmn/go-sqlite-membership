package main

import (
	"errors"
	"fmt"
)

func (this *Db) Get(id int64) (Registration, error) {
	var reg Registration
	stmt, err := this.db.Prepare("select id, email from registration where id = ?")
	if err != nil {
		fmt.Println("error getting user by id: " + err.Error())
		return reg, err
	}
	defer stmt.Close()
	var email string
	err = stmt.QueryRow(&id).Scan(&id, &email)
	if err != nil {
		fmt.Println("error in get function for " + email + " --- " + err.Error())
		return reg, err
	}
	fmt.Println("found", id, email)

	reg = Registration{id, email, ""}

	return reg, err
}
func (this *Db) GetByEmail(email string) (Registration, error) {
	var reg Registration
	stmt, err := this.db.Prepare("select id, email from registration where email = ?")
	if err != nil {
		fmt.Println("error getting " + email + " " + err.Error())
		return reg, err
	}
	defer stmt.Close()
	var id int64
	err = stmt.QueryRow(&id).Scan(&id, &email)
	if err != nil {
		fmt.Println("error in get function for " + email + " --- " + err.Error())
		return reg, err
	}
	fmt.Println("found", id, email)

	reg = Registration{id, email, ""}

	return reg, err
}
func (this *Db) List() error {
	rows, err := this.db.Query("select id, email, password from registration")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var email string
		var password string
		rows.Scan(&id, &email, &password)
		fmt.Println(id, email, password)
	}

	return errors.New("haven't implemented list of registration")
	//TODO: return set of registration on list

}
func (this *Db) Create(r *Registration) (*Registration, error) {

	_, err := this.GetByEmail(r.Email)

	if err == nil {
		return r, errors.New("User already exists.")
	} else {
		fmt.Println("user does not exist, okay to create user -- " + r.Email + " -- error:" + err.Error())
	}

	tx, err := this.db.Begin()
	if err != nil {
		fmt.Println(err)
		return r, err
	}
	stmt, err := tx.Prepare("insert into registration(id, email, password) values(?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return r, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(nil, r.Email, r.Password)
	tx.Commit()

	r.Password = ""
	return r, err

}
func (this *Db) Update(r *Registration) (*Registration, error) {
	var reg Registration
	reg, err := this.GetByEmail(r.Email)

	if reg.Email == "" {
		return r, errors.New("User does not exist.")
	}

	tx, err := this.db.Begin()
	if err != nil {
		fmt.Println(err)
		return r, err
	}
	stmt, err := tx.Prepare("update registration set email=?, password=? where id=?")
	if err != nil {
		fmt.Println(err)
		return r, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.Email, r.Password, r.ID)
	tx.Commit()

	r.Password = ""
	return r, err

}
func (this *Db) Delete(id int64) error {

	tx, err := this.db.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}
	stmt, err := tx.Prepare("delete from registration where id=?")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	tx.Commit()

	return err
}

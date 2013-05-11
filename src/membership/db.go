package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Db struct {
	Namespace string
	db        *sql.DB
}

func (this *Db) sCheck() error {
	db, err := sql.Open("sqlite3", "data/"+this.Namespace+".db")
	this.db = db
	if err != nil {
		fmt.Println(err)
	}

	this.setup()

	return err
}
func (this *Db) Close() {
	this.db.Close()
}
func (this *Db) setup() {
	sqls := []string{
		"create table registration (id integer not null primary key autoincrement, email text, password text)",
		"delete from registration",
	}

	//NOTE: this is using err for program continuation by allowing error if table exists
	//for _, sql := range sqls {
	_, err := this.db.Exec(sqls[0])
	if err != nil {
		fmt.Printf("%q: %s\n", err, sqls[0])
		return
	}
	//}

}

package db

import (
	"avatar.com/avatar/db/conf"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", conf.ConnString)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS avatar (id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY, user_id BIGINT UNIQUE, src_path TEXT);")
	if err != nil {
		log.Fatal("Cant' create table")
		return
	}
	log.Println("Created")
	if err != nil {
		log.Fatalln(err)
	}
}

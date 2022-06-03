package db

import (
	"avatar.com/avatar/db/conf"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var db *sql.DB

func init() {
	var err error
	for {
		db, err = sql.Open("postgres", conf.ConnString)
		if err == nil {
			break
		}
		log.Println(err)
		log.Println("Can't open connection with db\nTry again in 5s")
		time.Sleep(5 * time.Second)
	}

	for {
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS avatar (id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY, user_id BIGINT UNIQUE, src_path TEXT);")
		if err == nil {
			log.Println("Table `avatar` created")
			break
		}
		log.Println(err)
		log.Println("Cant' create table `avatar`\nTry again in 5s")
		time.Sleep(5 * time.Second)
	}
}

package db

import (
	"avatar.com/avatar/db/conf"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var db *sql.DB

const UserAvatarTableName = "avatar"
const TeamAvatarTableName = "team_avatar"

func openConn(driverName string, connString string) {
	var err error
	for {
		db, err = sql.Open(driverName, connString)
		if err == nil {
			break
		}
		log.Println(err)
		log.Println("Can't open connection with db\nTry again in 5s")
		time.Sleep(5 * time.Second)
	}
}

func initCreateTable(tableName string, query string) {
	for {
		_, err := db.Exec(query)
		if err == nil {
			log.Printf("Table `%s` created\n", tableName)
			break
		}
		log.Println(err)
		log.Printf("Cant' create table `%s`\nTry again in 5s\n", tableName)
		time.Sleep(5 * time.Second)
	}
}

func init() {
	openConn("postgres", conf.ConnString)
	initCreateTable(UserAvatarTableName, "CREATE TABLE IF NOT EXISTS avatar (id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY, user_id BIGINT UNIQUE, src_path TEXT);")
	initCreateTable(TeamAvatarTableName, "CREATE TABLE IF NOT EXISTS team_avatar (id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY, team_id BIGINT UNIQUE, src_path TEXT);")
}

package conf

import (
	"log"
	"os"
)

var dbUser = os.Getenv("DB_AVATAR_USER")
var dbPsw = os.Getenv("DB_AVATAR_PSW")
var DBName = os.Getenv("DB_AVATAR_DBNAME")
var dbSslmode = os.Getenv("DB_AVATAR_SSL")
var dbHostname = os.Getenv("DB_AVATAR_HOSTNAME")
var dbPort = os.Getenv("DB_AVATAR_PORT")

func init() {
	if dbUser == "" {
		log.Fatal("environment variable `DB_USER` must be specified and must be non empty value")
	}
	if dbPsw == "" {
		log.Fatal("environment variable `DB_PSW` must be specified and must be non empty value")
	}
	if DBName == "" {
		log.Fatal("environment variable `DB_NAME` must be specified and must be non empty value")
	}
	if dbSslmode == "" {
		log.Fatal("environment variable `DB_SSL` must be specified and must be non empty value")
	}
	if dbHostname == "" {
		log.Fatal("environment variable `DB_AVATAR_HOSTNAME` must be specified and must be non empty value")
	}
	if dbPort == "" {
		log.Fatal("environment variable `DB_AVATAR_PORT` must be specified and must be non empty value")
	}
}

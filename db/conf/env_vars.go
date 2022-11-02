package conf

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func checkEnvExists(envname string) string {
	val := os.Getenv(envname)
	log.Printf("%s = %s\n", envname, val)
	if val == "" {
		log.Fatalf("Environment variable `%s` must be specified and must be non empty value", envname)
	}
	return val
}

var dbUser string
var dbPsw string
var DBName string
var dbSslmode string
var dbHostname string
var dbPort string
var ConnString string

func init() {
	// Load the .env file in the current directory
	envErr := godotenv.Load()
	if envErr != nil {
		log.Println(".env file not found, load vars from session")
	}

	dbUser = checkEnvExists("POSTGRES_USER")
	dbPsw = checkEnvExists("POSTGRES_PASSWORD")
	dbHostname = checkEnvExists("DB_AVATAR_HOSTNAME")
	DBName = checkEnvExists("POSTGRES_DB")
	dbPort = checkEnvExists("PGPORT")
	dbSslmode = checkEnvExists("DB_AVATAR_SSL")

	ConnString = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHostname,
		dbUser,
		dbPsw,
		DBName,
		dbPort,
		dbSslmode)
}

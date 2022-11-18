package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func checkEnvExists(envname string) string {
	val := os.Getenv(envname)
	if val == "" {
		log.Fatalf("Environment variable `%s` must be specified and must be non empty value", envname)
	}
	return val
}

var UserServicePath string
var BaseAvatarPath string
var ServiceAvatarPort string

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

	BaseAvatarPath = checkEnvExists("AVATAR_IMAGES_DIR")

	dbUser = checkEnvExists("POSTGRES_USER")
	dbPsw = checkEnvExists("POSTGRES_PASSWORD")
	dbHostname = checkEnvExists("DB_AVATAR_HOSTNAME")
	DBName = checkEnvExists("POSTGRES_DB")
	dbPort = checkEnvExists("PGPORT")
	dbSslmode = checkEnvExists("DB_AVATAR_SSL")

	UserServicePath = checkEnvExists("USER_SERVICE_PATH")
	ServiceAvatarPort = checkEnvExists("SERVICE_AVATAR_PORT")

	ConnString = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHostname,
		dbUser,
		dbPsw,
		DBName,
		dbPort,
		dbSslmode)
}

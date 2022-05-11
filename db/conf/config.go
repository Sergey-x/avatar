package conf

import (
	"fmt"
)

var ConnString = fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
	dbHostname,
	dbUser,
	dbPsw,
	DBName,
	dbPort,
	dbSslmode)

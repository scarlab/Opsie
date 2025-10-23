package db

import (
	"fmt"
	"opsie/config"
)

// Build DSN (connection string)
var	dsn = fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.ENV.PG_User,
		config.ENV.PG_Password,
		config.ENV.PG_Host,
		config.ENV.PG_Port,
		config.ENV.PG_Database,
	)
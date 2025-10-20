package cli

import (
	"database/sql"
	"fmt"
	"log"
	"opsie/config"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	pgMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// ------------------------
// Migration helper
// ------------------------
func RunMigrate(cmd string, version int) {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.ENV.PG_User,
		config.ENV.PG_Password,
		config.ENV.PG_Host,
		config.ENV.PG_Port,
		config.ENV.PG_Database,
	)


	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("DB connection error: ", err)
	}
	defer dbConn.Close()
	

	driver, err := pgMigrate.WithInstance(dbConn, &pgMigrate.Config{
		MigrationsTable: "schema_migrations",
    	DatabaseName:    config.ENV.PG_Database,
	})
	if err != nil {
		log.Fatal("Migration driver error: ", err)
	}

	absPath, _ := filepath.Abs("./db/migrations")
	m, err := migrate.NewWithDatabaseInstance(
		"file://" + absPath,
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatal("Migration init error: ", err)
	}


	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Migration up error: ", err)
		}
		log.Println("✅ Migrations applied successfully", )

	case "down":
		if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Migration down error: ", err)
		}
		log.Println("✅ Migration rolled back successfully")

	case "force":
		if err := m.Force(version); err != nil {
			log.Fatal("Force migration error: ", err)
		}
		log.Printf("✅ Forced migration version set to %d", version)
	}
}

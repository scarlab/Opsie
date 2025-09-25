package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"watchtower/config"

	"github.com/golang-migrate/migrate/v4"
	pgMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "devx",
		Short: "DevX CLI for scaffolding domains and managing migrations",
	}


	// ------------------------
	// create-domain command
	// ------------------------
	newDomainCmd := &cobra.Command{
		Use:   "create-domain [name]",
		Short: "Scaffold a new domain",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := strings.ToLower(args[0])
			if err := createDomain(name); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("✅ Domain '%s' created successfully\n", name)
		},
	}

	rootCmd.AddCommand(newDomainCmd)

	// ------------------------
	// migrate command
	// ------------------------
	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "Run database migrations",
	}

	// migrate up
	migrateUpCmd := &cobra.Command{
		Use:   "up",
		Short: "Apply all pending migrations",
		Run: func(cmd *cobra.Command, args []string) {
			runMigrate("up", 0)
		},
	}

	// migrate down
	migrateDownCmd := &cobra.Command{
		Use:   "down",
		Short: "Rollback last migration",
		Run: func(cmd *cobra.Command, args []string) {
			runMigrate("down", 0)
		},
	}

	// migrate force
	migrateForceCmd := &cobra.Command{
		Use:   "force [version]",
		Short: "Force migration version",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			version, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal("Invalid version number")
			}
			runMigrate("force", version)
		},
	}

	migrateCmd.AddCommand(migrateUpCmd, migrateDownCmd, migrateForceCmd)
	rootCmd.AddCommand(migrateCmd)

	// ------------------------
	// Execute root command
	// ------------------------
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// ------------------------
// Domain scaffolding
// ------------------------
func createDomain(name string) error {
	basePath := filepath.Join("internal", "domain", name)
	if err := os.MkdirAll(basePath, 0755); err != nil {
		return err
	}

	templates := []string{
		"init.go.tpl",
		"handler.go.tpl",
		"service.go.tpl",
		"repository.go.tpl",
		"route.go.tpl",
		"type.go.tpl",
	}
	tplDir := filepath.Join("cmd", "cli", "templates", "domain")

	data := struct {
		PackageName string
		CreatedAt   string
	}{
		PackageName: name,
		CreatedAt:   time.Now().Format("2006/01/02 15:04:05"),
	}

	for _, tplFile := range templates {
		tplPath := filepath.Join(tplDir, tplFile)
		tpl, err := template.ParseFiles(tplPath)
		if err != nil {
			return err
		}

		outFile := filepath.Join(basePath, fmt.Sprintf("%s.%s", name, strings.TrimSuffix(tplFile, ".tpl")))
		f, err := os.Create(outFile)
		if err != nil {
			return err
		}
		defer f.Close()

		if err := tpl.Execute(f, data); err != nil {
			return err
		}
	}
	return nil
}

// ------------------------
// Migration helper
// ------------------------
func runMigrate(cmd string, version int) {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.ENV.PG_USER,
		config.ENV.PG_PASSWD,
		config.ENV.PG_HOST,
		config.ENV.PG_PORT,
		config.ENV.PG_DB,
	)
	

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("DB connection error: ", err)
	}
	defer dbConn.Close()

	driver, err := pgMigrate.WithInstance(dbConn, &pgMigrate.Config{})
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

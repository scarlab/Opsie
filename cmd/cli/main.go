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

	"opsie/config"

	"github.com/golang-migrate/migrate/v4"
	pgMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "devx",
		Short: "DevX CLI for scaffolding apis and managing migrations",
	}


	// ------------------------
	// create-api command
	// ------------------------
	newDomainCmd := &cobra.Command{
		Use:   "create-api [name]",
		Short: "Scaffold a new api inside core/api/",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := strings.ToLower(args[0])

			// Read the flag
			withWS, _ := cmd.Flags().GetBool("ws")

			fmt.Printf("🧱 Creating api '%s' (WebSocket: %v)\n", name, withWS)

			// Generate api files
			if err := createDomain(name, withWS); err != nil {
				log.Fatalf("❌ Failed to create api '%s': %v\n", name, err)
			}
		},
	}

	// Add flag
	newDomainCmd.Flags().BoolP("ws", "w", false, "Include WebSocket-enabled templates")


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
func createDomain(name string, withWS bool) error {
	basePath := filepath.Join("core", "api", name)

	// --- Check if the api already exists ---
	if _, err := os.Stat(basePath); err == nil {
		return fmt.Errorf("api '%s' already exists", name)
	}



	// --- Create base directory ---
	if err := os.MkdirAll(basePath, 0755); err != nil {
		return fmt.Errorf("failed to create api dir: %w", err)
	}

	// --- Template list ---
	templates := []string{
		"init.tpl",
		"handler.tpl",
		"service.tpl",
		"repository.tpl",
		"route.tpl",
		"type.tpl",
	}

	// if --ws flag is passed, use -ws.tpl versions where available
	if withWS {
		for i, name := range templates {
			wsPath := filepath.Join("cmd", "cli", "templates", "api", strings.Replace(name, ".tpl", "-ws.tpl", 1))
			if _, err := os.Stat(wsPath); err == nil {
				templates[i] = strings.Replace(name, ".tpl", "-ws.tpl", 1)
			}
		}
	}

	tplDir := filepath.Join("cmd", "cli", "templates", "api")
	data := struct {
		PackageName string
		CreatedAt   string
	}{
		PackageName: name,
		CreatedAt:   time.Now().Format("2006/01/02 15:04:05"),
	}

	// --- Generate files from templates ---
	for _, tplFile := range templates {
		tplPath := filepath.Join(tplDir, tplFile)
		tpl, err := template.ParseFiles(tplPath)
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", tplFile, err)
		}

		outFile := filepath.Join(basePath, fmt.Sprintf("%s.%s.go",
			name, strings.TrimSuffix(strings.TrimSuffix(tplFile, "-ws.tpl"), ".tpl")))

		f, err := os.Create(outFile)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", outFile, err)
		}
		defer f.Close()

		if err := tpl.Execute(f, data); err != nil {
			return fmt.Errorf("failed to execute template %s: %w", tplFile, err)
		}
		fmt.Printf("✅ Created %s\n", outFile)
	}

	fmt.Printf("🎉 Domain '%s' created successfully.\n", name)
	return nil
}

// ------------------------
// Migration helper
// ------------------------
func runMigrate(cmd string, version int) {
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

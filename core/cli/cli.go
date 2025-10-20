package cli

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)


func Init() {
	rootCmd := &cobra.Command{
		Use:   "devx",
		Short: "DevX CLI for scaffolding apis and managing migrations",
	}


	// ------------------------
	// api parent command
	// ------------------------
	apiCmd := &cobra.Command{
		Use:   "api",
		Short: "Manage backend API modules (create/delete/list)",
	}

	// ------------------------
	// api create command
	// ------------------------
	apiCreateCmd := &cobra.Command{
		Use:   "create [name]",
		Short: "Scaffold a new API inside core/api/",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := strings.ToLower(args[0])
			withWS, _ := cmd.Flags().GetBool("ws")

			fmt.Printf("🧱 Creating api '%s' (WebSocket: %v)\n", name, withWS)
			if err := CreateApi(name, withWS); err != nil {
				log.Fatalf("❌ Failed to create api '%s': %v\n", name, err)
			}
		},
	}
	apiCreateCmd.Flags().BoolP("ws", "w", false, "Include WebSocket-enabled templates")

	// ------------------------
	// api delete command
	// ------------------------
	apiDeleteCmd := &cobra.Command{
		Use:   "delete [name]",
		Short: "Delete an existing API and its related files",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := strings.ToLower(args[0])

			fmt.Printf("🗑️ Deleting api '%s'...\n", name)
			if err := DeleteApi(name); err != nil {
				log.Fatalf("❌ Failed to delete api '%s': %v\n", name, err)
			}
		},
	}

	// ------------------------
	// attach subcommands
	// ------------------------
	apiCmd.AddCommand(apiCreateCmd)
	apiCmd.AddCommand(apiDeleteCmd)

	// Register parent command with root
	rootCmd.AddCommand(apiCmd)


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
			RunMigrate("up", 0)
		},
	}

	// migrate down
	migrateDownCmd := &cobra.Command{
		Use:   "down",
		Short: "Rollback last migration",
		Run: func(cmd *cobra.Command, args []string) {
			RunMigrate("down", 0)
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
			RunMigrate("force", version)
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
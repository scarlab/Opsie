package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "devx",
		Short: "DevX CLI for scaffolding domains",
	}

	var author string

	// new-domain command
	newDomainCmd := &cobra.Command{
		Use:   "new-domain [name]",
		Short: "Scaffold a new domain",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := strings.ToLower(args[0])
			if author == "" {
				author = getGitAuthor()
			}

			if err := createDomain(name, author); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("✅ Domain '%s' created successfully by %s\n", name, author)
		},
	}

	newDomainCmd.Flags().StringVarP(&author, "author", "a", getGitAuthor(), "author info")

	// add subcommand to root
	rootCmd.AddCommand(newDomainCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// createDomain scaffolds a new domain folder with boilerplate files
func createDomain(name, author string) error {
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
		Author      string
		CreatedAt   string
	}{
		PackageName: name,
		Author:      author,
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

func getGitAuthor() string {
	name, _ := exec.Command("git", "config", "user.name").Output()
	// email, _ := exec.Command("git", "config", "user.email").Output()

	n := strings.TrimSpace(string(name))
	// e := strings.TrimSpace(string(email))

	if n != "" {
		return fmt.Sprint(n)
	}

	
	// Fallback to system environment
	user := os.Getenv("USER")
	if user == "" {
		user = os.Getenv("USERNAME") // Windows
	}

	if user == "" {
		user = "unknown" // final fallback
	}

	return user
}

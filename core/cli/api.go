package cli

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ------------------------
// Api scaffolding
// ------------------------
func CreateApi(name string, withWS bool) error {
    // --- Define output paths ---
    apiBase := filepath.Join("core", "api", name)
    serviceBase := filepath.Join("core", "services")
    repoBase := filepath.Join("core", "repositories")
    typeBase := filepath.Join("types")

    // --- Check if the API already exists ---
    if _, err := os.Stat(apiBase); err == nil {
        return fmt.Errorf("api '%s' already exists", name)
    }

    // --- Create directories ---
    dirs := []string{
        apiBase,                    // init, route, handler
        serviceBase,                // service
        repoBase,                   // repository
        typeBase,                   // type
    }

    for _, d := range dirs {
        if err := os.MkdirAll(d, 0755); err != nil {
            return fmt.Errorf("failed to create dir %s: %w", d, err)
        }
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

    // --- Adjust templates for WS if needed ---
    if withWS {
        for i, tplName := range templates {
            wsTplPath := filepath.Join("cmd", "cli", "templates", "api", strings.Replace(tplName, ".tpl", "-ws.tpl", 1))
            if _, err := os.Stat(wsTplPath); err == nil {
                templates[i] = strings.Replace(tplName, ".tpl", "-ws.tpl", 1)
            }
        }
    }

	caser := cases.Title(language.English)
    data := struct {
        PackageName string
        Name string
        CreatedAt   string
    }{
        PackageName: strings.ToLower(name),
        Name: caser.String(name),
        CreatedAt:   time.Now().Format("2006/01/02 15:04:05"),
    }

    // --- Generate files ---
    tplDir := filepath.Join("cmd", "cli", "templates", "api")


	for _, tplFile := range templates {
        tplPath := filepath.Join(tplDir, tplFile)
        tpl, err := template.ParseFiles(tplPath)
        if err != nil {
            return fmt.Errorf("failed to parse template %s: %w", tplFile, err)
        }

        // Determine output folder based on template type
        var outFile string
        switch {
        case strings.HasPrefix(tplFile, "init") || strings.HasPrefix(tplFile, "handler") || strings.HasPrefix(tplFile, "route"):
            outFile = filepath.Join(apiBase, fmt.Sprintf("%s.%s.go", name, strings.TrimSuffix(strings.TrimSuffix(tplFile, "-ws.tpl"), ".tpl")))
        case strings.HasPrefix(tplFile, "service"):
            outFile = filepath.Join(serviceBase, fmt.Sprintf("%s.%s.go", name, strings.TrimSuffix(strings.TrimSuffix(tplFile, "-ws.tpl"), ".tpl")))
        case strings.HasPrefix(tplFile, "repository"):
            outFile = filepath.Join(repoBase, fmt.Sprintf("%s.%s.go", name, strings.TrimSuffix(strings.TrimSuffix(tplFile, "-ws.tpl"), ".tpl")))
        case strings.HasPrefix(tplFile, "type"):
            outFile = filepath.Join(typeBase, fmt.Sprintf("%s.%s.go", name, strings.TrimSuffix(strings.TrimSuffix(tplFile, "-ws.tpl"), ".tpl")))
        default:
            return fmt.Errorf("unknown template type: %s", tplFile)
        }

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

    fmt.Printf("🎉 Api '%s' created successfully.\n", name)

	// --- Generate Redux files ---
	if err := GenerateReduxFiles(name); err != nil {
		return fmt.Errorf("redux generation failed: %w", err)
	}

	fmt.Printf("🎨 Redux layer for '%s' created successfully.\n", name)

    return nil
}


func DeleteApi(name string) error {
	caser := cases.Title(language.English)
	nameLower := strings.ToLower(name)
	nameTitle := caser.String(nameLower)

	// Ask for confirmation by typing the name again
	fmt.Printf("You are about to DELETE the API and related files(feature) for '%s'.\n", nameTitle)
	fmt.Printf("Type the name '%s' to confirm: ", nameLower)

	reader := bufio.NewReader(os.Stdin)
	inputRaw, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read confirmation: %w", err)
	}
	input := strings.ToLower(strings.TrimSpace(inputRaw))

	if input != nameLower {
		fmt.Println("Confirmation did not match. Aborting deletion.")
		return fmt.Errorf("confirmation mismatch")
	}

	// --- Backend folders ---
	backendPaths := []string{
		filepath.Join("core", "api", nameLower),
		filepath.Join("core", "services", fmt.Sprintf("%s.service.go", nameLower)),
		filepath.Join("core", "repositories", fmt.Sprintf("%s.repository.go", nameLower)),
		filepath.Join("types", fmt.Sprintf("%s.type.go", nameLower)),
	}

	for _, path := range backendPaths {
		if _, err := os.Stat(path); err == nil {
			if err := os.RemoveAll(path); err != nil {
				return fmt.Errorf("failed to remove %s: %w", path, err)
			}
			fmt.Printf("🗑️  Deleted %s\n", path)
		}
	}

	// --- Frontend (Redux) files ---
	reduxFiles := []string{
		filepath.Join("ui", "src", "cs-redux", "actions", fmt.Sprintf("%s.action.ts", nameLower)),
		filepath.Join("ui", "src", "cs-redux", "slices", fmt.Sprintf("%s.slice.ts", nameLower)),
	}

	for _, path := range reduxFiles {
		if _, err := os.Stat(path); err == nil {
			if err := os.Remove(path); err != nil {
				return fmt.Errorf("failed to remove %s: %w", path, err)
			}
			fmt.Printf("🗑️  Deleted %s\n", path)
		}
	}

	// --- Clean Redux index.ts entries ---
	if err := CleanReduxIndexes(nameTitle, nameLower); err != nil {
		return fmt.Errorf("failed to clean redux indexes: %w", err)
	}

	fmt.Printf("✅ Api '%s' fully removed.\n", name)
	return nil
}

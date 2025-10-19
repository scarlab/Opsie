package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// generateRedux creates Redux action + slice files from templates
func GenerateReduxFiles(name string) error {
	caser := cases.Title(language.English)
	nameLower := strings.ToLower(name)
	nameTitle :=caser.String(name)

	baseActions := filepath.Join("ui", "src", "cs-redux", "actions")
	baseSlices := filepath.Join("ui", "src", "cs-redux", "slices")

	// Ensure folders exist
	if err := os.MkdirAll(baseActions, 0755); err != nil {
		return fmt.Errorf("failed to create actions dir: %w", err)
	}
	if err := os.MkdirAll(baseSlices, 0755); err != nil {
		return fmt.Errorf("failed to create slices dir: %w", err)
	}

	data := struct {
		Name        string
		PackageName string
		CreatedAt   string
	}{
		Name:        nameTitle,
		PackageName: nameLower,
		CreatedAt:   time.Now().Format("2006/01/02 15:04:05"),
	}

	templates := []struct {
		FileName string
		OutPath  string
	}{
		{"action.tpl", filepath.Join(baseActions, fmt.Sprintf("%s.action.ts", nameLower))},
		{"slice.tpl", filepath.Join(baseSlices, fmt.Sprintf("%s.slice.ts", nameLower))},
	}

	for _, tplFile := range templates {
		tplPath := filepath.Join("cmd", "cli", "templates", "redux", tplFile.FileName)
		tpl, err := template.ParseFiles(tplPath)
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", tplFile.FileName, err)
		}

		outFile := tplFile.OutPath
		f, err := os.Create(outFile)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", outFile, err)
		}
		defer f.Close()

		if err := tpl.Execute(f, data); err != nil {
			return fmt.Errorf("failed to execute template %s: %w", tplFile.FileName, err)
		}
		fmt.Printf("✅ Created %s\n", outFile)
	}

	if err := updateReduxIndexes(nameTitle, nameLower); err != nil {
		return fmt.Errorf("failed to update index files: %w", err)
	}

	fmt.Printf("🎉 Redux files for '%s' created successfully.\n", name)
	return nil
}

// updateReduxIndexes updates actions/index.ts and slices/index.ts imports + exports
func updateReduxIndexes(nameTitle, nameLower string) error {
	actionsIndex := filepath.Join("ui", "src", "cs-redux", "actions", "index.ts")
	slicesIndex := filepath.Join("ui", "src", "cs-redux", "slices", "index.ts")

	// --- Update actions/index.ts ---
	if err := updateFile(actionsIndex,
		fmt.Sprintf(`import { %sAction } from "./%s.action";`, nameTitle, nameLower),
		fmt.Sprintf(`    %s: new %sAction(),`, nameLower, nameTitle),
		"export const Actions = {",
	); err != nil {
		return err
	}

	// --- Update slices/index.ts ---
	if err := updateFile(slicesIndex,
		fmt.Sprintf(`import %sSlice from "./%s.slice";`, nameTitle, nameLower),
		fmt.Sprintf(`    %s: %sSlice.reducer,`, nameLower, nameTitle),
		"const CsRootReducer = combineReducers({",
	); err != nil {
		return err
	}

	return nil
}

// updateFile inserts import and object entry if missing
func updateFile(filePath, importLine, objectLine, insertAfter string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", filePath, err)
	}
	text := string(content)

	if !strings.Contains(text, importLine) {
		text = importLine + "\n" + text
	}
	if !strings.Contains(text, objectLine) {
		lines := strings.Split(text, "\n")
		for i, line := range lines {
			if strings.Contains(line, insertAfter) {
				lines = append(lines[:i+1], append([]string{objectLine}, lines[i+1:]...)...)
				break
			}
		}
		text = strings.Join(lines, "\n")
	}

	return os.WriteFile(filePath, []byte(text), 0644)
}



func CleanReduxIndexes(nameTitle, nameLower string) error {
	actionsIndex := filepath.Join("ui", "src", "cs-redux", "actions", "index.ts")
	slicesIndex := filepath.Join("ui", "src", "cs-redux", "slices", "index.ts")

	if err := removeLinesContaining(actionsIndex,
		fmt.Sprintf("import { %sAction }", nameTitle),
		fmt.Sprintf("%s: new %sAction()", nameLower, nameTitle),
	); err != nil {
		return err
	}

	if err := removeLinesContaining(slicesIndex,
		fmt.Sprintf("import %sSlice", nameTitle),
		fmt.Sprintf("%s: %sSlice.reducer", nameLower, nameTitle),
	); err != nil {
		return err
	}

	return nil
}

func removeLinesContaining(filePath string, patterns ...string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", filePath, err)
	}

	lines := strings.Split(string(content), "\n")
	filtered := lines[:0]
	for _, line := range lines {
		skip := false
		for _, p := range patterns {
			if strings.Contains(line, p) {
				skip = true
				break
			}
		}
		if !skip {
			filtered = append(filtered, line)
		}
	}

	return os.WriteFile(filePath, []byte(strings.Join(filtered, "\n")), 0644)
}

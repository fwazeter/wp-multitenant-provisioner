package tenant

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/fwazeter/wp-multitenant-provisioner/internal/config"
)

func CreateTenant(request ProvisionRequest) error {
	tenantPath := filepath.Join(
		config.TenantBasePath,
		request.Domain)

	// Make tenant wp-content dir
	err := os.MkdirAll(filepath.Join(
		tenantPath,
		"wp-content"),
		0755)

	if err != nil {
		return fmt.Errorf("failed to create tenant directory: %w", err)
	}

	// Generate index.php from template
	err = renderTemplate(
		"index.php.tmpl",
		filepath.Join(tenantPath, "index.php"),
		map[string]string{
			"CorePath": "wp",
		})
	if err != nil {
		return fmt.Errorf("failed to render index.php from template: %w", err)
	}

	// Create placeholder files
	stubFiles := []string{
		"wp-config.php",
		"site-config.php",
	}

	for _, f := range stubFiles {
		content := "<?php // stub for " + f
		err := os.WriteFile(filepath.Join(tenantPath, f), []byte(content), 0644)
		if err != nil {
			return fmt.Errorf("failed to write %s: %w", f, err)
		}
	}

	// Symlink wp/ -> WP Core path.
	symlinkTarget := filepath.Join(tenantPath, "wp")
	err = os.Symlink(config.WPCorePath, symlinkTarget)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create symlink: %w", err)
	}

	return nil
}

// Helper function to render Go template files.
func renderTemplate(
	templateFile,
	outputPath string,
	data map[string]string,
) error {
	tmplPath := filepath.Join("internal", "tenant", "templates", templateFile)
	tmpl, err := template.ParseFiles(tmplPath)

	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close()

	return tmpl.Execute(f, data)
}

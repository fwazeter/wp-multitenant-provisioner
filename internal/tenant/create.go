package tenant

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fwazeter/wp-multitenant-provisioner/internal/config"
)

func CreateTenant(request ProvisionRequest) error {
	tenantPath := filepath.Join(
		config.TenantBasePath,
		request.Domain)

	// Make tenant dir
	err := os.MkdirAll(filepath.Join(
		tenantPath,
		"wp-content"),
		0755)

	if err != nil {
		return fmt.Errorf("failed to create tenant directory: %w", err)
	}

	// Create placeholder files
	files := []string{
		"index.php",
		"wp-config.php",
		"site-config.php"}

	for _, f := range files {
		content := "<?php // stub for " + f + "?>"
		err := os.WriteFile(filepath.Join(tenantPath, f), []byte(content), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

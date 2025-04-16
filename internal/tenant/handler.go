package tenant

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ProvisionRequest Struct to decode incoming JSON
type ProvisionRequest struct {
	Domain      string `json:"domain"`
	TablePrefix string `json:"prefix"`
	DBName      string `json:"db"`
}

// HandleProvision Function that handles the POST request
func HandleProvision(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST supported", http.StatusMethodNotAllowed)
		return
	}

	var req ProvisionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request JSON", http.StatusBadRequest)
		return
	}

	err = CreateTenant(req)
	if err != nil {
		http.Error(w, "Failed to create tenant: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Tenant created: %s\n", req.Domain)
}

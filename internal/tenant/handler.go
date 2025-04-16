package tenant

import (
	"encoding/json"
	"net/http"
)

type ProvisionRequest struct {
	Domain      string `json:"domain"`
	TablePrefix string `json:"prefix"`
	DBName      string `json:"db"`
}

func HandleProvision(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ProvisionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = CreateTenant(req)
	if err != nil {
		http.Error(w, "Failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Tenant created: " + req.Domain))
}

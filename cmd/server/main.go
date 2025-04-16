package main

import (
	"fmt"
	"net/http"

	"github.com/fwazeter/wp-multitenant-provisioner"
)

func main() {
	http.HandleFunc("/provision", tenant.HandleProvision)

	fmt.Println("Tenant Provisioning Service is running on http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}

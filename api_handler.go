package main

import (
	"encoding/json"
	"net/http"
)

// HomeHandler /api/v1 controller
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

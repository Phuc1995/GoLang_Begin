package handler

import "net/http"

func AuthenticateRequest(w http.ResponseWriter, r *http.Request) {
	// Redirect the users to login if they’re not authenticated
	authenticated := false
	if !authenticated {
		http.Redirect(w, r, "/register", http.StatusFound)
	}
}

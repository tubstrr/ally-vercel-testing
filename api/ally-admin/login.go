package handler

import (
	"net/http"

	ally "github.com/tubstrr/ally"
)

func AllyAdminLogin(w http.ResponseWriter, r *http.Request) {	
	ally.AdminLogin(w, r)
}
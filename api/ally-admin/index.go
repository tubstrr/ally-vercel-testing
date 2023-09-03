package handler

import (
	"net/http"

	ally "github.com/tubstrr/ally"
)

func AllyAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("cache-control", "no-cache, must-revalidate, private")
  ally.Admin(w, r)
}
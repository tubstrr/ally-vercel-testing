package handler

import (
	"net/http"

	ally "github.com/tubstrr/ally"
	"github.com/tubstrr/ally/network"
)

func AllyAdminLogin(w http.ResponseWriter, r *http.Request) {
	if (network.IsUserLoggedIn(w, r)) {
		w.Header().Set("cache-control", "no-cache, must-revalidate, private")
		w.WriteHeader(http.StatusOK)
	}
  ally.AdminLogin(w, r)
}
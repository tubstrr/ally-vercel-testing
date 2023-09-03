package handler

import (
	"net/http"

	ally "github.com/tubstrr/ally"
	"github.com/tubstrr/ally/network"
)

func AllyAdmin(w http.ResponseWriter, r *http.Request) {
	if (network.IsUserLoggedIn(w, r)) {
		w.Header().Set("cache-control", "no-cache, must-revalidate, private")
		w.WriteHeader(http.StatusOK)
	}
  ally.Admin(w, r)
}
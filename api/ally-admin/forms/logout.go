package handler

import (
	"net/http"

	network "github.com/tubstrr/ally/network"
)

func AllyAdminFormsLogout(w http.ResponseWriter, r *http.Request) {
  network.Logout(w, r)
}
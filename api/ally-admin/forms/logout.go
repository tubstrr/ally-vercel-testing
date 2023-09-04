package handler

import (
	"net/http"

	network "github.com/tubstrr/ally/network"
)

func AllyAdminFormsCreateAccount(w http.ResponseWriter, r *http.Request) {
  network.Logout(w, r)
}
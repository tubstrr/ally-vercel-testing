package handler

import (
	"net/http"

	database "github.com/tubstrr/ally/database"
	network "github.com/tubstrr/ally/network"
)

func AllyAdminFormsCreateAccount(w http.ResponseWriter, r *http.Request) {
	// Check the database
	database.CheckDatabase()
  network.CreateAccount(w, r)
}
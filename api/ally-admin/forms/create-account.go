package handler

import (
	"net/http"

	database "github.com/tubstrr/ally/database"
	network "github.com/tubstrr/ally/network"
)

func AllyAdminFormsCreateAccount(w http.ResponseWriter, r *http.Request) {
	// Check the database
	database.Check_database()
	
	if (network.IsUserLoggedIn(w, r)) {
		w.Header().Set("cache-control", "no-cache, must-revalidate, private")
		w.WriteHeader(http.StatusOK)
	}
  network.CreateAccount(w, r)
}
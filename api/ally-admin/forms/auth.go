package handler

import (
	"net/http"

	database "github.com/tubstrr/ally/database"
	network "github.com/tubstrr/ally/network"
)

func AllyAdminFormsAuth(w http.ResponseWriter, r *http.Request) {
	// Check the database
	database.Check_database()
	
  network.Authorization(w, r)
}
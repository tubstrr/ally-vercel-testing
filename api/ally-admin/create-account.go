package handler

import (
	"net/http"

	ally "github.com/tubstrr/ally"
)

func AllyAdminCreateAccount(w http.ResponseWriter, r *http.Request) {
  ally.AdminCreateAccount(w, r)
}
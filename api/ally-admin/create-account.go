package handler

import (
	"net/http"

	ally "github.com/tubstrr/ally"
)

func AllyAdmin(w http.ResponseWriter, r *http.Request) {
  ally.AdminCreateAccount(w, r)
}
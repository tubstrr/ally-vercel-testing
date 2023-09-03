package handler

import (
	"fmt"
	"net/http"

	ally_theme "github.com/tubstrr/ally-vercel-testing"
)


func Ally(w http.ResponseWriter, r *http.Request) {
  theme := ally_theme.GetTheme()
  // Get the template
	file, err := theme.ReadFile("theme/templates/front-end/index.html")
  if (err != nil) {
    fmt.Println("Error reading file")
		fmt.Println(err)
		// network.FourOhFour(w, r)
		return
	}

  w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "public, s-maxage=86400")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(file))

  // ally.Ally(w, r, theme)
}
package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	if player == "Hass" {
		fmt.Fprint(w, "20")
		return
	}

	if player == "LaughingMan" {
		fmt.Fprint(w, "10")
		return
	}
}

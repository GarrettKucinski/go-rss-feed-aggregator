package handlers

import "net/http"

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, `{ "status": "ok" }`)
}

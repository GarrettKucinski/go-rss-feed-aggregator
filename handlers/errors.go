package handlers

import "net/http"

func HandleErrors(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "something went wrong")
}

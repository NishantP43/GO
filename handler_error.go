package main

import "net/http"

func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, "SOMETHING WENT WRONG !! PLEASE TRY AGAIN !!")

}

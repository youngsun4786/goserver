package main

import (
	"net/http"
	"encoding/json"
)


func handlerValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body	string `json:"body"`
	}

	type returnVals struct {
		Valid	bool   `json:"valid"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)	
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxBodyLength = 140
	if len(params.Body) > maxBodyLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	respondWithJSON(w, http.StatusOK, returnVals{
		Valid: true,
	})
}


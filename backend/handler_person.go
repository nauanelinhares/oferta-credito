package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/nauanelinhares/ofertacredito/internal/database"
)

func (apiCfg *apiConfig) handlerCreatePerson(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Fname   string  `json:"fname"`
		Lname   string  `json:"lname"`
		Age     int32   `json:"age"`
		Email   string  `json:"email"`
		Job     string  `json:"job"`
		Savings float64 `json:"savings"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	person, err := apiCfg.DB.CreatePerson(r.Context(), database.CreatePersonParams{
		ID:        uuid.New(),
		Fname:     params.Fname,
		Lname:     params.Lname,
		Age:       params.Age,
		Email:     params.Email,
		Job:       params.Job,
		Savings:   params.Savings,
		Due:       0.0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating person: %v", err))
		return
	}

	respondWithJSON(w, 201, databasePersonToPerson(person))

}

func (apiCfg *apiConfig) handlerGetPerson(w http.ResponseWriter, r *http.Request) {
	personID := uuid.MustParse(chi.URLParam(r, "id"))
	request, err := apiCfg.DB.GetPerson(r.Context(), personID)

	fmt.Println(request)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting person: %v", err))
		return
	}

	respondWithJSON(w, 200, databasePersonToPerson(request))

}

func (apiCfg *apiConfig) handlerGetPersons(w http.ResponseWriter, r *http.Request) {
	persons, err := apiCfg.DB.GetPersons(r.Context())

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting persons: %v", err))
		return
	}

	respondWithJSON(w, 200, persons)
}

func (apiCfg *apiConfig) handlerDeletePerson(w http.ResponseWriter, r *http.Request) {
	personID := uuid.MustParse(chi.URLParam(r, "id"))

	err := apiCfg.DB.DeletePerson(r.Context(), personID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error deleting person: %v", err))
		return
	}

	respondWithJSON(w, 200, nil)
}

func (apiCfg *apiConfig) handlerUpdatePerson(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Savings float64 `json:"savings"`
	}

	requestID := uuid.MustParse(chi.URLParam(r, "id"))

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// check if request exists
	_, err = apiCfg.DB.GetPerson(r.Context(), requestID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting person: %v", err))
		return
	}

	err = apiCfg.DB.UpdatePerson(r.Context(), database.UpdatePersonParams{
		ID:        requestID,
		Savings:   params.Savings,
		UpdatedAt: time.Now(),
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error updating credit request: %v", err))
		return
	}

	respondWithJSON(w, 200, nil)
}

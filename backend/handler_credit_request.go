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

func (apiCfg *apiConfig) handlerCreateCreditRequest(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		PersonID uuid.UUID `json:"person_id"`
		Amount   float64   `json:"amount"`
		Reason   string    `json:"reason"`
		Note     string    `json:"note"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// check if person exists
	_, err = apiCfg.DB.GetPerson(r.Context(), params.PersonID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting person: %v", err))
		return
	}

	req, err := apiCfg.DB.CreateCreditRequest(r.Context(), database.CreateCreditRequestParams{
		ID:          uuid.New(),
		PersonID:    params.PersonID,
		StartAmount: params.Amount,
		Amount:      params.Amount,
		Status:      "Open",
		Reason:      params.Reason,
		Note:        params.Note,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating person: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseCreditRequestToCreditRequest(req))
}

func (apiCfg *apiConfig) handlerGetCreditRequests(w http.ResponseWriter, r *http.Request) {
	requests, err := apiCfg.DB.GetCreditRequests(r.Context())

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting credit requests: %v", err))
		return
	}

	respondWithJSON(w, 200, requests)
}

func (apiCfg *apiConfig) handlerGetCreditRequest(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.MustParse(chi.URLParam(r, "id"))
	request, err := apiCfg.DB.GetCreditRequest(r.Context(), requestID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting credit request: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseCreditRequestToCreditRequest(request))
}

func (apiCfg *apiConfig) handlerDeleteCreditRequest(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.MustParse(chi.URLParam(r, "id"))

	err := apiCfg.DB.DeleteCreditRequest(r.Context(), requestID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error deleting credit request: %v", err))
		return
	}

	respondWithJSON(w, 200, nil)
}

func (apiCfg *apiConfig) handlerUpdateCreditRequest(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Status string `json:"status"`
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
	_, err = apiCfg.DB.GetCreditRequest(r.Context(), requestID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting credit request: %v", err))
		return
	}

	err = apiCfg.DB.UpdateCreditRequest(r.Context(), database.UpdateCreditRequestParams{
		ID:        requestID,
		Status:    params.Status,
		UpdatedAt: time.Now(),
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error updating credit request: %v", err))
		return
	}

	respondWithJSON(w, 200, nil)

}

package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/nauanelinhares/ofertacredito/internal/database"
)

type person struct {
	ID        uuid.UUID `json:"id"`
	Fname     string    `json:"fname"`
	Lname     string    `json:"lname"`
	Age       int32     `json:"age"`
	Email     string    `json:"email"`
	Job       string    `json:"job"`
	Savings   float64   `json:"savings"`
	Due       float64   `json:"due"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// pedido de credito
type creditRequest struct {
	ID          uuid.UUID `json:"id"`
	PersonID    uuid.UUID `json:"person_id"`
	StartAmount float64   `json:"start_amount"`
	Amount      float64   `json:"amount"`
	Status      string    `json:"status"`
	Reason      string    `json:"reason"`
	Note        string    `json:"note"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func databasePersonToPerson(dbPerson database.Person) person {
	return person{
		ID:        dbPerson.ID,
		Fname:     dbPerson.Fname,
		Lname:     dbPerson.Lname,
		Age:       dbPerson.Age, // Convert int32 to int
		Email:     dbPerson.Email,
		Job:       dbPerson.Job,
		Savings:   dbPerson.Savings,
		Due:       dbPerson.Due,
		CreatedAt: dbPerson.CreatedAt,
		UpdatedAt: dbPerson.UpdatedAt,
	}
}

func databaseCreditRequestToCreditRequest(dbCreditRequest database.Creditrequest) creditRequest {
	return creditRequest{
		ID:          dbCreditRequest.ID,
		PersonID:    dbCreditRequest.PersonID,
		StartAmount: dbCreditRequest.StartAmount,
		Amount:      dbCreditRequest.Amount,
		Status:      dbCreditRequest.Status,
		Reason:      dbCreditRequest.Reason,
		Note:        dbCreditRequest.Note,
		CreatedAt:   dbCreditRequest.CreatedAt,
		UpdatedAt:   dbCreditRequest.UpdatedAt,
	}
}

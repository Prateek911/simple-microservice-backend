package request

import (
	"simple-microservice-backend/db/model"
)

type EmployeeCreate struct {
	Firstname  string             `json:"firstName,omitempty"`
	LastName   string             `json:"LastName,omitempty"`
	MiddleName string             `json:"MiddleName,omitempty"`
	Mnemonic   string             `json:"mnemonic,omitempty"`
	Department model.WorkGroup    `json:"department,omitempty"`
	Role       model.EmployeeRole `json:"role,omitempty"`
	Contact    ContactCreate      `json:"contact,omitempty" validator:"required,dive"`
}

func (e *EmployeeCreate) Validate() error {
	return validate.Struct(e)
}

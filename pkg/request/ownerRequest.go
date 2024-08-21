package request

import (
	"github.com/go-playground/validator/v10"
)

type OwnerCreate struct {
	Name        string             `json:"name,omitempty" validator:"required,alpha,max=255"`
	CRNumber    uint               `json:"crNo,omitempty" validator:"required,alphanum,max=32"`
	Contactable ContactablesCreate `json:"contactable,omitempty" validator:"required,dive"`
}

type ContactablesCreate struct {
	Contact  ContactCreate `json:"contact,omitempty" validator:"required,dive"`
	IsActive bool          `json:"isActive,omitempty" validator:"required,boolean"`
}

type ContactCreate struct {
	Email    string `json:"email,omitempty" validator:"required,alpha,max=255"`
	PhoneNo  uint   `json:"phoneNo,omitempty" validator:"required,numeric,max=10"`
	Location string `json:"location,omitempty" validator:"required,alpha,max=255"`
	Addr1    string `json:"address_1,omitempty" validator:"required,alpha,max=255"`
	Addr2    string `json:"address_2,omitempty"`
	Addr3    string `json:"address_3,omitempty"`
}

var validate = validator.New()

func (o *OwnerCreate) Validate() error {
	return validate.Struct(o)
}

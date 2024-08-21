package response

import "time"

type OwnerResponse struct {
	ClientResponse
	Name        string               `json:"name"`
	CRNumber    uint                 `json:"crNo"`
	Contactable ContactablesResponse `json:"contactable"`
	CreatedAt   time.Time            `json:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt"`
}

type ContactablesResponse struct {
	ClientResponse
	Contact  ContactResponse `json:"contact"`
	IsActive bool            `json:"isActive"`
}

type ContactResponse struct {
	ClientResponse
	Email    string `json:"email"`
	PhoneNo  uint   `json:"phoneNo"`
	Location string `json:"location"`
	Addr1    string `json:"address1"`
	Addr2    string `json:"address2"`
	Addr3    string `json:"address3"`
}

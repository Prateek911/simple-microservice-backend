package response

type OwnerCreate struct {
	Name        string             `json:"name,omitempty"`
	CRNumber    uint               `json:"crNo,omitempty"`
	Contactable ContactablesCreate `json:"contactable,omitempty"`
}
type ContactCreate struct {
	Email    string `json:"email,omitempty"`
	PhoneNo  uint   `json:"phoneNo,omitempty"`
	Location string `json:"location,omitempty"`
	Addr1    string `json:"address_1,omitempty"`
	Addr2    string `json:"address_2,omitempty"`
	Addr3    string `json:"address_3,omitempty"`
}

type ContactablesCreate struct {
	Contact  ContactCreate `json:"contact,omitempty"`
	IsActive bool          `json:"isActive,omitempty"`
}

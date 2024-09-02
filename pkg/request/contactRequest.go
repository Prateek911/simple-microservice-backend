package request

type ContactCreate struct {
	Email    string `json:"email,omitempty" validator:"required,alpha,max=255"`
	PhoneNo  uint   `json:"phoneNo,omitempty" validator:"required,numeric,max=10"`
	Location string `json:"location,omitempty" validator:"required,alpha,max=255"`
	Addr1    string `json:"address_1,omitempty" validator:"required,alpha,max=255"`
	Addr2    string `json:"address_2,omitempty"`
	Addr3    string `json:"address_3,omitempty"`
	IsActive bool   `json:"isActive,omitempty" validator:"required,boolean"`
}

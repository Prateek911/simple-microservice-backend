package response

type ContactResponse struct {
	BaseResponse
	Email    string `json:"email"`
	PhoneNo  uint   `json:"phoneNo"`
	Location string `json:"location"`
	Addr1    string `json:"address1"`
	Addr2    string `json:"address2"`
	Addr3    string `json:"address3"`
	IsActive bool   `json:"isActive"`
}

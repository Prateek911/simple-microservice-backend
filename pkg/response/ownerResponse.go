package response

import "time"

type OwnerResponse struct {
	BaseResponse
	Name      string          `json:"name,omitempty"`
	CRNumber  uint            `json:"crNo,omitempty"`
	Contact   ContactResponse `json:"contact,omitempty"`
	CreatedAt time.Time       `json:"createdAt,omitempty"`
	UpdatedAt time.Time       `json:"updatedAt,omitempty"`
}

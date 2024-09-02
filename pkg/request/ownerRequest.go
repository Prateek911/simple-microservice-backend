package request

type OwnerCreate struct {
	Name     string        `json:"name,omitempty" validator:"required,alpha,max=255"`
	CRNumber uint          `json:"crNo,omitempty" validator:"required,alphanum,max=32"`
	Contact  ContactCreate `json:"contact,omitempty" validator:"required,dive"`
}

func (o *OwnerCreate) Validate() error {
	return validate.Struct(o)
}

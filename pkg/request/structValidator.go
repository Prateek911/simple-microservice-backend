package request

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type Validate interface {
	Validate() error
}

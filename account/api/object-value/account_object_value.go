package objectvalue

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type SignUpAccount struct {
	CompanyName   string
	PhoneNumber   string
	Email         string
	Address       string
	BusinessType  int
	OutletsNumber int
}

func (p SignUpAccount) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.CompanyName, validation.Required, validation.Length(1, 0)),
		validation.Field(&p.PhoneNumber, validation.Required, validation.Length(1, 0)),
		validation.Field(&p.Email, validation.Required, validation.Length(1, 0), is.Email),
		validation.Field(&p.Address, validation.Required, validation.Length(1, 0)),
		validation.Field(&p.BusinessType, validation.Required),
		validation.Field(&p.OutletsNumber, validation.Required),
	)
}

// {
// 	"CompanyName" :   "OK",
// 	"PhoneNumber"  :   "OK",
// 	"Email"         :   "OK",
// 	"Address"       :   "OK",
// 	"BusinessType"  :   "OK",
// 	"OutletsNumber" :   123
// }

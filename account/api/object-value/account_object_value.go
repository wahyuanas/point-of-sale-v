package objectvalue

import validation "github.com/go-ozzo/ozzo-validation"

type SignUpAccount struct {
	CompanyName   string
	PhoneNumber   string
	Email         string
	Address       string
	BusinessType  string
	OutletsNumber int64
}

func (p SignUpAccount) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.CompanyName, validation.Required),
		validation.Field(&p.PhoneNumber, validation.Required),
		validation.Field(&p.Email, validation.Required),
		validation.Field(&p.Address, validation.Required),
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

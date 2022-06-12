package objectvalue

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type SignUp struct {
	CompanyName      string  `json:"companyName"`
	Address          string  `json:"address"`
	Email            string  `json:"email"`
	PhoneNumber      string  `json:"phoneNumber"`
	OutletsNumber    int     `json:"outletsNumber"`
	BusinessType     int     `json:"businessType"`
	MainBusinessType *int    `json:"mainBusinessType"`
	CoreBusinessType *string `json:"coreBusinessType"`
	//Files         []*multipart.File `json:"files"`
}

func (p SignUp) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.CompanyName, validation.Required.Error("Field companyName required"), validation.Length(1, 0).Error("can not be empty")),
		validation.Field(&p.PhoneNumber, validation.Required.Error("Field phoneNumber required"), validation.Length(1, 0).Error("can not be empty")),
		validation.Field(&p.Email, validation.Required.Error("Field email required"), validation.Length(1, 0).Error("can not be empty"), is.Email.Error("wrong email format ")),
		validation.Field(&p.Address, validation.Required.Error("Field address required"), validation.Length(1, 0).Error("can not be empty")),
		validation.Field(&p.BusinessType, validation.Required.Error("Field businessType required")),
		//validation.Field(&p.SubBusinessType, validation., validation.Required.Error("Field subBusinessType required")),
		validation.Field(&p.OutletsNumber, validation.Required.Error("Field outletsNumber required")),
	)

}

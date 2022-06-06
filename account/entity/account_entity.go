package entity

type Account struct {
	ID            int64  `json:"id"`
	CompanyName   string `json:"companyName"`
	Address       string `json:"address"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phoneNumber"`
	OutletsNumber int64  `json:"outletNumber"`
	BusinessType  string `json:"businessType"`
}

package entity

import (
	"github.com/uptrace/bun"
)

type Account struct {
	bun.BaseModel    `bun:"table:tbl_account"`
	ID               int64   `json:"id" bun:",pk,autoincrement"`
	CompanyName      string  `json:"companyName"`
	Address          string  `json:"address"`
	Email            string  `json:"email"`
	PhoneNumber      string  `json:"phoneNumber"`
	OutletsNumber    int     `json:"outletsNumber"`
	BusinessType     int     `json:"businessType"`
	MainBusinessType *int    `json:"mainBusinessType"`
	CoreBusinessType *string `json:"coreBusinessType"`
}

package repository

import (
	"context"
	"net/http"

	"github.com/uptrace/bun"
	"github.com/wahyuanas/point-of-sale-v/account"
	objectvalue "github.com/wahyuanas/point-of-sale-v/account/api/object-value"
	"github.com/wahyuanas/point-of-sale-v/account/entity"
)

type AccountRepository struct {
	DB *bun.DB
}

func NewAccountRepository(db *bun.DB) account.AccountRepository {
	return &AccountRepository{DB: db}
}

func (a *AccountRepository) Store(ctx context.Context, cmd *objectvalue.SignUp) (*entity.Account, interface{}, error) {

	// query := `INSERT tbl_account SET company_name=? , phone_number=? , email=?, address=? , business_type=? , outlets_number=?`
	// stmt, err := a.DB.PrepareContext(ctx, query)
	ec := entity.Account{CompanyName: cmd.CompanyName, PhoneNumber: cmd.PhoneNumber, Email: cmd.Email, Address: cmd.Address, BusinessType: cmd.BusinessType, OutletsNumber: cmd.OutletsNumber}
	_, err := a.DB.NewInsert().Model(&ec).Exec(ctx)

	if err != nil {
		return nil, http.StatusNotAcceptable, err
	}

	//res, err := stmt.ExecContext(ctx, cmd.CompanyName, cmd.PhoneNumber, cmd.Email, cmd.Address, cmd.BusinessType, cmd.OutletsNumber)

	// if err != nil {
	// 	return nil, http.StatusNotAcceptable, err
	// }
	//lastID, err := res.LastInsertId()
	// if err != nil {
	// 	return nil, 500, err
	// }

	// return &entity.Account{
	// 	ID:            lastID,
	// 	CompanyName:   cmd.CompanyName,
	// 	Address:       cmd.Address,
	// 	PhoneNumber:   cmd.PhoneNumber,
	// 	Email:         cmd.Email,
	// 	BusinessType:  cmd.BusinessType,
	// 	OutletsNumber: cmd.OutletsNumber,
	// }, http.StatusCreated, nil

	return &ec, http.StatusCreated, nil
}

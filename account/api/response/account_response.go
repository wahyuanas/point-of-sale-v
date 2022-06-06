package response

import "github.com/wahyuanas/point-of-sale-v/account/entity"

type SignUpAccountResponse struct {
	CommonResponse
	Account *entity.Account
}

type SignInAccountResponse struct {
	CommonResponse
	Account *entity.Account
}

type SignOutAccountResponse struct {
	CommonResponse
	ID int64
}

type UpdateAccountResponse struct {
	CommonResponse
	Account *entity.Account
}

type DeleteAccountResponse struct {
	CommonResponse
	ID int64
}

type GetAccountAccountResponse struct {
	CommonResponse
	Account *entity.Account
}

type GetAccountsAccountResponse struct {
	User []*entity.Account
}

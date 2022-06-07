package response

import "github.com/wahyuanas/point-of-sale-v/account/entity"

type AccountDataResponse struct {
	Account  *entity.Account   `json:"account,omitempty"`
	Accounts []*entity.Account `json:"accounts,omitempty"`
}

type SignUpAccountResponse struct {
	CommonResponse
	DataResponse[AccountDataResponse]
}

type SignInAccountResponse struct {
	CommonResponse
	DataResponse[AccountDataResponse]
}

type SignOutAccountResponse struct {
	CommonResponse
	ID int64
}

type UpdateAccountResponse struct {
	CommonResponse
	DataResponse[AccountDataResponse]
}

type DeleteAccountResponse struct {
	CommonResponse
	ID int64
}

type GetAccountAccountResponse struct {
	CommonResponse
	DataResponse[AccountDataResponse]
}

type GetAccountsAccountResponse struct {
	CommonResponse
	DataResponse[AccountDataResponse]
}

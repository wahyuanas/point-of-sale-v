package response

import "github.com/wahyuanas/point-of-sale-v/account/entity"

type SignUpAccountResponse struct {
	CommonResponse
	DataResponse[struct {
		Account *entity.Account `json:"account"`
	}]
}

type SignInAccountResponse struct {
	CommonResponse
	DataResponse[struct {
		Account *entity.Account `json:"account"`
	}]
}

type SignOutAccountResponse struct {
	CommonResponse
	ID int64
}

type UpdateAccountResponse struct {
	CommonResponse
	DataResponse[struct {
		Account *entity.Account `json:"account"`
	}]
}

type DeleteAccountResponse struct {
	CommonResponse
	ID int64
}

type GetAccountAccountResponse struct {
	CommonResponse
	DataResponse[struct {
		Account *entity.Account `json:"account"`
	}]
}

type GetAccountsAccountResponse struct {
	CommonResponse
	DataResponse[struct {
		Accounts []*entity.Account `json:"accounts"`
	}]
}

package usecase

import (
	"github.com/go-chi/jwtauth"
	"gitlab.com/marcosvto/sys-adv-api/internal/usecase/dtos"
	"gitlab.com/marcosvto/sys-adv-api/pkg/errorable"
)

type ICreateUser interface {
	Execute(input dtos.CreateUserInput) (dtos.UserOutput, error)
}

type IFindUser interface {
	Execute(input dtos.FindInput) (dtos.FindOutput[dtos.UserOutput], error)
}

type ICreateWallet interface {
	Execute(input dtos.CreateWalletInput) (dtos.WalletOutput, error)
}

type IFindWallets interface {
	Execute() ([]dtos.WalletOutput, error)
}

type ILoginUseCase interface {
	Execute(jwt *jwtauth.JWTAuth, jwtExpiresIn int, input dtos.LoginInput) (dtos.LoginOutput, *errorable.CtxError)
}

type ICreateCategory interface {
	Execute(input dtos.CreateCategoryInput) (dtos.CategoryOutput, error)
}

type ICreateTransaction interface {
	Execute(input dtos.CreateTransactionInput) (dtos.TransactionOutput, error)
}

type IFindTransaction interface {
	Execute(input dtos.FindTransactionInput) (dtos.FindOutput[dtos.TransactionOutput], error)
}

type IFindOneTransaction interface {
	Execute(id int) (dtos.FindOutput[dtos.TransactionOutput], error)
}

type IUpdateTransaction interface {
	Execute(id int, input dtos.UpdateTransactionInput) error
}

type IDeleteTransaction interface {
	Execute(id int) error
}

type IFindCategories interface {
	Execute() ([]dtos.CategoryOutput, error)
}

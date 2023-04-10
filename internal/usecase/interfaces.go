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
	Execute(input dtos.FindInput) (dtos.FindOutput, error)
}

type ICreateWallet interface {
	Execute(input dtos.CreateWalletInput) (dtos.CreateWalletOutput, error)
}

type ILoginUseCase interface {
	Execute(jwt *jwtauth.JWTAuth, jwtExpiresIn int, input dtos.LoginInput) (dtos.LoginOutput, *errorable.CtxError)
}

type ICreateCategory interface {
	Execute(input dtos.CreateCategoryInput) (dtos.CategoryOutput, error)
}

package usecase

import (
	"fmt"
	"time"

	"github.com/go-chi/jwtauth"
	"gitlab.com/marcosvto/sys-fin-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
	"gitlab.com/marcosvto/sys-fin-api/pkg/errorable"
)

type LoginUseCase struct {
	UserRepository database.IUserRepository
}

func NewLoginUseCase(userRepository database.IUserRepository) *LoginUseCase {
	return &LoginUseCase{
		UserRepository: userRepository,
	}
}

func (l *LoginUseCase) Execute(jwt *jwtauth.JWTAuth, jwtExpiresIn int, input dtos.LoginInput) (dtos.LoginOutput, *errorable.CtxError) {
	user, err := l.UserRepository.FindByEmail(input.Email)
	if err != nil {
		return dtos.LoginOutput{}, errorable.New(errorable.NOT_FOUND_REGISTER)
	}

	fmt.Println(user.Password)
	if !user.ValidatePassword(input.Password) {
		return dtos.LoginOutput{}, errorable.New(errorable.INVALID_PASSWORD)
	}

	_, tokenString, err := jwt.Encode(map[string]interface{}{
		"sub": user.ID,
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	return dtos.LoginOutput{
		AccessToken: tokenString,
	}, nil
}

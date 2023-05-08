package usecase

import (
	"gitlab.com/marcosvto/sys-adv-api/internal/entity"
	"gitlab.com/marcosvto/sys-adv-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-adv-api/internal/usecase/dtos"
)

type CreateUserUseCase struct {
	UserRepository database.IUserRepository
}

func NewCreateUser(userRepository database.IUserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *CreateUserUseCase) Execute(input dtos.CreateUserInput) (dtos.UserOutput, error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password, input.ConfirmPassword, input.Role)
	if err != nil {
		return dtos.UserOutput{}, err
	}

	if err := uc.UserRepository.Create(user); err != nil {
		return dtos.UserOutput{}, err
	}

	return dtos.UserOutput{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

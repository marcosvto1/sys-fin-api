package usecase

import (
	"gitlab.com/marcosvto/sys-adv-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-adv-api/internal/usecase/dtos"
)

type FindUserUseCase struct {
	UserRepository database.IUserRepository
}

func NewFindUserUseCase(userRepository database.IUserRepository) *FindUserUseCase {
	return &FindUserUseCase{
		UserRepository: userRepository,
	}
}

func (f *FindUserUseCase) Execute(input dtos.FindInput) (dtos.FindOutput, error) {
	offset := (input.PageNumber - 1) * input.PageSize

	users, total, err := f.UserRepository.Find(offset, input.PageSize, input.Id)
	if err != nil {
		return dtos.FindOutput{}, err
	}

	var output []dtos.UserOutput
	for _, user := range users {
		output = append(output, dtos.UserOutput{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		})
	}

	return dtos.FindOutput{
		Items: output,
		Meta: dtos.MetaFindOutput{
			Paginate: dtos.PaginateOutput{
				PageNumber:     input.PageNumber,
				PageSize:       input.PageSize,
				TotalPages:     total / input.PageSize,
				TotalRegisters: total,
			},
		},
	}, nil
}

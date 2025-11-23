package usecase

import (
	"gitlab.com/marcosvto/sys-fin-api/internal/infra/database"
	"gitlab.com/marcosvto/sys-fin-api/internal/usecase/dtos"
)

type FindMovieUC struct {
	movieRepository database.IMovieRepository
}

func NewFindMovieUC(movieRepository database.IMovieRepository) *FindMovieUC {
	return &FindMovieUC{
		movieRepository: movieRepository,
	}
}

func (uc *FindMovieUC) Execute(input dtos.FindMovieInput) (dtos.FindOutput[dtos.MovieOutput], error) {
	offset := (input.PageNumber - 1) * input.PageSize

	movies, count, err := uc.movieRepository.Find(offset, input.PageSize, database.FindMovieOptions{
		Title:   input.Title,
		Watched: input.Watched,
	})

	if err != nil {
		return dtos.FindOutput[dtos.MovieOutput]{}, err
	}

	var items []dtos.MovieOutput
	for _, movie := range movies {
		items = append(items, dtos.MovieOutput{
			ID:          movie.ID,
			Title:       movie.Title,
			Description: movie.Description,
			ImageUrl:    movie.ImageUrl,
			Watched:     movie.Watched,
		})
	}

	if len(items) == 0 {
		return dtos.FindOutput[dtos.MovieOutput]{
			Items: []dtos.MovieOutput{},
		}, nil
	}

	return dtos.FindOutput[dtos.MovieOutput]{
		Items: items,
		Meta: dtos.MetaFindOutput{
			Paginate: dtos.PaginateOutput{
				PageNumber:     input.PageNumber,
				PageSize:       input.PageSize,
				TotalPages:     (count + input.PageSize - 1) / input.PageSize,
				TotalRegisters: count,
			},
		},
	}, nil

}

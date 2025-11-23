package usecase

import "gitlab.com/marcosvto/sys-fin-api/internal/infra/database"

type DeleteMovieUC struct {
	movieRepository database.IMovieRepository
}

func NewDeleteMovieUC(movieRepository database.IMovieRepository) *DeleteMovieUC {
	return &DeleteMovieUC{
		movieRepository: movieRepository,
	}
}

func (uc *DeleteMovieUC) Execute(id int) error {
	_, err := uc.movieRepository.FindById(id)
	if err != nil {
		return err
	}

	err = uc.movieRepository.DeleteById(id)
	if err != nil {
		return err
	}

	return nil
}

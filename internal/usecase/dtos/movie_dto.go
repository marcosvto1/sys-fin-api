package dtos

type MovieOutput struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Watched     bool   `json:"watched"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}

type CreateMovieInput struct {
	Title       string `json:"title" validate:"required"`
	Watched     bool   `json:"watched" default:"false"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}

type UpdateMovieInput struct {
	Title       string `json:"title" validate:"required"`
	Watched     bool   `json:"watched" validate:"required"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}

type UpdateMovieOutput MovieOutput

type FindMovieInput struct {
	Title   string `json:"title" validate:"required"`
	Watched string `json:"watched"`

	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
}

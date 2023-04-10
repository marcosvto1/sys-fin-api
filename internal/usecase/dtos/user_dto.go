package dtos

type CreateUserInput struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Role            string `json:"role"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UserOutput struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type FindInput struct {
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
	Id         int `json:"id"`
}

type PaginateOutput struct {
	PageNumber     int `json:"page_number"`
	PageSize       int `json:"page_size"`
	TotalPages     int `json:"total_pages"`
	TotalRegisters int `json:"total_registers"`
}

type MetaFindOutput struct {
	Paginate PaginateOutput `json:"paginate"`
}

type FindOutput struct {
	Items []UserOutput   `json:"items"`
	Meta  MetaFindOutput `json:"meta"`
}

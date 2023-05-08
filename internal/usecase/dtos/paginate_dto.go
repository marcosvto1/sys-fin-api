package dtos

type PaginateOutput struct {
	PageNumber     int `json:"page_number"`
	PageSize       int `json:"page_size"`
	TotalPages     int `json:"total_pages"`
	TotalRegisters int `json:"total_registers"`
}

type MetaFindOutput struct {
	Paginate PaginateOutput `json:"paginate"`
}

type FindOutput[V UserOutput | CategoryOutput | TransactionOutput] struct {
	Items []V            `json:"items"`
	Meta  MetaFindOutput `json:"meta"`
}

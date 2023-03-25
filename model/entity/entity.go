package entity

type Pagination struct {
	Total   int64 `json:"total"`
	Page    int64 `json:"page"`
	PerPage int64 `json:"per_page"`
}

type Index struct {
	Rows interface{} `json:"rows"`
	Pagination
}

type PaginationRequests struct {
	Page    int64 `json:"page" query:"page" default:"1"`
	PerPage int64 `json:"per_page" query:"per_page" default:"10"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ResponseError struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
}

type ErrorResponse struct {
	FailedField string `json:"failedField"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

type Response200 struct {
	Success bool        `json:"success" example:"true"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message" example:"Success"`
}

type Response400 struct {
	Success bool        `json:"success" example:"false"`
	Message string      `json:"message" example:"Bad Request"`
	Error   interface{} `json:"error"`
}

type Response500 struct {
	Success bool        `json:"success" example:"false"`
	Message string      `json:"message" example:"Internal Server Error"`
	Error   interface{} `json:"error"`
}

type Sort struct {
	Field string `json:"field"`
	By    int64  `json:"by"`
}

type SortRequests struct {
	SortField string `query:"sort_field" json:"sort_field"`
	SortBy    int64  `query:"sort_by" json:"sort_by"`
}

type CountAggregate struct {
	Count int `json:"count"`
}

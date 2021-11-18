package helper

type Pagination struct {
	Total int `json:"total"`
	Count int `json:"count"`
	PerPage int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	TotalPage int `json:"total_page"`
}

type Meta struct {
	Message string `json:"message"`
	Code int `json:"code"`
	Status string `json:"status"`
	Pagination Pagination `json:"pagination"`
}

type Response struct {
	Data interface{} `json:"data"`
	Meta Meta `json:"meta"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}
	return jsonResponse
}

package responders

import "net/http"

type PaginationResponse struct {
	Payload []interface{}  `json:"payload"`
	Meta    PaginationMeta `json:"meta"`
}

type PaginationMeta struct {
	PaginationShortMeta
}

type PaginationShortMeta struct {
	Total        int `json:"total"`
	Count        int `json:"count"`
	Start        int `json:"start"`
	CountRecords int `json:"countRecords"`
}

func Pagination(w http.ResponseWriter, data []interface{}, meta PaginationShortMeta) {
	render.JSON(w, 200, PaginationResponse{
		Payload: data,
		Meta: PaginationMeta{
			meta,
		},
	})
}

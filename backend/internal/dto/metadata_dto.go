package dto

type Metadata struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalPages int64 `json:"total_pages"`
	TotalItems int64 `json:"total_items"`
}

func NewMetadata(page, limit, totalItems int64) Metadata {
	return Metadata{
		Page:       page,
		Limit:      limit,
		TotalPages: (totalItems / limit) + 1,
		TotalItems: totalItems,
	}
}

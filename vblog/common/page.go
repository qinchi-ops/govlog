package common

func NewPageRequest() *PageRequest {
	return &PageRequest{
		PageSize:   10,
		PageNumber: 1,
	}
}

type PageRequest struct {
	// 分页的大小
	PageSize int `json:"page_size"`
	// 当前页码
	PageNumber int `json:"page_number"`
}

func (req *PageRequest) Offset() int {
	return (req.PageNumber - 1) * req.PageSize
}

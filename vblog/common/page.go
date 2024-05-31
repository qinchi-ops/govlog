package common

func NewPageRequest() *PageRequest {
	return &PageRequest{
		PageSize:   10,
		PageNumber: 1,
	}
}

type PageRequest struct {
	// 分页的大小
	PageSize int
	// 当前页码
	PageNumber int
}

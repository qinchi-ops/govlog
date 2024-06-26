package common

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewPageRequest() *PageRequest {
	return &PageRequest{
		PageSize:   10,
		PageNumber: 1,
	}
}

func NewPageRequestFromGin(c *gin.Context) *PageRequest {
	p := NewPageRequest()
	pnstr := c.Query("page_number")
	psstr := c.Query("page_size")
	if pnstr != "" {
		pn, _ := strconv.ParseInt(pnstr, 10, 64)
		if pn != 0 {
			p.PageNumber = int(pn)
		}
	}

	if psstr != "" {
		ps, _ := strconv.ParseInt(psstr, 10, 64)
		if ps != 0 {
			p.PageSize = int(ps)
		}
	}
	return p
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

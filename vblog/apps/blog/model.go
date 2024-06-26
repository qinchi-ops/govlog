package blog

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/qinchi-ops/govlog/vblog/common"
)

//	func NewBlog() *Blog {
//		return &Blog{
//			common.NewMeta(),
//			&CreateBlogRequest{
//				Tags: map[string]string{},
//			},
//			&ChangedBlogStatusRequest{},
//		}
//	}
func NewBlog() *Blog {
	return &Blog{
		&Meta{
			CreatedAt: time.Now().Unix(),
		},
		&CreateBlogRequest{
			Tags: map[string]string{},
		},
		&ChangedBlogStatusRequest{},
	}
}

type Blog struct {
	*Meta
	*CreateBlogRequest
	*ChangedBlogStatusRequest
}

func (req *Blog) String() string {
	dj, _ := json.MarshalIndent(req, "", "	")
	return string(dj)
}

// func (req *Blog) String() string {
// 	dj, _ := json.Marshal(req, "", " ")
// 	fmt.Print("findis")
// 	return string(dj)

// }

type CreateBlogRequest struct {
	// 文章标题
	Title string `json:"title" gorm:"column:title" validate:"required"`
	// 作者
	Author string `json:"author" gorm:"column:author" validate:"required"`
	// 文章内容
	Content string `json:"content" gorm:"column:content" validate:"required"`
	// 文章概要信息
	Summary string `json:"summary" gorm:"column:summary"`
	// 创建人
	CreateBy string `json:"create_by" gorm:"column:create_by"`
	// 标签 https://gorm.io/docs/serializer.html
	Tags map[string]string `json:"tags" gorm:"column:tags;serializer:json"`
}

func (req *CreateBlogRequest) Validate() error {
	return common.Validate(req)
}

func (req *CreateBlogRequest) String() string {
	dj, _ := json.MarshalIndent(req, "", "	")
	fmt.Print("findis")
	return string(dj)

}

type ChangedBlogStatusRequest struct {
	// 发布时间
	PublishedAt int64 `json:"published_at" gorm:"column:published_at"`
	// 文章状态: 草稿/已发布
	Status Status `json:"status" gorm:"column:status"`
}

func (req *ChangedBlogStatusRequest) String() string {
	dj, _ := json.MarshalIndent(req, "", "	")
	fmt.Print("findis")
	return string(dj)

}

type Meta struct {
	// 用户Id
	Id int `json:"id" gorm:"column:id"`
	// 创建时间, 时间戳 10位, 秒
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	// 更新时间, 时间戳 10位, 秒
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`
}

type BlogSet struct {
	Total int64   `json:"total"`
	Items []*Blog `json:"items"`

	// common.Print
}

func (req *BlogSet) String() string {
	dj, _ := json.MarshalIndent(req, "", " ")
	return string(dj)
}

func NewBlogSet() *BlogSet {
	return &BlogSet{Items: []*Blog{}}
}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{
		Tags: map[string]string{},
	}
}

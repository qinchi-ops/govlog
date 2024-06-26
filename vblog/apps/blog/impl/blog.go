package impl

import (
	"context"
	"fmt"

	"github.com/qinchi-ops/govlog/vblog/apps/blog"
	"github.com/qinchi-ops/govlog/vblog/exception"
)

// 文章创建
func (i *BlogServiceImpl) CreateBlog(ctx context.Context, in *blog.CreateBlogRequest) (*blog.Blog, error) {

	// 1. 验证请求参数
	if err := in.Validate(); err != nil {
		return nil, exception.ErrValidateFailed(err.Error())
	}

	// 2. 构造实例对象
	ins := blog.NewBlog()
	ins.CreateBlogRequest = in
	fmt.Println(ins)
	fmt.Println(in)
	// 3. 入库返回
	// INSERT INTO `blogs` (`created_at`,`updated_at`,`title`,`author`,`content`,`summary`,`create_by`,`tags`,`published_at`,`status`) VALUES (1717832260,1717832260,'Go全站开发','will','Md内容填充','文章概要信息','','{}',0,0)
	err := i.db.WithContext(ctx).Create(ins).Error
	if err != nil {
		return nil, err
	}

	return ins, nil
}
func (i *BlogServiceImpl) QueryBlog(ctx context.Context, in *blog.QueryBlogRequest) (*blog.BlogSet, error) {
	// 1.因为有默认值，不需要用户传参数
	set := &blog.BlogSet{}
	// 2.直接查数据库，构造查询语句
	query := i.db.WithContext(ctx).Table("blogs")
	if in.KeyWords != "" {
		query = query.Where("title LIKE ?", "%"+in.KeyWords+"%")
	}
	// count 总数统计
	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}
	// 3. 具体的查询
	err = query.Limit(in.PageSize).Offset(in.Offset()).Find(&set.Items).Error
	if err != nil {
		return nil, err
	}
	return set, nil
}

// 文章详情
func (i *BlogServiceImpl) DescribeBlog(ctx context.Context, in *blog.DescribeBlogRequest) (*blog.Blog, error) {
	ins := blog.NewBlog()
	err := i.db.WithContext(ctx).First(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// 文章更新
func (i *BlogServiceImpl) UpdateBlog(ctx context.Context, in *blog.UpdateBlogRequest) (*blog.Blog, error) {
	return nil, nil
}

// 文章删除
func (i *BlogServiceImpl) DeleteBlog(ctx context.Context, in *blog.DeleteBlogRequest) (*blog.Blog, error) {
	return nil, nil
}

func (i *BlogServiceImpl) UpdateBlogStatus(ctx context.Context, in *blog.ChangedBlogStatusRequest) (*blog.Blog, error) {
	return nil, nil
}

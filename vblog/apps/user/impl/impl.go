package impl

import (
	"context"

	"github.com/qinchi-ops/govlog/vblog/apps/user"
	"github.com/qinchi-ops/govlog/vblog/common"
	"github.com/qinchi-ops/govlog/vblog/conf"
	"github.com/qinchi-ops/govlog/vblog/ioc"
	"gorm.io/gorm"
)

// func NewUserServiceImpl() *UserServiceImpl {
// 	// 每个业务对象,都可能依赖到数据库
// 	// db =create db
// 	return &UserServiceImpl{
// 		db: *conf.C().Mysql.GetDB(),
// 	}
// }

// import _ ----> init方来来注册 包里面的核心对象
func init() {
	ioc.Controller.Registry(user.AppName, &UserServiceImpl{})
}

func (i *UserServiceImpl) Init() error {
	i.db = conf.C().Mysql.GetDB()
	return nil
}

// 需要资源
// 需要数据库操作
type UserServiceImpl struct {
	// db conn 共享对象
	// mysql host port  ....
	db *gorm.DB
}

func (i *UserServiceImpl) CreateUser(ctx context.Context, in *user.CreateUserRequest) (*user.User, error) {
	// 1.校验请求的合法性
	if err := common.Validate(in); err != nil {
		return nil, err
	}
	//不要存储用户铭文密码
	if err := in.HashPassword(); err != nil {
		return nil, err

	}
	// 2. 创建user对象(资源)
	ins := user.NewUser(in)
	// 3. user对象保持入库
	/*
		读取数据库配置
		读取数据库连接
		操作连接 保证数据
	*/
	if err := i.db.Save(ins).Error; err != nil {
		return nil, err
	}
	// 4. 返回保持后的user对象
	return ins, nil
}

func (i *UserServiceImpl) QueryUser(
	ctx context.Context,
	in *user.QueryUserRequest) (
	*user.UserSet, error) {
	// 构造一个查询语句，TableName() select
	// WithContext
	set := user.NewUserSet()
	query := i.db.Model(&user.User{}).WithContext(ctx)
	// Where where username = ?
	if in.Username != "" {
		//注意：返回一个新的对象,并没有直接修改这个对象
		//func (db *DB) Where(query interface{}, args ...interface{}) (tx *DB) {」
		query = query.Where("username = ?", in.Username)
	}
	// 怎么查询Total, 需要把过滤条件: username ,key
	// 查询Total时能不能把分页参数带上
	// select COUNT(*) from xxx limit 10
	// select COUNT(*) from xxx
	// 不能携带分页参数
	if err := query.Count(&set.Total).Error; err != nil {
		return nil, err
	}
	if err := query.Offset(in.Offset()).Limit(in.PageSize).Find(&set.Items).Error; err != nil {
		return nil, err
	}
	return set, nil
}

func (i *UserServiceImpl) DeleteUser(ctx context.Context, in *user.DeleteUserRequest) (*user.UserSet, error) {

	set := user.NewUserSet()
	query := i.db.Model(&user.User{}).WithContext(ctx)
	// Where where username = ?
	if in.Username != "" {

		query = query.Where("username = ?", in.Username).Table("users").Delete(&user.User{})
	}
	if err := query.Count(&set.Total).Error; err != nil {
		return nil, err
	}
	if err := query.Offset(in.Offset()).Limit(in.PageSize).Find(&set.Items).Error; err != nil {
		return nil, err
	}
	return set, nil

}

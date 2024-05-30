package blog

// 构造函数,对对象进行初始化,同意对对象的初始者进行管理
// 保证构造的对象是可用的,不容易出现nil
type Status int

const (
	// 草稿
	STATUS_DRAFT = iota
	// 已发布
	STATUS_PUBLISH
)

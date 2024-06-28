package user

type Role int

// 通过定义满足类型的常量，来定义满足这个类型的列表
// ROLE_MEMBER/ROLE_ADMIN
const (
	// 当 值为0的时候, 就是默认值
	// 枚举命名风格:
	ROLE_ADMIN Role = iota
	ROLE_VISITOR
	ROLE_AUTHOR
)

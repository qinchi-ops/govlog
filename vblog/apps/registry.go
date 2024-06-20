package apps

import (
	// 业务对象注册
	_ "github.com/qinchi-ops/govlog/vblog/apps/blog"
	_ "github.com/qinchi-ops/govlog/vblog/apps/token/api"
	_ "github.com/qinchi-ops/govlog/vblog/apps/token/impl"
	_ "github.com/qinchi-ops/govlog/vblog/apps/user/impl"
)

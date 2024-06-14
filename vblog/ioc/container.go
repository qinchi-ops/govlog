package ioc

// Controller 是一个Container, 使用MapContainer 的实现
var Controller Container = &MapContainer{
	name: "controller",
	// [], 自定义结构
	storage: make(map[string]Object),
}

// Api 所有的对外接口对象都放这里

var Api Container = &MapContainer{
	// [], 自定义结构
	name:    "api",
	storage: make(map[string]Object),
}

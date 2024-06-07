package ioc

// Controller 是一个Container, 使用MapContainer 的实现
var Controller Container = &MapCotainer{
	// [], 自定义结构
	storage: make(map[string]Object),
}

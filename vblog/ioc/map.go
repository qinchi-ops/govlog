package ioc

import "fmt"

// Map　类型的IocContainer
type MapContainer struct {
	name    string
	storage map[string]Object
}

//注册对象

func (c *MapContainer) Registry(name string, obj Object) {
	c.storage[name] = obj
}

//获取对象

func (c *MapContainer) Get(name string) any {
	return c.storage[name]

}

// 调用所有被托管对象的Init()，对对象进行初始化
func (c *MapContainer) Init() error {
	for k, v := range c.storage {
		if err := v.Init(); err != nil {
			return fmt.Errorf("%s init failed,%s", k, err)
		}
		fmt.Printf("%s init success", k)
	}
	return nil

}

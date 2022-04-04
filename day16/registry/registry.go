package registry

import "context"

type Registry interface {
	// 插件名字
	Name() string
	// 初始化
	Init(ctx context.Context, opts ...Option) (err error)
	// 服务注册
	Regiter(ctx context.Context, service *Service) (err error)
	// 服务反注册
	UnRegiter(ctx context.Context, service *Service) (err error)
	// 服务发现
	GetService(ctx context.Context, name string) (service *Service, err error)
}

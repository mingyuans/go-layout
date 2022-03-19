package apiserver

//这里配置 Redis, Spanner 等等额外的服务组件；
type componentConfig struct {
}

type completedComponentConfig struct {
	*componentConfig
}

func (c *componentConfig) complete() *completedComponentConfig {
	return &completedComponentConfig{c}
}

func (c *completedComponentConfig) New() error {
	return nil
}

func newComponentConfig() *componentConfig {
	return &componentConfig{}
}

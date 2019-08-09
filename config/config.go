package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
)

func main() *toml.Tree {
	cfg, err := toml.LoadFile("./config/debug.toml")
	if err != nil {
		fmt.Println("获取配置错误", err)
	}
	return cfg
}

var (
	GetConfig = main()
)

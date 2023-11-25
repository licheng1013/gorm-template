package app

import "gopkg.in/yaml.v3"

// Config 全局配置变量
var Config *SystemConfig

type SystemConfig struct {
	// Mysql地址
	MysqlUrl string `yaml:"mysqlUrl"`
	// 排除地址
	ExcludePath []string `yaml:"excludePath"`
	// 端口
	Port int `yaml:"port"`
	// 上传文件
	UploadPath string `yaml:"uploadPath"`
}

func ParseAppConfig(data []byte) {
	err := yaml.Unmarshal(data, &Config)
	if err != nil {
		panic("解析配置失败!")
	}
}

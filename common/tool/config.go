package tool

import (
	"gopkg.in/yaml.v3"
	"os"
)

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
	// 上传路径
	UploadUrl string `yaml:"uploadUrl"`
}

func ReadConfig(filePath string) {
	MyLog.Println("读取配置")
	data, _ := os.ReadFile(filePath)
	err := yaml.Unmarshal(data, &Config)
	if err != nil {
		panic("解析配置失败!")
	}
}

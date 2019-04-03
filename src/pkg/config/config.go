package config

import (
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/sirupsen/logrus"
)

type Config struct {
	HttpServer HttpServer
	Mysql      Mysql
	RabbitMQ   RabbitMQ
}

type HttpServer struct {
	Port int
	Mode string
	Tag  string
}

type Mysql struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

type RabbitMQ struct {
	Host      string
	Port      int
	Username  string
	Password  string
	QueueName string
}

// config单例对象
var config *Config

// 创建配置对象
func InitConfig(confPath string) error {
	config = &Config{}
	err := loadFile(confPath, config)
	if err != nil {
		logrus.Errorf("loadFile error, res: %v", err)
		return err
	}
	return nil
}

// 获取配置
func GetConfig() *Config {
	if config == nil{
		logrus.Fatalln("config is nil, need to InitConfig")
	}
	return config
}

// 加载配置
func loadFile(path string, cfg interface{}) error {
	if file, err := os.Open(path); err != nil {
		return err
	} else {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err != nil {
			return err
		} else {
			return yaml.Unmarshal(data, cfg)
		}
	}
}

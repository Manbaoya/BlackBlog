package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Server struct { //服务器信息
	Port    string
	Address string
}
type DataBase struct {
	User      string `mapstructure:"User"`
	Password  string `mapstructure:"Password"`
	Name      string `mapstructure:"Name"`
	Port      string `mapstructure:"Port"`
	Charset   string `mapstructure:"Charset"`
	ParseTime string `mapstructure:"ParseTime"`
	Loc       string `mapstructure:"Loc"`
	Limit     int    `mapstructure:"Limit"`
}
type Log struct {
	MaxSize    int  `mapstructure:"Maxsize"`
	MaxAge     int  `mapstructure:"MaxAge"`
	MaxBackups int  `mapstructure:"MaxBackups"`
	Compress   bool `mapstructure:"Compress"`
}
type AppConfig struct {
	Server   Server   `mapstructure:"Server"`
	DataBase DataBase `mapstructure:"DataBase"`
	Log      Log      `mapstructure:"Log"`
}

const (
	Path = "D:/GolandProjects/BLACKBLOG/config/config.yaml"
)

var Conf = new(AppConfig)

func GetConfig() {
	viper.SetConfigFile(Path) //指定配置路径

	if err := ReadConfig(Conf); err != nil {
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("unmarshal conf failed, err:%s \n", err)
		}
	})
}

// 读取配置
func ReadConfig(Conf *AppConfig) error {
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("read file failed error: %s \n", err)
		return err

	}

	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("unmarshal error: %s \n", err)
		return err
	}

	return nil
}

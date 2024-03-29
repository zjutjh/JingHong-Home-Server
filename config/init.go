package config

import (
	"log"

	"github.com/spf13/viper"
)

type server struct {
	Port         string   `mapstructure:"port"`
	AllowOrigins []string `mapstructure:"allow_origins"`
}

type database struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Address  string `mapstructure:"address"`
	DbName   string `mapstructure:"db_name"`
}
type email struct {
	Sender   string `mapstructure:"sender"`
	Pwd      string `mapstructure:"pwd"`
	SmtpAddr string `mapstructure:"smtp_addr"`
	SmtpPort int    `mapstructure:"smtp_port"`
	Receiver string `mapstructure:"receiver"`
}

type ConfigData struct {
	Server   server   `mapstructure:"server"`
	Database database `mapstructure:"database"`
	Secret   string   `mapstructure:"secret"`
	Secret2  string   `mapstructure:"secret2"`
	Dev      bool     `mapstructure:"dev"`
	Email    email    `mapstructure:"email"`
}

var Config ConfigData

func init() {
	var config = viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("./config")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatalln("Config Error: ", err)
		panic(err)
	}
	config.Unmarshal(&Config)
}

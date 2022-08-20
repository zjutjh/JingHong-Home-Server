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
	Receiver string `mapstructure:"receiver"`
	Pwd      string `mapstructure:"pwd"`
	SmtpAddr string `mapstructure:"smtp_addr"`
	SmtpPort int    `mapstructure:"smtp_port"`
}
type jwt struct {
	Issuer  string `mapstructure:"issuer"`
	Secret  string `mapstructure:"secret"`
	Expires int    `mapstructure:"expires"`
}
type ConfigData struct {
	Server   server   `mapstructure:"server"`
	Database database `mapstructure:"database"`
	Secret   string   `mapstructure:"secret"`
	Secret2  string   `mapstructure:"secret2"`
	Email    email    `mapstructure:"email"`
	Dev      bool     `mapstructure:"dev"`
	Jwt      jwt      `mapstructure:"jwt"`
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

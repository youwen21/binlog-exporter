package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type OptsOfSMTP struct {
	SmtpHost string `toml:"smtp_host" mapstructure:"smtp_host"`
	Port     int    `toml:"smtp_port" mapstructure:"smtp_port"`
	User     string `toml:"smtp_user" mapstructure:"smtp_user"`
	Password string `toml:"smtp_password" mapstructure:"smtp_password"`
}

type Mysql struct {
	ServerId int    `toml:"server_id" mapstructure:"server_id"`
	Host     string `mapstructure:"db_host"`
	Username string `mapstructure:"db_username"`
	Password string `mapstructure:"db_password"`
	Port     int    `mapstructure:"db_port"`
	Database string `mapstructure:"db_database"`
	Charset  string `mapstructure:"db_charset"`
}

type AppConfig struct {
	AppEnv        string     `mapstructure:"app_env"`
	ListenAddress string     `mapstructure:"listen_address"`
	TelemetryPath string     `mapstructure:"telemetry_path"`
	DefaultMysql  Mysql      `mapstructure:",squash"`
	Smtp          OptsOfSMTP `mapstructure:",squash"`
}

var Config *AppConfig

func init() {
	//viper.AutomaticEnv()
	viper.SetConfigName("config") // name of config file (without extension)  config.type{env, yaml, toml , etc...}
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	//settings := viper.AllSettings()
	//fmt.Println(settings)

	err = viper.Unmarshal(&Config)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("viper Unmarshal error: %w \n", err))
	}
}

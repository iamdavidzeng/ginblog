package setting

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

var (
	// RunMode 1
	RunMode string

	// HTTPPort 1
	HTTPPort int
	// ReadTimeout 1
	ReadTimeout time.Duration
	// WriteTimeout 1
	WriteTimeout time.Duration

	// PageSize define
	PageSize int
	// JwtSecret represent something
	JwtSecret string
)

type Config struct {
	RunMode  string   `mapstructure:"RUN_MODE"`
	App      App      `mapstructure:"APP"`
	Server   Server   `mapstructure:"SERVER"`
	Database Database `mapstructure:"DATABASE"`
}

type App struct {
	PageSize  int    `mapstructure:"PAGE_SIZE"`
	JwtSecret string `mapstructure:"JWT_SECRET"`
}

type Server struct {
	HTTPPort     int `mapstructure:"HTTP_PORT"`
	ReadTimeout  int `mapstructure:"READ_TIMEOUT"`
	WriteTimeout int `mapstructure:"WRITE_TIMEOUT"`
}

type Database struct {
	Type        string `mapstructure:"TYPE"`
	User        string `mapstructure:"USER"`
	Password    string `mapstructure:"PASSWORD"`
	Host        string `mapstructure:"HOST"`
	Port        int    `mapstructure:"PORT"`
	Name        string `mapstructure:"NAME"`
	TablePrefix string `mapstructure:"TABLE_PREFIX"`
}

var Cfg Config

func init() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	viper.Unmarshal(&Cfg)

	LoadBase()
	LoadApp()
	LoadServer()
}

// LoadBase parse base config from ini file
func LoadBase() {
	RunMode = Cfg.RunMode
}

// LoadServer parse server config
func LoadServer() {
	HTTPPort = Cfg.Server.HTTPPort
	ReadTimeout = time.Duration(Cfg.Server.ReadTimeout) * time.Second
	WriteTimeout = time.Duration(Cfg.Server.WriteTimeout) * time.Second
}

// LoadApp parse app config
func LoadApp() {
	JwtSecret = Cfg.App.JwtSecret
	PageSize = Cfg.App.PageSize
}

package setting

import (
	"fmt"
	"os"
	"regexp"
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
	RunMode  string `mapstructure:"RUN_MODE"`
	App      App    `mapstructure:"APP"`
	Server   Server `mapstructure:"SERVER"`
	Database struct {
		Type        string `mapstructure:"TYPE"`
		TablePrefix string `mapstructure:"TABLE_PREFIX"`
		URI         string `mapstructure:"URI"`
	} `mapstructure:"DB"`
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

var Cfg Config

func init() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	for _, key := range viper.AllKeys() {
		value := viper.GetString(key)
		envValueOrDefault := Parse(value)
		viper.Set(key, envValueOrDefault)
	}

	err := viper.Unmarshal(&Cfg)
	if err != nil {
		panic(fmt.Errorf("Fatal error load config: %s", err))
	}

	// Load config to constant
	LoadBase()
	LoadApp()
	LoadServer()
}

// Parse use to get ENV value or just default value
func Parse(str string) string {
	search := regexp.MustCompile(`\$\{([^}:]+):?([^}]+)?\}`)
	replacedBody := search.ReplaceAllFunc([]byte(str), func(b []byte) []byte {
		group1 := search.ReplaceAllString(string(b), `$1`)
		group2 := search.ReplaceAllString(string(b), `$2`)

		envValue := os.Getenv(group1)
		if len(envValue) > 0 {
			return []byte(envValue)
		}
		return []byte(group2)
	})

	return string(replacedBody)
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

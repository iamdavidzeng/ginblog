package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	// Cfg use ini.File object get config
	Cfg *ini.File

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

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

// LoadBase parse base config from ini file
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

// LoadServer parse server config
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

// LoadApp parse app config
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("123456")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

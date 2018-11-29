package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	/*LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string*/
	LogLevel string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var cfg *ini.File

func Setup(runmode string) {
	var err error
	switch runmode {
	case "dev":
		cfg, err = ini.Load("conf/dev.ini")
	case "test":
		cfg, err = ini.Load("conf/test.ini")
	case "pre":
		cfg, err = ini.Load("conf/pre.ini")
	case "prod":
		cfg, err = ini.Load("conf/prod.ini")
	default:
		logrus.Fatal("invalid runmode,must be one of [dev,test,pre,prod]!")
	}

	if err != nil {
		log.Fatalf("Fail to parse config file: %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}

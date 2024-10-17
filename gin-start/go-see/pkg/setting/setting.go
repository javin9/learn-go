package setting

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
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

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  int
	WriteTimeout int
}

var ServerSetting = &Server{}

// 数据库 相关 start
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
	IdleTimeout int
}

var RedisSetting = &Redis{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatalf("setting Setup,fail to load config/config.ini:%v", err.Error())
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)

	fmt.Printf("map AppSetting:  %+v \n\n", AppSetting)
	fmt.Printf("map ServerSetting:  %+v \n\n", ServerSetting)
	fmt.Printf("map DatabaseSetting:  %+v \n\n", DatabaseSetting)
	fmt.Printf("map RedisSetting:  %+v \n\n", RedisSetting)
}

func mapTo(selection string, target any) {
	err := cfg.Section(selection).MapTo(target)
	if err != nil {
		log.Fatalf("cfg.MapTo selection: %s,err:%v", selection, err)
	}
}

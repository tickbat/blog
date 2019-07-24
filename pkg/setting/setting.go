package setting

import (
	"time"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	conf 		config
	App 		app
	Server 		server
	Database 	database
	Log			logs
	Image		image
)

func init() {
	// 里面的字段开头必须大写才能decode，解析的时候不区分大小写
	if _, err := toml.DecodeFile("conf/app.toml", &conf); err != nil {
		log.Fatal("decode config error: %v" + err.Error())
	}

	conf.Image.MaxSize = conf.Image.MaxSize * 1024 * 1024

	App 		= conf.App
	Server 		= conf.Server
	Database 	= conf.Database
	Log 		= conf.Log
	Image 		= conf.Image
}

type config struct {
	App     	app
	Server 		server
	Database 	database
	Log 		logs
	Image		image
}

type app struct {
	RunMode string
	PageSize  int
	JwtSecret string
}

type server struct {
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type logs struct {
	LogSavePath	string
	LogSaveName	string
	LogFileExt	string
	TimeFormat	string
}

type image struct {
	PrefixUrl	string
	SavePath	string
	AllowExts  	[]string
	MaxSize 	int
}

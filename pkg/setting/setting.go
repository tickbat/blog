package setting

import (
	"time"
	"github.com/BurntSushi/toml"
	"blog/pkg/logging"
)

var (
	conf 		config
	App 		app
	Server 		server
	Database 	database
	Log			log
)

func init() {
	if _, err := toml.DecodeFile("conf/app.toml", &conf); err != nil {
		logging.Error("decode config error: %v" + err.Error())
	}
	App 		= conf.app
	Server 		= conf.server
	Database 	= conf.database
	Log 		= conf.log
}

type config struct {
	app     	app
	server 		server
	database 	database
	log 		log
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

type log struct {
	LogSavePath	string `toml:"LOG_SAVE_PATH"`
	LogSaveName	string `toml:"LOG_SAVE_NAME"`
	LogFileExt	string `toml:"LOG_FILE_EXT"`
	TimeFormat	string `toml:"TIME_FORMAT"`
}

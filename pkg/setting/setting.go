package setting

import (
	"github.com/BurntSushi/toml"
	"log"
	"time"
)

var (
	App      app
	Server   server
	Database database
	Log      logs
	Image    image
	Redis    redis
	Excel    excel
)

func init() {
	var conf config
	// 里面的字段开头必须大写才能decode，解析的时候不区分大小写
	if _, err := toml.DecodeFile("conf/app.toml", &conf); err != nil {
		log.Fatal("decode config error: %v" + err.Error())
	}

	conf.Image.MaxSize = conf.Image.MaxSize * 1024 * 1024

	App = conf.App
	Server = conf.Server
	Database = conf.Database
	Log = conf.Log
	Image = conf.Image
	Redis = conf.Redis
	Excel = conf.Excel
}

type config struct {
	App      app
	Server   server
	Database database
	Log      logs
	Image    image
	Redis    redis
	Excel    excel
}

type app struct {
	RunMode   string
	PageSize  int
	JwtSecret string
	PrefixUrl string
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
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

type image struct {
	PrefixUrl string
	SavePath  string
	AllowExts []string
	MaxSize   int
}

type redis struct {
	Password string
	Host     string
}

type excel struct {
	SavePath string
}

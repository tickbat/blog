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
)

func init() {
	if _, err := toml.DecodeFile("conf/app.toml", &conf); err != nil {
		logging.Error("decode config error: %v" + err.Error())
	}
	App 		= conf.App
	Server 		= conf.Server
	Database 	= conf.Database
}

type config struct {
	App     	app
	Server 		server
	Database 	database
}

type app struct {
	RunMode string `toml:"RUN_MODE"`
	PageSize  int    `toml:"PAGE_SIZE"`
	JwtSecret string `toml:"JWT_SECRET"`
}

type server struct {
	HttpPort     int 			`toml:"HTTP_PORT"`
	ReadTimeout  time.Duration `toml:"READ_TIMEOUT"`
	WriteTimeout time.Duration `toml:"WRITE_TIMEOUT"`
}

type database struct {
	Type        string `toml:"TYPE"`
	User        string `toml:"USER"`
	Password    string `toml:"PASSWORD"`
	Host        string `toml:"HOST"`
	Name        string `toml:"NAME"`
	TablePrefix string `toml:"TABLE_PREFIX"`
}

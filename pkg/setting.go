package pkg

import (
	"github.com/BurntSushi/toml"
	"log"
)

var (
	Config = new(config)
)

func init() {
	if _, err := toml.DecodeFile("conf/app.toml", Config); err != nil {
		// handle error
		log.Printf("decode config error: %v", err)
	}
}

type config struct {
	RunMode string `toml:"RUN_MODE"`
	App     *app
}

type app struct {
	PageSize  int    `toml:"PAGE_SIZE"`
	JwtSecret string `toml:"JWT_SECRET"`
}

type server struct {
	HttpPort     int `toml:"HTTP_PORT"`
	ReadTimeout  int `toml:"READ_TIMEOUT"`
	WriteTimeout int `toml:"WRITE_TIMEOUT"`
}

type database struct {
	Type        string `toml:"TYPE"`
	User        string `toml:"USER"`
	Password    string `toml:"PASSWORD"`
	Host        string `toml:"HOST"`
	Name        string `toml:"NAME"`
	TablePrefix string `toml:"TABLE_PREFIX"`
}

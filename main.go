package main

import (
	"blog/pkg/setting"
	"fmt"
	"blog/routers"
	"time"
	"net/http"
	_"blog/cron"
)

func main() {
	router := routers.InitRouters()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.Server.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.Server.ReadTimeout * time.Second,
		WriteTimeout:   setting.Server.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

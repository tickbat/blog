package main

import (
	"blog/pkg/setting"
	"fmt"
	"blog/routers"
	"time"
	"net/http"
)

func main() {
	router := routers.InitRouters()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.Config.Server.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.Config.Server.ReadTimeout * time.Second,
		WriteTimeout:   setting.Config.Server.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

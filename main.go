package main

import (
	"blog/pkg/setting"
	"fmt"
	"net/http"
	"blog/routers"
)

func main() {
	router := routers.InitRouters()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.Config.Server.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.Config.Server.ReadTimeout,
		WriteTimeout:   setting.Config.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

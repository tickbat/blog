package main

import (
	_ "blog/cron"
	_ "blog/pkg/gredis"
	"blog/pkg/logging"
	"blog/pkg/setting"
	"blog/routers"
	_ "blog/validator"
	"fmt"
	"net/http"
	"time"
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
	if err := s.ListenAndServe(); err != nil {
		logging.Fatal("ListenAndServe error: ", err.Error())
	}
}

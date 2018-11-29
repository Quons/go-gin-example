package main

import (
	"fmt"
	"net/http"
	"runtime"
	"syscall"

	"github.com/fvbock/endless"

	"github.com/Quons/go-gin-example/models"
	"github.com/Quons/go-gin-example/pkg/gredis"
	"github.com/Quons/go-gin-example/pkg/logging"
	"github.com/Quons/go-gin-example/pkg/setting"
	"github.com/Quons/go-gin-example/routers"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()
// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/Quons/go-gin-example

// @license.name MIT
// @license.url https://github.com/Quons/go-gin-example/blob/master/LICENSE
func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	logrus.Debug("...Debug log")
	logrus.Info("...Info log")
	logrus.Warn("...Warn log")
	logrus.Error("...Error log")
	logrus.WithFields(logrus.Fields{
		"name": "quon",
		"age":  12,
	}).Info("with field info log")
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	if runtime.GOOS == "windows" {
		server := &http.Server{
			Addr:           endPoint,
			Handler:        routersInit,
			ReadTimeout:    readTimeout,
			WriteTimeout:   writeTimeout,
			MaxHeaderBytes: maxHeaderBytes,
		}

		server.ListenAndServe()
		return
	}

	endless.DefaultReadTimeOut = readTimeout
	endless.DefaultWriteTimeOut = writeTimeout
	endless.DefaultMaxHeaderBytes = maxHeaderBytes
	server := endless.NewServer(endPoint, routersInit)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}

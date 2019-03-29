package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Quons/go-gin-example/models"
	"github.com/Quons/go-gin-example/pkg/gredis"
	"github.com/Quons/go-gin-example/pkg/logging"
	"github.com/Quons/go-gin-example/pkg/setting"
	"github.com/Quons/go-gin-example/routers"
	"github.com/gin-contrib/cors"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
)

func init() {
	var runmode string
	flag.StringVar(&runmode, "runmode", "dev", "runmode:dev,test,pre,prod;default dev mode")
	flag.Parse()
	setting.Setup(runmode)
	logging.Setup()
	models.Setup()
	gredis.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/Quons/go-gin-example

// @license.name MIOd
// @license.url https://github.com/Quons/go-gin-example/blob/master/LICENSE
func main() {
	routersInit := routers.InitRouter()

	//跨域请求设置
	corsConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization",  "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}
	routersInit.Use(cors.New(corsConfig))
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
	//服务器设置
	srv := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ErrorLog:       log.New(logging.GetLogrusWriter(), "server err:", log.LstdFlags),
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()
	logrus.Info("server started")
	//平滑重启设置 福
	//Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logrus.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("Server Shutdown:", err)
	}
	logrus.Println("Server exiting")
}

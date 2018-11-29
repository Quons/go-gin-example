package logging

import (
	"github.com/Quons/go-gin-example/pkg/file"
	"github.com/pkg/errors"
	"time"
	"path/filepath"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/rifflock/lfshook"
	"github.com/Quons/go-gin-example/pkg/setting"
)

func Setup() {
	//获取执行目录
	logPath, err := file.MkRdir("log")
	if err != nil {
		logrus.Fatal("get log path error")
	}
	dirName, err := file.GetDirName()
	if err != nil {
		logrus.Fatal("get dirName error")
	}
	//设置日志级别
	logLevel, err := logrus.ParseLevel(setting.AppSetting.LogLevel)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	logrus.SetLevel(logLevel)
	//打印行号，funcName
	logrus.SetReportCaller(true)
	//输出设置
	writer, err := rotatelogs.New(
		filepath.Join(logPath, dirName+".%Y%m%d%H%M"),
		rotatelogs.WithLinkName(filepath.Join(logPath, dirName+".log")), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(10*time.Hour*24),                          // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Hour*24),                       // 日志切割时间间隔
	)
	if err != nil {
		logrus.Fatalf("config local file system logger error.%+v", errors.WithStack(err))
	}
	//设置local file system hook
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{
		//是否显示颜色
		ForceColors: true,
		//输出字段排序设置
		DisableSorting: true,
		//设置日志格式
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//添加hook
	logrus.AddHook(lfsHook)
}

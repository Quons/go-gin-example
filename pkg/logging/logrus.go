package logging

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Quons/go-gin-example/pkg/file"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Quons/go-gin-example/pkg/logging/logstash"
	"github.com/astaxie/beego"
	"strings"
)

var logPath string
var dirName string

func Setup() {
	//获取执行目录
	var err error
	logPath, err = file.MkRdir("logs")
	if err != nil {
		logrus.Fatal("get log path error")
	}
	dirName, err = file.GetDirName()
	if err != nil {
		logrus.Fatal("get dirName error")
	}
	//设置日志级别
	logLevel, err := logrus.ParseLevel(beego.AppConfig.String("logLevel"))
	if err != nil {
		logrus.Fatal(err.Error())
	}
	logrus.SetLevel(logLevel)
	//打印行号，funcName
	logrus.SetReportCaller(true)
	//输出设置
	writer := GetLogrusWriter()
	//设置local file system hook
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &CodeFormatter{})
	//添加hook
	logrus.AddHook(lfsHook)
	//elasticSearch 推送配置，如果推动地址不为空，则进行推送
	enableEsPush := beego.AppConfig.DefaultBool("enableEsPush", false)
	if enableEsPush {
		appName := beego.AppConfig.String("appname") + "_" + beego.AppConfig.String("runmode")
		logtashHook, err := logstash.NewHookWithFields("udp", beego.AppConfig.String("esPushUrl"), "", logrus.Fields{"appname": appName})
		if err != nil {
			logrus.Fatal(err)
		}
		logrus.AddHook(logtashHook)
	}
}

//定义formatter ,实现logrus formatter接口
type CodeFormatter struct{}

func (f *CodeFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	fileLineNum := ""
	if entry.Caller != nil {
		srcIndex := strings.Index(entry.Caller.File, "src")
		fileLineNum = string([]rune(entry.Caller.File)[srcIndex+4:])
		fileLineNum = fmt.Sprintf("%s:%v ", fileLineNum, strconv.Itoa(entry.Caller.Line))
	}

	b.WriteString(entry.Time.Format("2006-01-02 15:04:05"))
	b.WriteString(" [")
	b.WriteString(entry.Level.String())
	b.WriteString("] ")
	b.WriteString(fileLineNum)

	if len(entry.Data) != 0 {
		b.WriteString("param:")
		data, _ := json.Marshal(entry.Data)
		b.WriteString(fmt.Sprintf("%+v ", string(data)))
	}

	b.WriteString("msg:")
	b.WriteString(entry.Message)
	b.WriteByte('\n')
	return b.Bytes(), nil
}

func GetLogrusWriter() *rotatelogs.RotateLogs {
	logrusPath, err := file.MkRdir("logs/" + dirName)
	if err != nil {
		logrus.Fatal("get log path error")
	}
	writer, err := rotatelogs.New(
		filepath.Join(logrusPath, dirName+".%Y%m%d%H%M"),
		rotatelogs.WithLinkName(filepath.Join(logPath, dirName+".log")), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(10*time.Hour*24),                          // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Hour*24),                       // 日志切割时间间隔
	)
	if err != nil {
		logrus.Fatalf("config local file system logger error.%+v", errors.WithStack(err))
	}
	return writer
}
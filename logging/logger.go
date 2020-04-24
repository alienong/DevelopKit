/**
 * @Author: alienongwlx@gmail.com
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/4/24 10:07
 */

package logging

import (
	"github.com/alienong/DevelopKit/utils/misc"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Logger = logrus.New()

func init() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	Logger.Out = src
	Logger.SetLevel(logrus.DebugLevel)
	if ok, _ := misc.PathExists("./log"); !ok {
		err = os.Mkdir("log", os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	logPath := "./log/api.log"
	writer, err := rotatelogs.New(
		logPath+".%Y%m%d.log",
		// WithLinkName为最新的日志建立软连接,以方便随着找到当前日志文件
		rotatelogs.WithLinkName(logPath),
		// WithRotationTime设置日志分割的时间,这里设置为一小时分割一次
		rotatelogs.WithRotationTime(time.Hour*24),
		// WithMaxAge和WithRotationCount二者只能设置一个,
		// WithMaxAge设置文件清理前的最长保存时间,
		// WithRotationCount设置文件清理前最多保存的个数.
		//rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationCount(3),
	)
	if err != nil {
		Logger.Errorf("config local file system for logger error: %v", err)
	}
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{DisableColors: true})
	Logger.AddHook(lfsHook)
}

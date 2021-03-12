package logger_logrus

import (
	"fmt"
	formatter "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	consoleFmt := getFormatter(true)
	fileFmt := getFormatter(false)
	Logger.SetFormatter(consoleFmt)
	Logger.SetReportCaller(true)
	infoWriter := getWriter("/data/log/goChatDemo/im_info.log")
	errorWriter := getWriter("/data/log/goChatDemo/im_error.log")
	Logger.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  infoWriter,
			logrus.ErrorLevel: errorWriter,
		},
		fileFmt,
	))

}

func getFormatter(isConsole bool) *formatter.Formatter {
	fmtter := &formatter.Formatter{
		TimestampFormat: "2006-01-02 15:04:05.299",
		CallerFirst:     true,
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			funcInfo := runtime.FuncForPC(frame.PC)
			if funcInfo == nil {
				return "error during runtime.FuncForPC"
			}
			fullPath, line := funcInfo.FileLine(frame.PC)
			return fmt.Sprintf(" [%v:%v]", filepath.Base(fullPath), line)
		},
	}
	if isConsole {
		fmtter.NoColors = false
	} else {
		fmtter.NoColors = true
	}
	return fmtter
}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的logger 实际生成的文件名 demo.log.YYmmddHH
	// *.log是指向最新日志的链接
	// 保存30天内的日志，每天分割一次日志
	hook, err := rotatelogs.New(
		strings.Replace(filename, ".log", "", -1)+"_%Y%m%d.log", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

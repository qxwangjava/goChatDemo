package logger

import (
	"fmt"
	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-gonic/gin"
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

type ErrHook struct {
}

const (
	InfoFileName  = "/data/log/goChatDemo/im_info.log"
	ErrorFileName = "/data/log/goChatDemo/im_error.log"
)

// Levels 只定义 error 和 panic 等级的日志,其他日志等级不会触发 hook
func (h *ErrHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
		logrus.PanicLevel,
	}
}

// Fire 将异常栈打印出来
func (h *ErrHook) Fire(entry *logrus.Entry) error {
	var errStack = StackTrace(false)
	entry.Message = errStack
	return nil
}

func init() {
	Logger = logrus.New()
	consoleFmt := getFormatter(true)
	fileFmt := getFormatter(false)
	Logger.SetFormatter(consoleFmt)
	Logger.SetReportCaller(true)
	Logger.SetLevel(logrus.DebugLevel)
	infoWriter := getWriter(InfoFileName)
	errorWriter := getWriter(ErrorFileName)
	Logger.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  infoWriter,
			logrus.DebugLevel: infoWriter,
			logrus.ErrorLevel: errorWriter,
		},
		fileFmt,
	))

	// gin log配置
	gin.DefaultWriter = infoWriter
	gin.DefaultErrorWriter = errorWriter
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
	fmtter.NoColors = true
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

func StackTrace(all bool) string {

	// Reserve 10K buffer at first
	buf := make([]byte, 10240)
	for {
		size := runtime.Stack(buf, all)
		// The size of the buffer may be not enough to hold the stacktrace,
		// so double the buffer size
		if size == len(buf) {
			buf = make([]byte, len(buf)<<1)
			continue
		}
		break
	}
	return string(buf)
}

package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

const (
	bold      = 1
	underline = 4
	black     = 30
	red       = 31
	green     = 32
	yellow    = 33
	blue      = 34
	purple    = 35
	cyan      = 36
	white     = 37
)

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = white
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = green
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	basicConfig := GetBasicConfig()
	// 自定义日期格式
	timeStamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		// 自定义日志格式
		// 输出格式：前缀 时间 日志级别 文件:行号 函数名 日志内容
		// 其中文件:行号 函数名 可以根据需求选择是否输出
		funcName := entry.Caller.Function
		filePath := entry.Caller.File + ":" + strconv.Itoa(entry.Caller.Line)
		_, err := fmt.Fprintf(b, "%s[%s]\x1b[%d;%dm[%s]\x1b[0m %s %s \x1b[%d;%dm[%s]\x1b[0m\n",
			basicConfig.Logger.Prefix,
			timeStamp,
			bold,
			levelColor,
			strings.ToUpper(entry.Level.String()),
			filePath,
			funcName,
			bold,
			cyan,
			entry.Message,
		)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := fmt.Fprintf(b, "%s[%s]\x1b[%dm[%s]\x1b[0m[%s]\n",
			basicConfig.Logger.Prefix,
			timeStamp,
			levelColor,
			strings.ToUpper(entry.Level.String()),
			entry.Message,
		)
		if err != nil {
			return nil, err
		}
	}
	return b.Bytes(), nil
}

// initLogger 初始化日志记录器
func initLogger() *logrus.Logger {
	basicConfig := GetBasicConfig()
	mLog := logrus.New()
	mLog.SetOutput(os.Stdout)          // 设置输出类型
	mLog.SetReportCaller(true)         // 开启返回函数名和行号
	mLog.SetFormatter(&LogFormatter{}) // 设置自定义的 Formatter
	level, err := logrus.ParseLevel(basicConfig.Logger.LogLevel)
	if err != nil {
		level = logrus.DebugLevel // 设置最低的日志级别
	}
	mLog.SetLevel(level) // 设置最低的日志级别
	return mLog
}

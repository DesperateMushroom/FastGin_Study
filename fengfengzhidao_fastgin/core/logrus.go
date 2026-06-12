package core

import (
	"bytes"
	"fmt"
	"path"

	"github.com/sirupsen/logrus"
)

type MyLog struct {
}

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

// 假设你定义了这个方法，你在项目里就可以这样把 Logrus 的默认格式替换成你的 MyLog
func (MyLog) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据不同level展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() {
	logrus.SetLevel(logrus.DebugLevel) // 设置从什么level开始记录日记
	logrus.SetReportCaller(true)       //设置行号
	logrus.SetFormatter(MyLog{})       //设置格式（自定义格式
	// logrus.SetFormatter(&logrus.JSONFormatter{}) //设置json格式：一般向外发（发给其他程序）的时候会用json格式
}

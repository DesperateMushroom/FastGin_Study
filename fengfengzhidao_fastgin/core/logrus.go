package core

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"sync"

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
	logrus.AddHook(&MyHook{
		logPath: "logs",
	})
}

type MyHook struct {
	file *os.File //当前打开的日志文件
	// 用*而不用原型：主要原因是 Go 标准库的 os.Open()、os.Create() 返回的就是 *os.File，而且文件对象本身应该被多个函数共享，而不是复制。
	errFile *os.File   //错误的日志文件
	date    string     // 当前日志的时间
	logPath string     // 日志的目录
	mu      sync.Mutex // 加把锁防止并发问题
}

// 在这个钩子函数中实现：1.产生日志写到文件中去；2.时间切片；3.错误的日志单独存放；
func (hook *MyHook) Fire(entry *logrus.Entry) error {

	hook.mu.Lock()
	defer hook.mu.Unlock()

	// 1.产生日志写到文件中去；

	// 错误示范：直接写下面四行 - 如果有大量日志生成，这里会频繁打开关闭file，会非常非常耗时耗资源
	// file, _ := os.OpenFile("xxx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// msg, _ := entry.String()
	// file.Write([]byte(msg))
	// file.Close()

	// 正确操作：在MyHook struct里定义一个file
	// 如果那个file为空时，创建新的file
	msg, _ := entry.String()
	// 2.时间切片；
	date := entry.Time.Format("2006-01-02")
	if hook.date != date {
		// 换时间，file文件要创建新的了
		hook.rotateFiles(date) // 在这个函数里创建file
		hook.date = date
	}

	// 3.错误的日志单独存放；写入的时候判断
	if entry.Level <= logrus.ErrorLevel {
		hook.errFile.Write([]byte(msg))
	}
	hook.file.Write([]byte(msg))

	// fmt.Println(entry)
	return nil
}

// 轮换时间日志
func (hook *MyHook) rotateFiles(timer string) error {
	// 开始轮换前，查看是否有file还在写，还有的话要close掉
	if hook.file != nil {
		hook.file.Close() //close了才可以创建新的
	}
	if hook.errFile != nil {
		hook.errFile.Close() //close了才可以创建新的
	}
	if hook.file == nil {
		// 根据date创建目录
		logDir := fmt.Sprintf("%s/%s", hook.logPath, timer)
		os.MkdirAll(logDir, 0666)
		logPath := fmt.Sprintf("%s/info.log", logDir)

		// 文件名，打开方式，文件权限
		// os.O_CREATE: 文件不存在就创建
		// os.O_WRONLY：只写不读
		// os.O_APPEND：追加模式，写在末尾，不覆盖
		// 0666 linux文件权限: 读4，写2，执行1 - 6可读可写；位置从左到右：owner, group, other
		file, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

		// 3.错误的日志单独存放；
		errLogPath := fmt.Sprintf("%s/err.log", logDir)
		errFile, _ := os.OpenFile(errLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

		hook.file = file
		hook.errFile = errFile
	}
	return nil
}

func (*MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

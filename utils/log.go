package utils

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"runtime"
	"strings"
)

// consoleLogs开发模式下日志
var consoleLogs *logs.BeeLogger

// fileLogs 生产环境下日志
var fileLogs *logs.BeeLogger
var fileAccessLogs *logs.BeeLogger
var runmode string

func init() {
	consoleLogs = logs.NewLogger(1)
	consoleLogs.SetLogger(logs.AdapterConsole)
	fileLogs = logs.NewLogger(10000)
	fileLogs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/debox-chain.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],"level":7,"daily":true,"maxdays":5}`)

	//fileAccessLogs = logs.NewLogger(10000)
	//fileAccessLogs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/debox-access.log","level":7,"daily":true,"maxdays":5}`)

	runmode = "dev"
}

func LogOut(level, v interface{}) {
	var format string
	format = "%s:%d %s"
	doLogOut(format, level, v)
}

func doLogOut(format string, level, v interface{}) {
	if level == "" {
		level = "debug"
	}
	_, file, line, _ := runtime.Caller(2)
	if lastIndex := strings.LastIndex(file, "/"); lastIndex != -1 {
		file = file[lastIndex+1:]
	}

	if runmode == "dev" {
		switch level {
		case "emergency":
			fileLogs.Emergency(format, file, line, v)
		case "alert":
			fileLogs.Alert(format, file, line, v)
		case "critical":
			fileLogs.Critical(format, file, line, v)
		case "error":
			fileLogs.Error(format, file, line, v)
		case "warning":
			fileLogs.Warning(format, file, line, v)
		case "notice":
			fileLogs.Notice(format, file, line, v)
		case "informational":
			fileLogs.Informational(format, file, line, v)
		case "debug":
			fileLogs.Debug(format, file, line, v)
		case "warn":
			fileLogs.Warn(format, file, line, v)
		case "info":
			fileLogs.Info(format, file, line, v)
		case "trace":
			fileLogs.Trace(format, file, line, v)
		default:
			fileLogs.Debug(format, file, line, v)
		}
	}
	switch level {
	case "emergency":
		consoleLogs.Emergency(format, file, line, v)
	case "alert":
		consoleLogs.Alert(format, file, line, v)
	case "critical":
		consoleLogs.Critical(format, file, line, v)
	case "error":
		consoleLogs.Error(format, file, line, v)
	case "warning":
		consoleLogs.Warning(format, file, line, v)
	case "notice":
		consoleLogs.Notice(format, file, line, v)
	case "informational":
		fileLogs.Informational(format, file, line, v)
	case "debug":
		consoleLogs.Debug(format, file, line, v)
	case "warn":
		consoleLogs.Warn(format, file, line, v)
	case "info":
		consoleLogs.Info(format, file, line, v)
	case "trace":
		consoleLogs.Trace(format, file, line, v)
	default:
		consoleLogs.Debug(format, file, line, v)
	}
}

func LogAccessOut(level, v interface{}) {
	format := "%s:%d %s"
	if level == "" {
		level = "debug"
	}
	_, file, line, _ := runtime.Caller(1)
	if lastIndex := strings.LastIndex(file, "/"); lastIndex != -1 {
		file = file[lastIndex+1:]
	}

	if runmode == "dev" {
		switch level {
		case "emergency":
			fileAccessLogs.Emergency(format, file, line, v)
		case "alert":
			fileAccessLogs.Alert(format, file, line, v)
		case "critical":
			fileAccessLogs.Critical(format, file, line, v)
		case "error":
			fileAccessLogs.Error(format, file, line, v)
		case "warning":
			fileAccessLogs.Warning(format, file, line, v)
		case "notice":
			fileAccessLogs.Notice(format, file, line, v)
		case "informational":
			fileAccessLogs.Informational(format, file, line, v)
		case "debug":
			fileAccessLogs.Debug(format, file, line, v)
		case "warn":
			fileAccessLogs.Warn(format, file, line, v)
		case "info":
			fileAccessLogs.Info(format, file, line, v)
		case "trace":
			fileAccessLogs.Trace(format, file, line, v)
		default:
			fileAccessLogs.Debug(format, file, line, v)
		}
	}

}

// FLogOut 格式化打印日志
func FLogOut(level interface{}, format string, a ...any) {
	var f string
	f = "%s:%d %s"
	doLogOut(f, level, fmt.Sprintf(format, a...))
}

// LogOutWithCtx 通过传递gin.Context上下文对象，获取request_id
func LogOutWithCtx(ctx *gin.Context, level interface{}, format string, a ...any) {
	if requestId, found := ctx.Get("request_id"); found {
		format = "[" + requestId.(string) + "]" + "%s:%d %s"
	} else {
		format = "%s:%d %s"
	}
	doLogOut(format, level, fmt.Sprintf(format, a...))
}

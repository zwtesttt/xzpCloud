package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// Level 日志级别类型
type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

var levelNames = map[Level]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
	FatalLevel: "FATAL",
}

var levelColors = map[Level]string{
	DebugLevel: "\033[36m", // cyan
	InfoLevel:  "\033[32m", // green
	WarnLevel:  "\033[33m", // yellow
	ErrorLevel: "\033[31m", // red
	FatalLevel: "\033[35m", // magenta
}

const colorReset = "\033[0m"

// Logger 日志器结构体
type Logger struct {
	level      Level
	output     io.Writer
	colorize   bool
	showCaller bool
	prefix     string
}

// Config 日志配置
type Config struct {
	Level      Level
	Output     io.Writer
	Colorize   bool
	ShowCaller bool
	Prefix     string
}

var defaultLogger *Logger

func init() {
	// 初始化默认日志器
	defaultLogger = New(Config{
		Level:      InfoLevel,
		Output:     os.Stdout,
		Colorize:   true,
		ShowCaller: true,
		Prefix:     "",
	})
}

// New 创建一个新的日志器
func New(config Config) *Logger {
	if config.Output == nil {
		config.Output = os.Stdout
	}

	return &Logger{
		level:      config.Level,
		output:     config.Output,
		colorize:   config.Colorize,
		showCaller: config.ShowCaller,
		prefix:     config.Prefix,
	}
}

// SetLevel 设置默认日志器的日志级别
func SetLevel(level Level) {
	defaultLogger.SetLevel(level)
}

// SetLevel 设置日志级别
func (l *Logger) SetLevel(level Level) {
	l.level = level
}

// SetOutput 设置输出目标
func (l *Logger) SetOutput(output io.Writer) {
	l.output = output
}

// SetColorize 设置是否启用颜色
func (l *Logger) SetColorize(colorize bool) {
	l.colorize = colorize
}

// log 内部日志方法
func (l *Logger) log(level Level, args ...interface{}) {
	if level < l.level {
		return
	}

	// 获取调用者信息
	var caller string
	if l.showCaller {
		_, file, line, ok := runtime.Caller(2)
		if ok {
			caller = fmt.Sprintf(" %s:%d", filepath.Base(file), line)
		}
	}

	// 构建日志消息
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	levelName := levelNames[level]
	message := fmt.Sprint(args...)

	var logLine string
	if l.colorize {
		color := levelColors[level]
		logLine = fmt.Sprintf("%s [%s%s%s]%s %s%s\n",
			timestamp, color, levelName, colorReset, caller, l.prefix, message)
	} else {
		logLine = fmt.Sprintf("%s [%s]%s %s%s\n",
			timestamp, levelName, caller, l.prefix, message)
	}

	l.output.Write([]byte(logLine))

	// Fatal级别直接退出程序
	if level == FatalLevel {
		os.Exit(1)
	}
}

// logf 内部格式化日志方法
func (l *Logger) logf(level Level, format string, args ...interface{}) {
	if level < l.level {
		return
	}

	message := fmt.Sprintf(format, args...)
	l.log(level, message)
}

// Debug 调试级别日志
func (l *Logger) Debug(args ...interface{}) {
	l.log(DebugLevel, args...)
}

// Debugf 格式化调试级别日志
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logf(DebugLevel, format, args...)
}

// Info 信息级别日志
func (l *Logger) Info(args ...interface{}) {
	l.log(InfoLevel, args...)
}

// Infof 格式化信息级别日志
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logf(InfoLevel, format, args...)
}

// Warn 警告级别日志
func (l *Logger) Warn(args ...interface{}) {
	l.log(WarnLevel, args...)
}

// Warnf 格式化警告级别日志
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logf(WarnLevel, format, args...)
}

// Error 错误级别日志
func (l *Logger) Error(args ...interface{}) {
	l.log(ErrorLevel, args...)
}

// Errorf 格式化错误级别日志
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logf(ErrorLevel, format, args...)
}

// Fatal 致命错误级别日志（会退出程序）
func (l *Logger) Fatal(args ...interface{}) {
	l.log(FatalLevel, args...)
}

// Fatalf 格式化致命错误级别日志（会退出程序）
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logf(FatalLevel, format, args...)
}

// 包级别的便利函数，使用默认日志器

// Debug 调试级别日志
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Debugf 格式化调试级别日志
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

// Info 信息级别日志
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Infof 格式化信息级别日志
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// Warn 警告级别日志
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Warnf 格式化警告级别日志
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

// Error 错误级别日志
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Errorf 格式化错误级别日志
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// Fatal 致命错误级别日志（会退出程序）
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// Fatalf 格式化致命错误级别日志（会退出程序）
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

// With 创建带有前缀的新日志器
func With(prefix string) *Logger {
	return &Logger{
		level:      defaultLogger.level,
		output:     defaultLogger.output,
		colorize:   defaultLogger.colorize,
		showCaller: defaultLogger.showCaller,
		prefix:     defaultLogger.prefix + "[" + prefix + "] ",
	}
}

// GetLogger 返回默认日志器
func GetLogger() *Logger {
	return defaultLogger
}

// ParseLevel 解析日志级别字符串
func ParseLevel(levelStr string) (Level, error) {
	switch strings.ToUpper(levelStr) {
	case "DEBUG":
		return DebugLevel, nil
	case "INFO":
		return InfoLevel, nil
	case "WARN", "WARNING":
		return WarnLevel, nil
	case "ERROR":
		return ErrorLevel, nil
	case "FATAL":
		return FatalLevel, nil
	default:
		return InfoLevel, fmt.Errorf("invalid log level: %s", levelStr)
	}
}

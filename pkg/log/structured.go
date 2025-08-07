package log

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// Field 结构化日志字段
type Field struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

// StructuredLogger 结构化日志器
type StructuredLogger struct {
	level      Level
	output     io.Writer
	jsonFormat bool
}

// NewStructuredLogger 创建新的结构化日志器
func NewStructuredLogger(level Level, output io.Writer, jsonFormat bool) *StructuredLogger {
	if output == nil {
		output = os.Stdout
	}
	return &StructuredLogger{
		level:      level,
		output:     output,
		jsonFormat: jsonFormat,
	}
}

// LogEntry 日志条目结构
type LogEntry struct {
	Timestamp string                 `json:"timestamp"`
	Level     string                 `json:"level"`
	Message   string                 `json:"message"`
	Caller    string                 `json:"caller,omitempty"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
}

// WithFields 添加字段
func (sl *StructuredLogger) WithFields(fields ...Field) *StructuredLogger {
	return sl
}

// logStructured 结构化日志输出
func (sl *StructuredLogger) logStructured(level Level, message string, fields ...Field) {
	if level < sl.level {
		return
	}

	// 获取调用者信息
	var caller string
	_, file, line, ok := runtime.Caller(2)
	if ok {
		caller = fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}

	entry := LogEntry{
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Level:     levelNames[level],
		Message:   message,
		Caller:    caller,
		Fields:    make(map[string]interface{}),
	}

	// 添加字段
	for _, field := range fields {
		entry.Fields[field.Key] = field.Value
	}

	var output []byte
	var err error

	if sl.jsonFormat {
		output, err = json.Marshal(entry)
		if err != nil {
			output = []byte(fmt.Sprintf(`{"error":"failed to marshal log entry: %v"}`, err))
		}
		output = append(output, '\n')
	} else {
		// 简化的键值对格式
		fieldsStr := ""
		if len(entry.Fields) > 0 {
			fieldsStr = " "
			for k, v := range entry.Fields {
				fieldsStr += fmt.Sprintf("%s=%v ", k, v)
			}
		}
		output = []byte(fmt.Sprintf("%s [%s] %s %s%s\n",
			entry.Timestamp, entry.Level, entry.Caller, entry.Message, fieldsStr))
	}

	sl.output.Write(output)

	if level == FatalLevel {
		os.Exit(1)
	}
}

// Debug 结构化调试日志
func (sl *StructuredLogger) Debug(message string, fields ...Field) {
	sl.logStructured(DebugLevel, message, fields...)
}

// Info 结构化信息日志
func (sl *StructuredLogger) Info(message string, fields ...Field) {
	sl.logStructured(InfoLevel, message, fields...)
}

// Warn 结构化警告日志
func (sl *StructuredLogger) Warn(message string, fields ...Field) {
	sl.logStructured(WarnLevel, message, fields...)
}

// Error 结构化错误日志
func (sl *StructuredLogger) Error(message string, fields ...Field) {
	sl.logStructured(ErrorLevel, message, fields...)
}

// Fatal 结构化致命错误日志
func (sl *StructuredLogger) Fatal(message string, fields ...Field) {
	sl.logStructured(FatalLevel, message, fields...)
}

// 便利函数用于创建字段
func String(key, value string) Field {
	return Field{Key: key, Value: value}
}

func Int(key string, value int) Field {
	return Field{Key: key, Value: value}
}

func Int64(key string, value int64) Field {
	return Field{Key: key, Value: value}
}

func Float64(key string, value float64) Field {
	return Field{Key: key, Value: value}
}

func Bool(key string, value bool) Field {
	return Field{Key: key, Value: value}
}

func Any(key string, value interface{}) Field {
	return Field{Key: key, Value: value}
}

func Duration(key string, value time.Duration) Field {
	return Field{Key: key, Value: value.String()}
}

func Time(key string, value time.Time) Field {
	return Field{Key: key, Value: value.Format("2006-01-02 15:04:05")}
}

func ErrorField(err error) Field {
	if err == nil {
		return Field{Key: "error", Value: nil}
	}
	return Field{Key: "error", Value: err.Error()}
}

// 全局结构化日志器实例
var defaultStructuredLogger = NewStructuredLogger(InfoLevel, os.Stdout, false)

// 包级别的结构化日志函数
func StructuredDebug(message string, fields ...Field) {
	defaultStructuredLogger.Debug(message, fields...)
}

func StructuredInfo(message string, fields ...Field) {
	defaultStructuredLogger.Info(message, fields...)
}

func StructuredWarn(message string, fields ...Field) {
	defaultStructuredLogger.Warn(message, fields...)
}

func StructuredError(message string, fields ...Field) {
	defaultStructuredLogger.Error(message, fields...)
}

func StructuredFatal(message string, fields ...Field) {
	defaultStructuredLogger.Fatal(message, fields...)
}

// SetStructuredFormat 设置结构化日志格式
func SetStructuredFormat(jsonFormat bool) {
	defaultStructuredLogger.jsonFormat = jsonFormat
}

// SetStructuredLevel 设置结构化日志级别
func SetStructuredLevel(level Level) {
	defaultStructuredLogger.level = level
}

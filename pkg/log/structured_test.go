package log

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
	"testing"
	"time"
)

// TestNewStructuredLogger 测试创建结构化日志器
func TestNewStructuredLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := NewStructuredLogger(InfoLevel, &buf, true)

	if logger == nil {
		t.Error("NewStructuredLogger返回nil")
	}

	if logger.level != InfoLevel {
		t.Errorf("期望日志级别为InfoLevel，实际为%v", logger.level)
	}

	if logger.output != &buf {
		t.Error("输出目标设置错误")
	}

	if !logger.jsonFormat {
		t.Error("JSON格式设置错误")
	}
}

// TestNewStructuredLoggerWithNilOutput 测试使用nil输出创建日志器
func TestNewStructuredLoggerWithNilOutput(t *testing.T) {
	logger := NewStructuredLogger(InfoLevel, nil, false)

	// 应该使用默认的os.Stdout
	if logger.output == nil {
		t.Error("输出不应该为nil")
	}
}

// TestStructuredLoggerInfo 测试Info级别日志
func TestStructuredLoggerInfo(t *testing.T) {
	var buf bytes.Buffer
	logger := NewStructuredLogger(InfoLevel, &buf, false)

	logger.Info("测试消息", String("key1", "value1"), Int("key2", 42))

	output := buf.String()
	if !strings.Contains(output, "[INFO]") {
		t.Error("输出中应包含[INFO]标记")
	}
	if !strings.Contains(output, "测试消息") {
		t.Error("输出中应包含消息内容")
	}
	if !strings.Contains(output, "key1=value1") {
		t.Error("输出中应包含字段信息")
	}
	if !strings.Contains(output, "key2=42") {
		t.Error("输出中应包含数字字段")
	}
}

// TestStructuredLoggerJSONFormat 测试JSON格式输出
func TestStructuredLoggerJSONFormat(t *testing.T) {
	var buf bytes.Buffer
	logger := NewStructuredLogger(InfoLevel, &buf, true)

	logger.Info("JSON测试消息", String("service", "test"), Int("port", 8080))

	output := strings.TrimSpace(buf.String())

	// 验证输出是有效的JSON
	var entry LogEntry
	err := json.Unmarshal([]byte(output), &entry)
	if err != nil {
		t.Errorf("JSON解析失败: %v", err)
	}

	// 验证JSON字段
	if entry.Level != "INFO" {
		t.Errorf("期望级别为INFO，实际为%s", entry.Level)
	}
	if entry.Message != "JSON测试消息" {
		t.Errorf("期望消息为'JSON测试消息'，实际为%s", entry.Message)
	}
	if entry.Fields["service"] != "test" {
		t.Errorf("期望service字段为'test'，实际为%v", entry.Fields["service"])
	}
	if entry.Fields["port"] != float64(8080) { // JSON数字解析为float64
		t.Errorf("期望port字段为8080，实际为%v", entry.Fields["port"])
	}
}

// TestLogLevels 测试所有日志级别
func TestLogLevels(t *testing.T) {
	testCases := []struct {
		name     string
		logFunc  func(*StructuredLogger, string, ...Field)
		expected string
	}{
		{"Debug", (*StructuredLogger).Debug, "DEBUG"},
		{"Info", (*StructuredLogger).Info, "INFO"},
		{"Warn", (*StructuredLogger).Warn, "WARN"},
		{"Error", (*StructuredLogger).Error, "ERROR"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			logger := NewStructuredLogger(DebugLevel, &buf, false)

			tc.logFunc(logger, "测试消息")

			output := buf.String()
			if !strings.Contains(output, "["+tc.expected+"]") {
				t.Errorf("期望输出包含[%s]，实际输出: %s", tc.expected, output)
			}
		})
	}
}

// TestLogLevelFiltering 测试日志级别过滤
func TestLogLevelFiltering(t *testing.T) {
	var buf bytes.Buffer
	logger := NewStructuredLogger(WarnLevel, &buf, false)

	// 这些日志不应该输出
	logger.Debug("调试消息")
	logger.Info("信息消息")

	// 这些日志应该输出
	logger.Warn("警告消息")
	logger.Error("错误消息")

	output := buf.String()

	if strings.Contains(output, "调试消息") {
		t.Error("DEBUG级别日志不应该输出")
	}
	if strings.Contains(output, "信息消息") {
		t.Error("INFO级别日志不应该输出")
	}
	if !strings.Contains(output, "警告消息") {
		t.Error("WARN级别日志应该输出")
	}
	if !strings.Contains(output, "错误消息") {
		t.Error("ERROR级别日志应该输出")
	}
}

// TestFieldFunctions 测试字段创建函数
func TestFieldFunctions(t *testing.T) {
	testTime := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	testDuration := 5 * time.Second
	testError := errors.New("测试错误")

	// 测试可比较的类型
	testCases := []struct {
		name     string
		field    Field
		expected interface{}
	}{
		{"String", String("name", "test"), "test"},
		{"Int", Int("count", 42), 42},
		{"Int64", Int64("timestamp", 1234567890), int64(1234567890)},
		{"Float64", Float64("price", 19.99), 19.99},
		{"Bool", Bool("active", true), true},
		{"Duration", Duration("elapsed", testDuration), testDuration.String()},
		{"Time", Time("created", testTime), testTime.Format(time.RFC3339)},
		{"ErrorField", ErrorField(testError), testError.Error()},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.field.Value != tc.expected {
				t.Errorf("期望值为%v，实际值为%v", tc.expected, tc.field.Value)
			}
		})
	}

	// 单独测试Any字段（不可比较的类型）
	t.Run("Any", func(t *testing.T) {
		customMap := map[string]int{"a": 1}
		field := Any("custom", customMap)

		if field.Key != "custom" {
			t.Errorf("期望键为'custom'，实际为%s", field.Key)
		}

		// 验证值的类型
		if actualMap, ok := field.Value.(map[string]int); ok {
			if actualMap["a"] != 1 {
				t.Errorf("期望map中'a'的值为1，实际为%d", actualMap["a"])
			}
		} else {
			t.Error("期望值为map[string]int类型")
		}
	})
}

// TestErrorFieldWithNil 测试错误字段处理nil情况
func TestErrorFieldWithNil(t *testing.T) {
	field := ErrorField(nil)
	if field.Key != "error" {
		t.Errorf("期望键为'error'，实际为%s", field.Key)
	}
	if field.Value != nil {
		t.Errorf("期望值为nil，实际为%v", field.Value)
	}
}

// TestPackageLevelFunctions 测试包级别的函数
func TestPackageLevelFunctions(t *testing.T) {
	var buf bytes.Buffer

	// 保存原始的默认日志器
	originalLogger := defaultStructuredLogger
	defer func() {
		defaultStructuredLogger = originalLogger
	}()

	// 设置测试用的日志器
	defaultStructuredLogger = NewStructuredLogger(DebugLevel, &buf, false)

	StructuredDebug("调试消息")
	StructuredInfo("信息消息")
	StructuredWarn("警告消息")
	StructuredError("错误消息")

	output := buf.String()

	levels := []string{"DEBUG", "INFO", "WARN", "ERROR"}
	messages := []string{"调试消息", "信息消息", "警告消息", "错误消息"}

	for i, level := range levels {
		if !strings.Contains(output, "["+level+"]") {
			t.Errorf("输出中应包含[%s]", level)
		}
		if !strings.Contains(output, messages[i]) {
			t.Errorf("输出中应包含'%s'", messages[i])
		}
	}
}

// TestSetStructuredFormat 测试设置结构化格式
func TestSetStructuredFormat(t *testing.T) {
	var buf bytes.Buffer

	// 保存原始的默认日志器
	originalLogger := defaultStructuredLogger
	defer func() {
		defaultStructuredLogger = originalLogger
	}()

	// 设置测试用的日志器
	defaultStructuredLogger = NewStructuredLogger(InfoLevel, &buf, false)

	// 切换到JSON格式
	SetStructuredFormat(true)
	StructuredInfo("JSON格式测试")

	output := strings.TrimSpace(buf.String())

	// 验证是JSON格式
	var entry LogEntry
	err := json.Unmarshal([]byte(output), &entry)
	if err != nil {
		t.Errorf("JSON解析失败: %v", err)
	}

	if entry.Message != "JSON格式测试" {
		t.Errorf("期望消息为'JSON格式测试'，实际为%s", entry.Message)
	}
}

// TestSetStructuredLevel 测试设置结构化日志级别
func TestSetStructuredLevel(t *testing.T) {
	var buf bytes.Buffer

	// 保存原始的默认日志器
	originalLogger := defaultStructuredLogger
	defer func() {
		defaultStructuredLogger = originalLogger
	}()

	// 设置测试用的日志器
	defaultStructuredLogger = NewStructuredLogger(DebugLevel, &buf, false)

	// 设置为ERROR级别
	SetStructuredLevel(ErrorLevel)

	StructuredInfo("不应该输出的信息")
	StructuredError("应该输出的错误")

	output := buf.String()

	if strings.Contains(output, "不应该输出的信息") {
		t.Error("INFO级别日志不应该输出")
	}
	if !strings.Contains(output, "应该输出的错误") {
		t.Error("ERROR级别日志应该输出")
	}
}

// TestWithFields 测试WithFields方法
func TestWithFields(t *testing.T) {
	var buf bytes.Buffer
	logger := NewStructuredLogger(InfoLevel, &buf, false)

	// WithFields方法目前返回相同的logger，这是一个简单的实现
	newLogger := logger.WithFields(String("service", "test"))

	if newLogger != logger {
		t.Log("WithFields返回了不同的logger实例") // 这不是错误，只是记录
	}
}

// TestCallerInformation 测试调用者信息
func TestCallerInformation(t *testing.T) {
	var buf bytes.Buffer
	logger := NewStructuredLogger(InfoLevel, &buf, true)

	logger.Info("测试调用者信息")

	output := strings.TrimSpace(buf.String())

	var entry LogEntry
	err := json.Unmarshal([]byte(output), &entry)
	if err != nil {
		t.Errorf("JSON解析失败: %v", err)
	}

	if entry.Caller == "" {
		t.Error("调用者信息不应该为空")
	}

	// 调用者信息应该包含文件名和行号
	if !strings.Contains(entry.Caller, ".go:") {
		t.Error("调用者信息应该包含文件名和行号")
	}
}

// TestLogEntryStructure 测试日志条目结构
func TestLogEntryStructure(t *testing.T) {
	var buf bytes.Buffer
	logger := NewStructuredLogger(InfoLevel, &buf, true)

	logger.Info("结构测试",
		String("service", "test-service"),
		Int("port", 8080),
		Bool("debug", true))

	output := strings.TrimSpace(buf.String())

	var entry LogEntry
	err := json.Unmarshal([]byte(output), &entry)
	if err != nil {
		t.Errorf("JSON解析失败: %v", err)
	}

	// 验证基本字段
	if entry.Level != "INFO" {
		t.Errorf("期望级别为INFO，实际为%s", entry.Level)
	}
	if entry.Message != "结构测试" {
		t.Errorf("期望消息为'结构测试'，实际为%s", entry.Message)
	}
	if entry.Timestamp == "" {
		t.Error("时间戳不应该为空")
	}
	if entry.Caller == "" {
		t.Error("调用者信息不应该为空")
	}

	// 验证字段
	if len(entry.Fields) != 3 {
		t.Errorf("期望3个字段，实际有%d个", len(entry.Fields))
	}

	if entry.Fields["service"] != "test-service" {
		t.Errorf("期望service为'test-service'，实际为%v", entry.Fields["service"])
	}
	if entry.Fields["port"] != float64(8080) {
		t.Errorf("期望port为8080，实际为%v", entry.Fields["port"])
	}
	if entry.Fields["debug"] != true {
		t.Errorf("期望debug为true，实际为%v", entry.Fields["debug"])
	}
}

// BenchmarkStructuredLogging 基准测试
func BenchmarkStructuredLogging(b *testing.B) {
	var buf bytes.Buffer
	logger := NewStructuredLogger(InfoLevel, &buf, false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("基准测试消息",
			String("iteration", "test"),
			Int("number", i),
			Bool("active", true))
		buf.Reset()
	}
}

// BenchmarkStructuredLoggingJSON JSON格式基准测试
func BenchmarkStructuredLoggingJSON(b *testing.B) {
	var buf bytes.Buffer
	logger := NewStructuredLogger(InfoLevel, &buf, true)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("JSON基准测试消息",
			String("iteration", "test"),
			Int("number", i),
			Bool("active", true))
		buf.Reset()
	}
}

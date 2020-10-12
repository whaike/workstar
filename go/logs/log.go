package adapter

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ILogger interface {
	Info(msg string, field ...zap.Field)
	Error(msg string, field ...zap.Field)
}

type stdLogger struct {
	Logger *zap.Logger
}

func (s stdLogger) Info(msg string, field ...zap.Field) {
	s.Logger.Info(msg, field...)
}

func (s stdLogger) Error(msg string, field ...zap.Field) {
	s.Logger.Error(msg, field...)
}

type gatewayLogger struct {
	Logger *zap.Logger
}

func (g gatewayLogger) Info(msg string, field ...zap.Field) {
	g.Logger.Info(msg, field...)
}

func (g gatewayLogger) Error(msg string, field ...zap.Field) {
	g.Logger.Error(msg, field...)
}

type sqlLogger struct {
	Logger *zap.Logger
}

func (s sqlLogger) Info(msg string, field ...zap.Field) {
	s.Logger.Info(msg, field...)
}

func (s sqlLogger) Error(msg string, field ...zap.Field) {
	s.Logger.Error(msg, field...)
}

func NewStdLogger(path string, maxSize, maxBackups, maxAge int, compress bool) ILogger {
	logger := NewLogger(path, zapcore.InfoLevel, maxSize, maxBackups, maxAge, compress)
	return &stdLogger{Logger: logger}
}

func NewSqlLogger(path string, maxSize, maxBackups, maxAge int, compress bool) ILogger {
	logger := NewLogger(path, zapcore.InfoLevel, maxSize, maxBackups, maxAge, compress)
	return &sqlLogger{Logger: logger}
}

func NewGatewayLogger(path string, maxSize, maxBackups, maxAge int, compress bool) ILogger {
	logger := NewLogger(path, zapcore.InfoLevel, maxSize, maxBackups, maxAge, compress)
	return &gatewayLogger{Logger: logger}
}

/*
 * NewLogger        获取日志
 * filePath         日志文件路径
 * level            日志级别
 * maxSize          每个日志文件保存的最大尺寸 单位：M
 * maxBackups       日志文件最多保存多少个备份
 * maxAge           文件最多保存多少天
 * compress         是否压缩
 * serviceName      服务名
 */
// NewLogger 创建新的日志管理器
func NewLogger(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool) *zap.Logger {
	core := newCore(filePath, level, maxSize, maxBackups, maxAge, compress)
	return zap.New(core, zap.AddCaller(), zap.Development())
}

// newCore zapCore 构造
func newCore(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool) zapcore.Core {
	hook := &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
		Compress:   compress,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	// 编码器
	encoderConfig := &zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	return zapcore.NewCore(
		zapcore.NewJSONEncoder(*encoderConfig),                                         // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)
}

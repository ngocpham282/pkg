package centerlog

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type logEntry struct {
	Time   time.Time `json:"time"`
	Level  string    `json:"level"`
	Msg    string    `json:"msg"`
	Fields []any     `json:"fields,omitempty"`
}

type Logger interface {
	Info(msg string, fields ...any)
	Debug(msg string, fields ...any)
	Warn(msg string, fields ...any)
	Error(msg string, fields ...any)
}

type ginLogger struct {
	ctx *gin.Context
}

const loggerKey = "centerlog.entries"

func Gin(ctx *gin.Context) Logger {
	// Set up done callback if not already set
	if _, exists := ctx.Get(loggerKey); !exists {
		ctx.Set(loggerKey, make([]logEntry, 0))
		ctx.Set("centerlog.setup", true)
		
		// Register callback to print logs when context is done
		ctx.Next()
		entries, _ := ctx.Get(loggerKey)
		if logs, ok := entries.([]logEntry); ok {
			for _, entry := range logs {
				logJSON, _ := json.Marshal(entry)
				fmt.Println(string(logJSON))
			}
		}
	}
	
	return &ginLogger{
		ctx: ctx,
	}
}

func appendTo(ctx *gin.Context, level string, msg string, fields ...any) {
	entry := logEntry{
		Time:   time.Now(),
		Level:  level,
		Msg:    msg,
		Fields: fields,
	}
	
	if entries, exists := ctx.Get(loggerKey); exists {
		if logs, ok := entries.([]logEntry); ok {
			logs = append(logs, entry)
			ctx.Set(loggerKey, logs)
		}
	}
}

func (l *ginLogger) Info(msg string, fields ...any) {
	appendTo(l.ctx, "INFO", msg, fields...)
}

func (l *ginLogger) Debug(msg string, fields ...any) {
	appendTo(l.ctx, "DEBUG", msg, fields...)
}

func (l *ginLogger) Warn(msg string, fields ...any) {
	appendTo(l.ctx, "WARN", msg, fields...)
}

func (l *ginLogger) Error(msg string, fields ...any) {
	appendTo(l.ctx, "ERROR", msg, fields...)
}

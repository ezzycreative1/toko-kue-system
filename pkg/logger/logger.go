package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"reflect"
	"runtime"
	"time"

	"github.com/ezzycreative1/toko-kue-system/pkg/constants"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var logger *zap.Logger

// InitializeLogger initializes the Zap logger
func InitializeLogger() {
	cfg := zap.NewDevelopmentConfig()
	logger, _ = cfg.Build()
	defer logger.Sync()
}

// LogInfoRequest logs information about incoming requests
func LogInfoRequest(ctx context.Context, logType, message string, request LoggerRequestData) {
	tcID := constants.PrefixForTCID + uuid.New().String()

	ts := ctx.Value(constants.ContextKeyTimeExecutionStart).(time.Time)
	te := GetValueExecutionEndFromContext(ctx, ts)

	body := new(bytes.Buffer)
	_ = json.Compact(body, request.Body)

	logger.Info(message,
		zap.String("transaction_id", ctx.Value(constants.ContextKeyTRID).(string)),
		zap.String("app_name", os.Getenv("APP_NAME")),
		zap.String("app_version", os.Getenv("APP_VERSION")),
		zap.Any("request", request),
		zap.String("response", ""),
		zap.String("log_type", logType),
		zap.String("endpoint", request.Path),
		zap.String("method_type", request.Method),
		zap.String("content_type", request.Header["Content-Type"]),
		zap.String("function_name", ""),
		zap.String("user_info", ""),
		zap.String("execution_time", te),
		zap.String("server_ip", request.Network.ServerIP),
		zap.String("trace_id", tcID),
		zap.String("body", body.String()),
		zap.String("log_flagging", "START"),
	)
}

// LogInfoResponse logs information about outgoing responses
func LogInfoResponse(ctx context.Context, logType, message string, response interface{}) {
	var request LoggerRequestData

	tcID := constants.PrefixForTCID + uuid.New().String()
	ts := ctx.Value(constants.ContextKeyTimeExecutionStart).(time.Time)
	te := GetValueExecutionEndFromContext(ctx, ts)

	if ctx.Value(constants.ContextKeyRequest) != nil {
		request = ctx.Value(constants.ContextKeyRequest).(LoggerRequestData)
	} else {
		request = LoggerRequestData{}
	}

	body := new(bytes.Buffer)
	_ = json.Compact(body, request.Body)

	logger.Info(message,
		zap.String("transaction_id", ctx.Value(constants.ContextKeyTRID).(string)),
		zap.String("app_name", os.Getenv("APP_NAME")),
		zap.String("app_version", os.Getenv("APP_VERSION")),
		zap.Any("request", request),
		zap.Any("response", response),
		zap.String("log_type", logType),
		zap.String("endpoint", request.Path),
		zap.String("method_type", request.Method),
		zap.String("content_type", request.Header["Content-Type"]),
		zap.String("function_name", ""),
		zap.String("user_info", ""),
		zap.String("execution_time", te),
		zap.String("server_ip", request.Network.ServerIP),
		zap.String("trace_id", tcID),
		zap.String("body", body.String()),
		zap.String("log_flagging", "STOP"),
	)
}

// LogError logs errors
func LogError(ctx context.Context, logType, errType, message string) {
	var request LoggerRequestData

	tcID := constants.PrefixForTCID + uuid.New().String()

	if ctx.Value(constants.ContextKeyRequest) != nil {
		request = ctx.Value(constants.ContextKeyRequest).(LoggerRequestData)
	} else {
		request = LoggerRequestData{}
	}

	body := new(bytes.Buffer)
	_ = json.Compact(body, request.Body)

	logger.Error(message,
		zap.String("transaction_id", ctx.Value(constants.ContextKeyTRID).(string)),
		zap.String("app_name", os.Getenv("APP_NAME")),
		zap.String("app_version", os.Getenv("APP_VERSION")),
		zap.String("error_type", errType),
		zap.String("log_type", logType),
		zap.String("endpoint", request.Path),
		zap.String("method_type", request.Method),
		zap.String("content_type", request.Header["Content-Type"]),
		zap.String("function_name", ""),
		zap.String("user_info", ""),
		zap.String("execution_time", ""),
		zap.String("server_ip", request.Network.ServerIP),
		zap.String("trace_id", tcID),
		zap.String("body", body.String()),
		zap.String("log_flagging", "TOP"),
	)
}

// GetFunctionName returns the name of a function
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// GetValueExecutionEndFromContext calculates the execution time
func GetValueExecutionEndFromContext(ctx context.Context, startTime time.Time) (timeSince string) {
	timeStart := ctx.Value(constants.ContextKeyTimeExecutionStart).(time.Time)
	timeSince = time.Since(timeStart).String()

	return
}

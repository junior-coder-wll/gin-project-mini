//@Author: wulinlin
//@Description:
//@File:  logger
//@Version: 1.0.0
//@Date: 2023/03/10 03:38

package middlewares

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求开始时间
		startTime := time.Now()
		// 获取请求方法和路径
		method := c.Request.Method
		path := c.Request.URL.Path
		// 处理请求
		c.Next()

		// 获取请求结束时间
		endTime := time.Now()

		// 计算请求处理时间
		latency := endTime.Sub(startTime)

		// 获取响应状态码
		statusCode := c.Writer.Status()

		// 构建日志记录器
		// 预设相应公用字段
		logFields := []zapcore.Field{
			zap.String("ip", c.ClientIP()),
			zap.Int("status_code", statusCode),
			zap.String("method", method),
			zap.String("path", path),
			zap.Duration("latency", latency),
		}

		// 判断响应状态码，并添加相应的日志级别
		switch {
		case statusCode >= 500:
			logger.Error("Server Error", logFields...)
		case statusCode >= 400:
			logger.Warn("Client Error", logFields...)
		default:
			logger.Info("Request processed", logFields...)
		}
	}
}

//
// NewLogger
//  @Description: 初始化zap的logger对象
//  @param logPath
//  @return *zap.Logger
//
func NewLogger(logPath string, maxSize int, maxAge int) *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// 定义日志切割的参数 当然下面的配置也可以从配置文件中读取，这里为了方便就不写了
	h := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    maxSize, // 单位MB
		MaxBackups: 30,
		MaxAge:     maxAge, // 单位天
		Compress:   false,
	}

	// 定义两个输出流，一个输出到控制台，一个输出到文件
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
			zapcore.InfoLevel,
		),
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(h)),
			zapcore.InfoLevel,
		),
	)

	lg := zap.New(core)
	zap.ReplaceGlobals(lg) // 替换zap包中全局的logger实例，之后在项目中使用 zap.L().xxx即可使用日志
	return lg
}

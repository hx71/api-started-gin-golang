package helpers

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"time"

	gin "github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Log to file
func LoggerToFile() gin.HandlerFunc {

	currentTime := time.Now()
	crnTime := currentTime.Format("01-02-2006")
	// log file
	fileLog := "log-file-" + crnTime + ".log"
	fileName := path.Join("logging", fileLog)

	// write file
	//src, err := os.OpenFile("lintasarta.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	src, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("err", err)
	}
	// instantiation
	logger := logrus.New()
	// Set output
	logger.Out = src
	// Set log level
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	// Set rotatelogs
	// logWriter, err := rotatelogs.New(
	// 	// Split file name
	// 	fileName+".%Y%m%d.log",
	// 	// Generate soft chain, point to the latest log file
	// 	rotatelogs.WithLinkName(fileName),
	// 	// Set maximum save time (7 days)
	// 	rotatelogs.WithMaxAge(7*24*time.Hour),
	// 	// Set log cutting interval (1 day)
	// 	rotatelogs.WithRotationTime(24*time.Hour),
	// )
	// writeMap := lfshook.WriterMap{
	// 	logrus.InfoLevel:  logWriter,
	// 	logrus.FatalLevel: logWriter,
	// 	logrus.DebugLevel: logWriter,
	// 	logrus.WarnLevel:  logWriter,
	// 	logrus.ErrorLevel: logWriter,
	// 	logrus.PanicLevel: logWriter,
	// }
	// lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// })
	// // Add Hook
	// logger.AddHook(lfHook)
	return func(c *gin.Context) {
		// start time
		startTime := time.Now()
		// Processing request
		c.Next()
		// Stop time
		endTime := time.Now()
		// execution time
		latencyTime := endTime.Sub(startTime)
		// Request mode
		reqMethod := c.Request.Method
		// Request routing
		reqUri := c.Request.RequestURI
		// Status code
		statusCode := c.Writer.Status()
		// Request IP
		clientIP := c.ClientIP()
		// Log format
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"ip":           clientIP,
			"method":       reqMethod,
			"uri":          reqUri,
		}).Info()
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

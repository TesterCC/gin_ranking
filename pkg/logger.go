package logger

import "github.com/sirupsen/logrus"

// go get github.com/sirupsen/logrus

func init()  {
	// 设置日志格式为json格式
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetReportCaller(false)

}
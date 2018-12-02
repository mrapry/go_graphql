package helper

import (
	log "github.com/sirupsen/logrus"
)

const (
	// TOPIC for setting topic of log
	TOPIC = "apigateway-log"
	// LogTag default log tag
	LogTag = "apigateway"
)

// LogContext function for logging the context of echo
// c string context
// s string scope
func LogContext(c string, s string) *log.Entry {
	return log.WithFields(log.Fields{
		"topic":   TOPIC,
		"context": c,
		"scope":   s,
	})
}

// Log function for returning entry type
// level log.Level
// message string message of log
// context string context of log
// scope string scope of log

func Log(level log.Level, message string, context string, scope string) {
	//fmt.Println(level)
	//log.SetFormatter(&log.JSONFormatter{})
	//syslogOutput, err := logrusSyslog.NewSyslogHook("", "", syslog.LOG_INFO, LogTag)
	//log.AddHook(syslogOutput)
	//
	//if err != nil {
	//	log.Fatal("Unable to setup syslog output")
	//}

	entry := LogContext(context, scope)
	switch level {
	case log.DebugLevel:
		entry.Debug(message)
	case log.InfoLevel:
		entry.Info(message)
	case log.WarnLevel:
		entry.Warn(message)
	case log.ErrorLevel:
		entry.Error(message)
	case log.FatalLevel:
		entry.Fatal(message)
	case log.PanicLevel:
		entry.Panic(message)
	}
}

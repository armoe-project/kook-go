package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

type LoggerFormatter struct {
	logrus.TextFormatter
}

func (f *LoggerFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	time := entry.Time.Format("2006-01-02 15:04:05")
	time = "\033[34m" + time + "\033[0m"
	level := strings.ToUpper(entry.Level.String())
	level = colorize(level)
	message := fmt.Sprintf("[%s] [%s]: %s\n", time, level, entry.Message)
	return []byte(message), nil
}

func colorize(level string) string {
	switch level {
	case "PANIC":
		return "\033[35m" + level + "\033[0m"
	case "FATAL":
		return "\033[31m" + level + "\033[0m"
	case "ERROR":
		return "\033[31m" + level + "\033[0m"
	case "WARNING":
		return "\033[33mWARN\033[0m"
	case "INFO":
		return "\033[32m" + level + "\033[0m"
	case "DEBUG":
		return "\033[36m" + level + "\033[0m"
	case "TRACE":
		return "\033[37m" + level + "\033[0m"
	}
	return level
}

func InitLogger() {
	setLevel()
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&LoggerFormatter{})
}

func setLevel() {
	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = "INFO"
	}
	if level == "WARN" {
		level = "WARNING"
	}

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.Fatalf("Invalid log level: %s", level)
		os.Exit(1)
	}
	logrus.SetLevel(lvl)
}

func main() {
	InitLogger()
	logrus.Info("Hello, World!")
}

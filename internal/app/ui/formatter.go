package ui

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

func Title(format string, args ...interface{}) {
	color.Blue(format+"\n", args...)
}

func Phase(format string, args ...interface{}) {
	fmt.Printf("- "+format+"\n", args...)
}

func Step(format string, args ...interface{}) {
	color.Green("  "+format+"\n", args...)
}

func Trace(format string, args ...interface{}) {
	logrus.Tracef(format, args...)
}

func Debug(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Info(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Warn(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func Error(format string, args ...interface{}) {
	color.Red(format, args...)
	// logrus.Errorf(format, args...)
}

func Panic(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

func Fatal(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}
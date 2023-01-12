// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-12
// Based on adapter by liasica, magicrolan@qq.com.

package loki

var (
    std = New()
)

func StandardLogger() *Logger {
    return std
}

func SetUrl(url string) {
    std.url = url
}

func SetJob(job string) {
    std.job = []byte(job)
}

func SetReportCaller(report bool) {
    std.reportCaller = report
}

func SetFormatter(f Formatter) {
    std.Formatter = f
}

func Wait() {
    std.WaitGroup.Wait()
}

func Trace(args ...any) {
    std.Log(std.job, TraceLevel, args...)
}

func Debug(args ...any) {
    std.Log(std.job, DebugLevel, args...)
}

func Info(args ...any) {
    std.Log(std.job, InfoLevel, args...)
}

func Warn(args ...any) {
    std.Log(std.job, WarnLevel, args...)
}

func Error(args ...any) {
    std.Log(std.job, ErrorLevel, args...)
}

func Fatal(args ...any) {
    std.Log(std.job, FatalLevel, args...)
}

func Tracef(format string, args ...any) {
    std.Logf(std.job, TraceLevel, format, args...)
}

func Debugf(format string, args ...any) {
    std.Logf(std.job, DebugLevel, format, args...)
}

func Warnf(format string, args ...any) {
    std.Logf(std.job, WarnLevel, format, args...)
}

func Infof(format string, args ...any) {
    std.Logf(std.job, InfoLevel, format, args...)
}

func Errorf(format string, args ...any) {
    std.Logf(std.job, ErrorLevel, format, args...)
}

func Fatalf(format string, args ...any) {
    std.Logf(std.job, FatalLevel, format, args...)
}

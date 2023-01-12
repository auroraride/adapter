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

func SetCallerSplitter(splitter func(s string) string) {
    std.CallerSplitter = splitter
}

func SetFormatter(f Formatter) {
    std.Formatter = f
}

func Wait() {
    std.WaitGroup.Wait()
}

func Trace(args ...any) {
    std.Trace(args...)
}

func Debug(args ...any) {
    std.Debug(args...)
}

func Info(args ...any) {
    std.Info(args...)
}

func Warn(args ...any) {
    std.Warn(args...)
}

func Error(args ...any) {
    std.Error(args...)
}

func Fatal(args ...any) {
    std.Fatal(args...)
}

func Tracef(format string, args ...any) {
    std.Tracef(format, args...)
}

func Debugf(format string, args ...any) {
    std.Debugf(format, args...)
}

func Warnf(format string, args ...any) {
    std.Warnf(format, args...)
}

func Infof(format string, args ...any) {
    std.Infof(format, args...)
}

func Errorf(format string, args ...any) {
    std.Errorf(format, args...)
}

func Fatalf(format string, args ...any) {
    std.Fatalf(format, args...)
}

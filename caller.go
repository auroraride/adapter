// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-21
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "runtime"
    "strings"
    "sync"
)

var (
    callerPackage      string
    minimumCallerDepth int
    callerInitOnce     sync.Once
)

const (
    maximumCallerDepth int = 25
)

func init() {
    // start at the bottom of the stack before the package-name cache is primed
    minimumCallerDepth = 1
}

func getPackageName(f string) string {
    for {
        lastPeriod := strings.LastIndex(f, ".")
        lastSlash := strings.LastIndex(f, "/")
        if lastPeriod > lastSlash {
            f = f[:lastPeriod]
        } else {
            break
        }
    }

    return f
}

func GetCaller(skip int) *runtime.Frame {
    // cache this package's fully-qualified name
    callerInitOnce.Do(func() {
        pcs := make([]uintptr, maximumCallerDepth)
        _ = runtime.Callers(0, pcs)

        // dynamic get the package name and the minimum caller depth
        for i := 0; i < maximumCallerDepth; i++ {
            funcName := runtime.FuncForPC(pcs[i]).Name()
            if strings.Contains(funcName, "GetCaller") {
                callerPackage = getPackageName(funcName)
                minimumCallerDepth = i
                break
            }
        }
    })

    // Restrict the lookback frames to avoid runaway lookups
    pcs := make([]uintptr, maximumCallerDepth)
    depth := runtime.Callers(minimumCallerDepth, pcs)
    frames := runtime.CallersFrames(pcs[skip:depth])

    for f, again := frames.Next(); again; f, again = frames.Next() {
        pkg := getPackageName(f.Function)

        // If the caller isn't part of this package, we're done
        if pkg != callerPackage {
            return &f //nolint:scopelint
        }
    }

    // if we got here, we failed to find the caller's context
    return nil
}

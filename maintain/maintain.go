// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-08
// Based on adapter by liasica, magicrolan@qq.com.

package maintain

import "os"

var (
    maintainFile = "runtime/MAINTAIN"
)

func SetFile(path string) {
    maintainFile = path
}

func File() string {
    return maintainFile
}

func Create() error {
    f, err := os.OpenFile(maintainFile, os.O_CREATE, 0644)
    if err != nil {
        return err
    }
    return f.Close()
}

func Remove() (err error) {
    if Exists() {
        err = os.Remove(maintainFile)
    }
    return
}

func Exists() bool {
    _, err := os.Stat(maintainFile)
    return err == nil
}

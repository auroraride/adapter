// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-06
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

func RemoveSliceDuplicate[T comparable](items []T) (list []T) {
    m := make(map[T]struct{})
    for _, item := range items {
        if _, ok := m[item]; !ok {
            m[item] = struct{}{}
            list = append(list, item)
        }
    }
    return
}

func ConvertSliceToMap[T any, K comparable](arr []T, f func(T) (K, T, bool)) (data map[K]T) {
    if len(arr) == 0 {
        return
    }
    data = make(map[K]T)
    for _, t := range arr {
        k, v, ok := f(t)
        if ok {
            data[k] = v
        } else {
            data[k] = t
        }
    }
    return
}

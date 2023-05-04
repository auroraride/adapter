// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-17
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

func ChSafeClose[T any](ch chan T) (justClosed bool) {
	defer func() {
		if recover() != nil {
			// The return result can be altered
			// in a defer function call.
			justClosed = false
		}
	}()

	// assume ch != nil here.
	close(ch)   // panic if ch is closed
	return true // <=> justClosed = true; return
}

func ChSafeSend[T any](ch chan T, value T) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()

	ch <- value  // panic if ch is closed
	return false // <=> closed = false; return
}

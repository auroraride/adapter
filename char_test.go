// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-06
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "testing"
)

func TestChar(t *testing.T) {
    if string(Newline) != "\n" {
        t.Fail()
    }
    if string(Hyphen) != "-" {
        t.Fail()
    }
    if string(Equal) != "=" {
        t.Fail()
    }
    if string(Colon) != ":" {
        t.Fail()
    }
    if string(Comma) != "," {
        t.Fail()
    }
    if string(Period) != "." {
        t.Fail()
    }
    if string(LeftSquareBracket) != "[" {
        t.Fail()
    }
    if string(RightSquareBracket) != "]" {
        t.Fail()
    }
    if string(LeftBracket) != "(" {
        t.Fail()
    }
    if string(RightBracket) != ")" {
        t.Fail()
    }
    if string(LeftCurlyBracket) != "{" {
        t.Fail()
    }
    if string(RightCurlyBracket) != "}" {
        t.Fail()
    }
}

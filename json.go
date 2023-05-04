// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-04
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

// DefaultJSONSerializer implements JSON encoding using encoding/jsoniter.
type DefaultJSONSerializer struct{}

// Serialize converts an interface into a json and writes it to the response.
// You can optionally use the indent parameter to produce pretty JSONs.
func (d DefaultJSONSerializer) Serialize(c echo.Context, i interface{}, indent string) error {
	enc := jsoniter.NewEncoder(c.Response())
	if indent != "" {
		enc.SetIndent("", indent)
	}
	return enc.Encode(i)
}

// Deserialize reads a JSON from a request body and converts it into an interface.
func (d DefaultJSONSerializer) Deserialize(c echo.Context, i interface{}) error {
	err := jsoniter.NewDecoder(c.Request().Body).Decode(i)
	return err
}

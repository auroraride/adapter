// Copyright (C) adapter. 2024-present.
//
// Created at 2024-03-25, by liasica

package pb

import jsoniter "github.com/json-iterator/go"

func (c *CabinetExchangeResponse) MarshalBinary() ([]byte, error) {
	return jsoniter.Marshal(c)
}

func (c *CabinetExchangeResponse) UnmarshalBinary(data []byte) error {
	return jsoniter.Unmarshal(data, c)
}

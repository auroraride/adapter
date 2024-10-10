// Copyright (C) cabservd. 2024-present.
//
// Created at 2024-05-30, by liasica

package es

import (
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"github.com/stretchr/testify/require"
)

type Document struct {
	Serial    string    `json:"serial"`
	Value     float64   `json:"value"`
	Timestamp time.Time `json:"@timestamp"`
}

func TestEs(t *testing.T) {
	err := Create("b19FdHpJOEI1VEVyeZVjOUs1WUM6MHIxUktUTGpROS1VaGdzZXRMQkVUQQ==", "development-electric-consumption-cabinet-", []string{"http://172.16.1.108:9200"})
	require.NoError(t, err)
	New().CreateDocument(&Document{
		Serial:    "x-test",
		Value:     198.00,
		Timestamp: time.Now(),
	})
	require.NoError(t, err)
	New().CreateDocument(&Document{
		Serial:    "x-test",
		Value:     98.00,
		Timestamp: time.Now().AddDate(0, 0, -1),
	})
}

func TestSearch(t *testing.T) {
	err := Create("b019FdHpJOEI1VEVyeZVjOUs1WUM6MHIxUktUTGpROS1VaGdzZXRMQkVUQQ==", "electric-consumption-cabinet-", []string{"http://172.16.1.108:9200"})
	require.NoError(t, err)

	items := NewSearch[Document](New().GetIndexWizard()).DoRequest(&search.Request{
		Query: &types.Query{
			Match: map[string]types.MatchQuery{
				"serial": {Query: "x-test"},
			},
		},
		Sort: []types.SortCombinations{
			types.SortOptions{SortOptions: map[string]types.FieldSort{
				DefaultFieldTimestamp: {Order: &sortorder.Asc},
			}},
		},
	})
	t.Log(items)
}

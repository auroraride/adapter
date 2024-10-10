// Copyright (C) adapter. 2024-present.
//
// Created at 2024-10-10, by liasica

package es

import (
	"context"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"

	"github.com/auroraride/adapter/log"
)

type ElasticSearch[T any] struct {
	*search.Search
	Options *SearchOptions
}

// SearchOptions 搜索选项
// 参考文档 https://www.elastic.co/guide/en/elasticsearch/reference/8.11/search-search.html#search-search-api-query-params
type SearchOptions struct {
	// 每页返回文档数量
	Size int

	// 总返回文档数量
	Pick int

	// 排序规则
	Sort map[string]types.FieldSort
}

type SearchOption func(*SearchOptions)

// SearchWithSize 设置每页返回的文档数量
func SearchWithSize(size int) SearchOption {
	return func(o *SearchOptions) {
		o.Size = size
	}
}

// SearchWithPick 设置总返回的文档数量
func SearchWithPick(num int) SearchOption {
	return func(o *SearchOptions) {
		o.Pick = num
	}
}

// SearchWithSort 设置排序规则
func SearchWithSort(field string, sort types.FieldSort) SearchOption {
	return func(o *SearchOptions) {
		if o.Sort == nil {
			o.Sort = make(map[string]types.FieldSort)
		}
		o.Sort[field] = sort
	}
}

// NewSearch 创建搜索
func NewSearch[T any](index string) *ElasticSearch[T] {
	return &ElasticSearch[T]{
		Search: GetInstance().client.Search().Index(index),
		Options: &SearchOptions{
			Size: 999,
		},
	}
}

// DoRequest 请求搜索
// https://www.elastic.co/guide/en/elasticsearch/reference/8.11/paginate-search-results.html#search-after
// https://ost.51cto.com/posts/11773
func (s *ElasticSearch[T]) DoRequest(request *search.Request, options ...SearchOption) (output []*T) {
	// 应用选项
	for _, option := range options {
		option(s.Options)
	}

	if s.Options.Size > 0 {
		request.Size = &s.Options.Size
	}

	if s.Options.Sort != nil {
		request.Sort = []types.SortCombinations{
			types.SortOptions{SortOptions: s.Options.Sort},
		}
	}

	res, err := s.Request(request).Do(context.Background())

	if err != nil {
		zap.L().Error("搜索失败", logTag(), log.Payload(request), zap.Error(err))
		return nil
	}

	hits := res.Hits.Hits

	// 如果没有命中文档，直接返回
	if len(hits) == 0 {
		return
	}

	// 解析命中的文档
	for _, hit := range hits {
		request.SearchAfter = hit.Sort
		item := new(T)
		err = jsoniter.Unmarshal(hit.Source_, item)
		if err != nil {
			zap.L().Error("搜索结果解析失败", logTag(), zap.Error(err), zap.String("source", string(hit.Source_)))
			continue
		}
		output = append(output, item)
	}

	// 如果需要返回的文档数量等于请求的文档数量，说明已经请求完毕
	if s.Options.Pick == len(output) {
		return
	}

	// 如果返回的文档数量小于请求的文档数量，说明已经请求完毕
	if len(hits) < s.Options.Size {
		return
	}

	// 如果返回的文档数量等于总文档数量，说明已经请求完毕
	if len(output) == int(res.Hits.Total.Value) {
		return
	}

	// 递归请求
	output = append(output, s.DoRequest(request)...)

	return
}

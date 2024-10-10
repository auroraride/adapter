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

	options *searciOption
}

// searciOption 搜索选项
// 参考文档 https://www.elastic.co/guide/en/elasticsearch/reference/8.11/search-search.html#search-search-api-query-params
type searciOption struct {
	// 索引
	index string

	// 每页返回文档数量
	size int

	// 总返回文档数量
	pick *int

	// 排序规则
	sort map[string]types.FieldSort
}

type SearchOption interface {
	apply(*searciOption)
}

type searchOptionFunc func(*searciOption)

func (f searchOptionFunc) apply(option *searciOption) {
	f(option)
}

func SearchWithIndex(index string) SearchOption {
	return searchOptionFunc(func(o *searciOption) {
		o.index = index
	})
}

// SearchWithSize 设置每页返回的文档数量
func SearchWithSize(size int) SearchOption {
	return searchOptionFunc(func(o *searciOption) {
		o.size = size
	})
}

// SearchWithPick 设置总返回的文档数量
func SearchWithPick(num int) SearchOption {
	return searchOptionFunc(func(o *searciOption) {
		o.pick = &num
	})
}

// SearchWithSort 设置排序规则
func SearchWithSort(field string, sort types.FieldSort) SearchOption {
	return searchOptionFunc(func(o *searciOption) {
		if o.sort == nil {
			o.sort = make(map[string]types.FieldSort)
		}
		o.sort[field] = sort
	})
}

// NewSearch 创建搜索
func NewSearch[T any](es *Elastic, options ...SearchOption) *ElasticSearch[T] {
	o := &searciOption{}

	index := es.GetIndexWizard()

	// 应用选项
	for _, option := range options {
		option.apply(o)
	}

	if o.index != "" {
		index = o.index
	}

	return &ElasticSearch[T]{
		Search:  es.client.Search().Index(index),
		options: o,
	}
}

// DoRequest 请求搜索
// https://www.elastic.co/guide/en/elasticsearch/reference/8.11/paginate-search-results.html#search-after
// https://ost.51cto.com/posts/11773
func (s *ElasticSearch[T]) DoRequest(req *search.Request) (output []*T) {
	if req.Size == nil {
		req.Size = &s.options.size
	}

	if *req.Size == 0 {
		*req.Size = 999
	}

	if s.options.sort != nil {
		req.Sort = []types.SortCombinations{
			types.SortOptions{SortOptions: s.options.sort},
		}
	}

	res, err := s.Request(req).Do(context.Background())

	if err != nil {
		zap.L().Error("搜索失败", logTag(), log.Payload(req), zap.Error(err))
		return nil
	}

	hits := res.Hits.Hits

	// 如果没有命中文档，直接返回
	if len(hits) == 0 {
		return
	}

	// 解析命中的文档
	for _, hit := range hits {
		req.SearchAfter = hit.Sort
		item := new(T)
		err = jsoniter.Unmarshal(hit.Source_, item)
		if err != nil {
			zap.L().Error("搜索结果解析失败", logTag(), zap.Error(err), zap.String("source", string(hit.Source_)))
			continue
		}
		output = append(output, item)
	}

	// 如果需要返回的文档数量小于或等于请求的文档数量，说明已经请求完毕
	if s.options.pick != nil && *s.options.pick <= len(output) {
		output = output[:*s.options.pick]
		return
	}

	// 如果返回的文档数量小于请求的文档数量，则是所有数据
	if len(hits) < s.options.size {
		return
	}

	// 如果返回的文档数量等于总文档数量，说明已经请求完毕
	if len(output) == int(res.Hits.Total.Value) {
		return
	}

	// 递归请求
	output = append(output, s.DoRequest(req)...)

	return
}

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

type Pagination struct {
	Total       int64  `json:"total"`
	SearchAfter string `json:"searchAfter"`
}

type ElasticSearch[T any] struct {
	*search.Search

	options *searchOption
}

// searchOption 搜索选项
// 参考文档 https://www.elastic.co/guide/en/elasticsearch/reference/8.11/search-search.html#search-search-api-query-params
type searchOption struct {
	// 索引
	index string

	// 每页返回文档数量
	size int

	// 每页最大返回文档数量
	maxSize int

	// 总返回文档数量
	pick *int

	// 排序规则
	sort map[string]types.FieldSort
}

type SearchOption interface {
	apply(*searchOption)
}

type searchOptionFunc func(*searchOption)

func (f searchOptionFunc) apply(option *searchOption) {
	f(option)
}

func SearchWithIndex(index string) SearchOption {
	return searchOptionFunc(func(o *searchOption) {
		o.index = index
	})
}

// SearchWithSize 设置每页返回的文档数量
func SearchWithSize(size int) SearchOption {
	return searchOptionFunc(func(o *searchOption) {
		o.size = size
	})
}

// SearchWithPick 设置总返回的文档数量
func SearchWithPick(num int) SearchOption {
	return searchOptionFunc(func(o *searchOption) {
		o.pick = &num
	})
}

// SearchWithSort 设置排序规则
func SearchWithSort(field string, sort types.FieldSort) SearchOption {
	return searchOptionFunc(func(o *searchOption) {
		if o.sort == nil {
			o.sort = make(map[string]types.FieldSort)
		}
		o.sort[field] = sort
	})
}

func SearchWithMaxSize(size int) SearchOption {
	return searchOptionFunc(func(o *searchOption) {
		o.maxSize = size
	})
}

// NewSearch 创建搜索
func NewSearch[T any](es *Elastic, options ...SearchOption) *ElasticSearch[T] {
	o := &searchOption{
		maxSize: 500,
	}

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

// doRequest 请求搜索
func (s *ElasticSearch[T]) doRequest(req *search.Request) *search.Response {
	if req.Size == nil {
		req.Size = &s.options.size
	}

	if *req.Size == 0 || *req.Size > s.options.maxSize {
		*req.Size = s.options.maxSize
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

	return res
}

// parseResponse 解析搜索结果
func (s *ElasticSearch[T]) parseResponse(res *search.Response) (output []*T) {
	if res == nil {
		return
	}

	hits := res.Hits.Hits

	// 如果没有命中文档，直接返回
	if len(hits) == 0 {
		return
	}

	// 解析命中的文档
	for _, hit := range hits {
		item := new(T)
		err := jsoniter.Unmarshal(hit.Source_, item)
		if err != nil {
			zap.L().Error("搜索结果解析失败", logTag(), zap.Error(err), zap.String("source", string(hit.Source_)))
			continue
		}
		output = append(output, item)
	}

	return
}

// DoRequest 请求搜索 <递归不分页>
// https://www.elastic.co/guide/en/elasticsearch/reference/8.11/paginate-search-results.html#search-after
// https://ost.51cto.com/posts/11773
func (s *ElasticSearch[T]) DoRequest(req *search.Request) (output []*T) {
	res := s.doRequest(req)

	// 解析结果
	output = s.parseResponse(res)

	// 修改下次请求的游标
	req.SearchAfter = res.Hits.Hits[len(res.Hits.Hits)-1].Sort

	// 如果需要返回的文档数量小于或等于请求的文档数量，说明已经请求完毕
	if s.options.pick != nil && *s.options.pick <= len(output) {
		output = output[:*s.options.pick]
		return
	}

	// 如果返回的文档数量小于请求的文档数量，则是所有数据
	if len(res.Hits.Hits) < s.options.size {
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

// DoRequestWithResponse 请求搜索，不递归搜索全部文档，返回原始结果
func (s *ElasticSearch[T]) DoRequestWithResponse(req *search.Request) (output []*T, res *search.Response) {
	res = s.doRequest(req)
	output = s.parseResponse(res)
	return
}

package main

import (
	"testing"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type Filter struct {
	Field      string        `json:"field,omitempty"`
	Gt         interface{}   `json:"gt,omitempty"`
	Gte        interface{}   `json:"gte,omitempty"`
	Lt         interface{}   `json:"lt,omitempty"`
	Lte        interface{}   `json:"lte,omitempty"`
	Equals     []interface{} `json:"equals,omitempty"`
	Or         []*Filter     `json:"or,omitempty"`
	And        []*Filter     `json:"and,omitempty"`
	Search     interface{}   `json:"search,omitempty"`
	NotEqual   []interface{} `json:"not_equal,omitempty"`
	MultiMatch []interface{} `json:"multi_match,omitempty"`
	Exists     interface{}   `json:"exists,omitempty"`
}

func populateNumberRangeQuery(numberRangeQuery *types.NumberRangeQuery, gt, gte, lt, lte *float64) {
	var floatValue types.Float64
	if gt != nil {
		floatValue = types.Float64(*gt)
		numberRangeQuery.Gt = &floatValue
	}
	if gte != nil {
		floatValue = types.Float64(*gte)
		numberRangeQuery.Gte = &floatValue
	}
	if lt != nil {
		floatValue = types.Float64(*lt)
		numberRangeQuery.Lt = &floatValue
	}
	if lte != nil {
		floatValue = types.Float64(*lte)
		numberRangeQuery.Lte = &floatValue
	}
}

func populateNumberRangeQuery2(numberRangeQuery *types.NumberRangeQuery, gt, gte, lt, lte *float64) *types.NumberRangeQuery {
	var floatValue types.Float64
	if gt != nil {
		floatValue = types.Float64(*gt)
		numberRangeQuery.Gt = &floatValue
	}
	if gte != nil {
		floatValue = types.Float64(*gte)
		numberRangeQuery.Gte = &floatValue
	}
	if lt != nil {
		floatValue = types.Float64(*lt)
		numberRangeQuery.Lt = &floatValue
	}
	if lte != nil {
		floatValue = types.Float64(*lte)
		numberRangeQuery.Lte = &floatValue
	}

	return numberRangeQuery
}

func BenchmarkMethod1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := &types.NumberRangeQuery{}
		ff := float64(10000)
		populateNumberRangeQuery(f, &ff, &ff, &ff, &ff)

	}
}

func BenchmarkMethod2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := &types.NumberRangeQuery{}
		ff := float64(10000)

		f = populateNumberRangeQuery2(f, &ff, &ff, &ff, &ff)
	}
}

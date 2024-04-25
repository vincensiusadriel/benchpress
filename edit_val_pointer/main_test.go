package main

import (
	"testing"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

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

func populateNumberRangeQuery2(numberRangeQuery types.NumberRangeQuery, gt, gte, lt, lte *float64) types.NumberRangeQuery {
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

// func BenchmarkMethod1(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		f := types.NumberRangeQuery{}
// 		ff := float64(10000)
// 		populateNumberRangeQuery(&f, &ff, &ff, &ff, &ff)

// 		if ff != float64(*f.Gt) {
// 			panic("not same")
// 		}

// 	}
// }

// func BenchmarkMethod2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		f := types.NumberRangeQuery{}
// 		ff := float64(10000)

// 		f = populateNumberRangeQuery2(f, &ff, &ff, &ff, &ff)
// 		if ff != float64(*f.Gt) {
// 			panic("not same")
// 		}
// 	}
// }

type A struct {
	X, Y, Z, F, G, H, J, K string
}

func (a *A) EditPointer(str string) {
	a.X = str
	a.Y = str
	a.Z = str
	a.F = str
	a.G = str
	a.H = str
	a.J = str
	a.K = str
}

func (a A) EditNon(str string) A {
	a.X = str
	a.Y = str
	a.Z = str
	a.F = str
	a.G = str
	a.H = str
	a.J = str
	a.K = str

	return a
}

const STR = "hi hello world boi some boi boi yo boi boi test boi boi asddhlskdjflaksdjflakjdfalkjdsfhlakjdfhlaskjdfhlakjdfblaskdfblakjdsblskaj vlaskdj laksdj flakjd flkas dflakjsd flakjsd flajskd flaksjdfalksdjfhlasjdhlfajkhdflakjhfd"

func BenchmarkMethodPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := &A{}

		a.EditPointer(STR)

		if a.K != STR {
			panic("not same")
		}
	}
}

func BenchmarkMethodNonPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := A{}

		temp := a.EditNon(STR)

		if temp.K != STR {
			panic("not same")
		}
	}
}

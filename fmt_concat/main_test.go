package main

import (
	"fmt"
	"strconv"
	"testing"
)

const (
	cacheKey            = "cache:userlabel:%v:%v:%v"
	uid          int64  = 10000
	label        string = "some label"
	contextValue string = "some context"
)

func withFmt() string {
	return fmt.Sprintf(cacheKey, uid, label, contextValue)
}

func withConcat() string {
	return "cache:userlabel:" + strconv.FormatInt(uid, 10) + ":" + label + ":" + contextValue

}

func BenchmarkMethodWithConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withConcat()
	}
}

func BenchmarkMethodWithFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withFmt()
	}
}

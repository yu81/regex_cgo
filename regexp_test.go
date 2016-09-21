package cregexp

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestRegmatch(t *testing.T) {
	assertion := assert.New(t)
	okPatternResult := RegMatch("/path/to/hoge/1234", "/path/to/hoge/[0-9]+", REG_EXTENDED)
	ngPatternResult := RegMatch("/path/to/fuga/1234", "/path/to/hoge/[0-9]+", REG_EXTENDED)
	assertion.True(okPatternResult)
	assertion.False(ngPatternResult)
}

func BenchmarkRegmatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RegMatch("/path/to/hoge/1234", "/path/to/hoge/[0-9]+", REG_EXTENDED)
	}
}

var compiledRegexp = regexp.MustCompile("/path/to/hoge/[0-9]+")

func BenchmarkRegexpMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		compiledRegexp.MatchString("/path/to/hoge/1234")
	}
}

package yesql

import (
	"strings"
	"unicode"
)

// simpleNamingStrategy simple naming strategy
type simpleNamingStrategy struct{}

// fieldUpNamingStrategy field up runne naming strategy
type fieldUpNamingStrategy struct{}

func (simpleNamingStrategy) FiledNaming(name string) string {
	return name
}

func (fieldUpNamingStrategy) FiledNaming(name string) string {
	buf := &strings.Builder{}
	toUp := true
	for _, c := range name {
		if c == '_' {
			toUp = true
			continue
		}
		if toUp {
			buf.WriteRune(unicode.ToUpper(c))
			toUp = false
		} else {
			buf.WriteRune(c)
		}
	}
	return buf.String()
}

// NewSimpleNamingStrategy return a simple naming strategy instance
func NewSimpleNamingStrategy() NamingStrategy {
	return simpleNamingStrategy{}
}

// NewFieldUpNamingStrategy retuan a filed up runne naming stratefy instance
func NewFieldUpNamingStrategy() NamingStrategy {
	return fieldUpNamingStrategy{}
}

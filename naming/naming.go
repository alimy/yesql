package naming

import (
	"strings"
	"unicode"
)

var (
	_defaultNamingStrategy = NewFieldUpNamingStrategy()
)

// NamingStrategy naming strategy interface
type NamingStrategy interface {
	Naming(string) string
}

// simpleNamingStrategy simple naming strategy
type simpleNamingStrategy struct{}

// fieldUpNamingStrategy field up runne naming strategy
type fieldUpNamingStrategy struct{}

func (simpleNamingStrategy) Naming(name string) string {
	return name
}

func (fieldUpNamingStrategy) Naming(name string) string {
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

func SetNamingStrategy(ns NamingStrategy) {
	if ns != nil {
		_defaultNamingStrategy = ns
	}
}

func Naming(name string) string {
	return _defaultNamingStrategy.Naming(name)
}

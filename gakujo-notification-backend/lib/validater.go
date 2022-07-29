package lib

import (
	"unicode"
)

type ValidateOption struct {
	minLen         int
	maxLen         int
	allowMultibyte bool
	whiteList      map[rune]struct{}
	blackList      map[rune]struct{}
}

type ValidateOptionFunc func(opt *ValidateOption)

var AlphanumericCharacters = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '_', '-'}

func NewValidateOption(opts ...ValidateOptionFunc) *ValidateOption {
	opt := DefaultValidateOption()
	for _, f := range opts {
		f(opt)
	}
	return opt
}

func DefaultValidateOption() *ValidateOption {
	return &ValidateOption{
		minLen:         0,
		maxLen:         1<<31 - 1,
		allowMultibyte: false,
		whiteList:      make(map[rune]struct{}),
		blackList:      make(map[rune]struct{}),
	}
}

func MinLen(n int) ValidateOptionFunc {
	return func(opt *ValidateOption) {
		opt.minLen = n
	}
}

func MaxLen(n int) ValidateOptionFunc {
	return func(opt *ValidateOption) {
		opt.maxLen = n
	}
}

func AllowMultibyte() ValidateOptionFunc {
	return func(opt *ValidateOption) {
		opt.allowMultibyte = true
	}
}

func WhiteList(r []rune) ValidateOptionFunc {
	mp := make(map[rune]struct{})
	for _, r := range r {
		mp[r] = struct{}{}
	}
	return func(opt *ValidateOption) {
		opt.whiteList = mp
	}
}

func BlackList(r []rune) ValidateOptionFunc {
	mp := make(map[rune]struct{})
	for _, r := range r {
		mp[r] = struct{}{}
	}
	return func(opt *ValidateOption) {
		opt.blackList = mp
	}
}

func ValidateString(s string, opts ...ValidateOptionFunc) bool {
	opt := NewValidateOption(opts...)
	return opt.ValidateString(s)
}

func (opt *ValidateOption) ValidateString(s string) bool {
	var length int
	if opt.allowMultibyte {
		for _, r := range s {
			if !unicode.IsPrint(r) {
				return false
			}
		}
		length = len([]rune(s))
	} else {
		if len(s) != len([]rune(s)) {
			return false
		}
		length = len(s)
	}
	if length < opt.minLen {
		return false
	}
	if length > opt.maxLen {
		return false
	}
	for _, r := range s {
		if _, ok := opt.whiteList[r]; !ok {
			return false
		}
	}
	for _, r := range s {
		if _, ok := opt.blackList[r]; ok {
			return false
		}
	}
	return true
}

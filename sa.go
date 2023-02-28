package bwt

import (
	"fmt"
	"math"
)

type SuffixConstraints interface {
	SuffixArray | SampleSuffixArray
}

type Suffix interface {
	Get(index int) (int, bool)
	Set(index, value int)
}

type OptionsSuffix struct {
	mod int
}

func WithCompression(mod int) func(*OptionsSuffix) {
	return func(s *OptionsSuffix) {
		s.mod = mod
	}
}

func NewSuffix[T SuffixConstraints](size int, options ...func(*OptionsSuffix)) Suffix {
	opts := &OptionsSuffix{}
	for _, o := range options {
		o(opts)
	}

	if opts.mod < 2 {
		opts.mod = 2
	}

	suffix := new(T)
	switch t := any(suffix).(type) {
	case *SuffixArray:
		t.sa = make([]int, size)
		return t
	case *SampleSuffixArray:
		t.opts = *opts
		l := int(math.Ceil(float64(size) / float64(opts.mod)))
		t.sa = make([]int, l)
		t.size = l
		t.length = size
		return t
	}
	return nil
}

type SuffixArray struct {
	sa []int
}

func (s *SuffixArray) Get(index int) (int, bool) {
	if index < 0 || index >= s.Length() {
		panic(fmt.Sprintf("index %v out of range", index))
	}

	return s.sa[index], true
}

func (s *SuffixArray) Set(index, value int) {
	if index < 0 || index >= s.Length() {
		panic(fmt.Sprintf("index %v out of range", index))
	}
	s.sa[index] = value

}

func (s *SuffixArray) Length() int {
	return len(s.sa)
}

func (s *SampleSuffixArray) Get(index int) (int, bool) {
	if index < 0 || index >= s.Length() {
		panic(fmt.Sprintf("index %v out of range", index))
	}

	if index%s.opts.mod == 0 {
		i := index / s.size
		return s.sa[i], true
	}

	return 0, false
}

func (s *SampleSuffixArray) get(index, count int) int {
	if index%s.opts.mod == 0 {
		i := index / s.size
		return s.sa[i] - count
	}

	index++
	count++
	return s.get(index, count)
}

func (s *SampleSuffixArray) Set(index, value int) {
	if index < 0 || index >= s.Length() {
		panic(fmt.Sprintf("index %v out of range", index))
	}

	if index%s.opts.mod == 0 {
		i := index / s.size
		s.sa[i] = value
	}
}

type SampleSuffixArray struct {
	sa     []int
	opts   OptionsSuffix
	size   int
	length int
}

func (s *SampleSuffixArray) Length() int {
	return s.length
}

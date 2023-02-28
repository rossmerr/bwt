package suffixarray

import (
	"math"
)

type Suffix interface {
	Has(index int) bool
	Get(index int) int
	Set(index, value int)
	Enumerate() SuffixIterator
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

type SuffixIterator interface {
	HasNext() bool
	Next() (int, int)
}

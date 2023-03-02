package suffixarray

import "fmt"

type SampleSuffixArray struct {
	sa      []int
	opts    OptionsSuffix
	size    int
	length  int
	version int
}

func (s *SampleSuffixArray) Has(index int) bool {
	if index < 0 || index >= s.Length() {
		panic(fmt.Sprintf("index %v out of range", index))
	}

	return index%s.opts.mod == 0
}

func (s *SampleSuffixArray) Get(index int) int {
	if index < 0 || index >= s.Length() {
		panic(fmt.Sprintf("index %v out of range", index))
	}

	if index%s.opts.mod == 0 {
		i := index / s.opts.mod
		return s.sa[i]
	}

	return -1
}

// func (s *SampleSuffixArray) get(index, count int) int {
// 	if index < 0 || index >= s.Length() {
// 		panic(fmt.Sprintf("index %v out of range", index))
// 	}

// 	if index%s.opts.mod == 0 {
// 		i := index / s.opts.mod
// 		return s.sa[i] + count
// 	}

// 	return 0
// }

// func (s *SampleSuffixArray) walk(i, count int) int {
// 	if s.Has(i) {
// 		return s.get(i, count)
// 	} else {
// 		return s.walk(i+1, count+1)
// 	}
// }

func (s *SampleSuffixArray) Set(index, value int) {
	if index < 0 || index >= s.Length() {
		panic(fmt.Sprintf("index %v out of range", index))
	}

	if index%s.opts.mod == 0 {
		i := index / s.opts.mod
		s.sa[i] = value
	}

	s.version++
}

func (s *SampleSuffixArray) Length() int {
	return s.length
}

func (s *SampleSuffixArray) Enumerate() SuffixIterator {
	return NewSampleSuffixArrayIterator(s)
}

func NewSampleSuffixArrayIterator(suffix *SampleSuffixArray) *SampleSuffixArrayIterator {

	return &SampleSuffixArrayIterator{
		suffix:     suffix,
		indexStart: 0,
		indexEnd:   suffix.size,
		version:    suffix.version,
	}
}

type SampleSuffixArrayIterator struct {
	suffix     *SampleSuffixArray
	version    int
	indexStart int
	indexEnd   int
}

func (s *SampleSuffixArrayIterator) HasNext() bool {
	return s.indexStart < s.indexEnd
}

func (s *SampleSuffixArrayIterator) Next() (int, int) {
	if s.version != s.suffix.version {
		panic("version failed")
	}

	if s.indexStart < s.indexEnd {
		index := s.indexStart
		currentElement := s.suffix.Get(index)
		s.indexStart += s.suffix.opts.mod
		return currentElement, index
	}

	s.indexStart = s.indexEnd

	return 0, s.indexStart
}

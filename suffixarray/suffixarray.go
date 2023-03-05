package suffixarray

import "fmt"

type SuffixArray struct {
	sa      []int
	version int
}

func NewSuffixArray(size int) Suffix {

	suffix := &SuffixArray{
		sa: make([]int, size),
	}

	return suffix
}

func (s *SuffixArray) Get(index int) int {
	if index < 0 || index >= s.Length() {
		panic(fmt.Sprintf("index %v out of range", index))
	}

	return s.sa[index]
}

func (s *SuffixArray) Set(index, value int) {
	if index < 0 || index >= s.Length() {
		panic(fmt.Sprintf("index %v out of range", index))
	}
	s.sa[index] = value

	s.version++
}

func (s *SuffixArray) Has(index int) bool {
	if index < 0 || index >= s.Length() {
		panic(fmt.Sprintf("index %v out of range", index))
	}

	return true
}

func (s *SuffixArray) Length() int {
	return len(s.sa)
}

func (s *SuffixArray) Enumerate() SuffixIterator {
	return NewSuffixArrayIterator(s)
}

func NewSuffixArrayIterator(suffix *SuffixArray) *SuffixArrayIterator {

	return &SuffixArrayIterator{
		suffix:     suffix,
		indexStart: 0,
		indexEnd:   suffix.Length(),
		version:    suffix.version,
	}
}

type SuffixArrayIterator struct {
	suffix     *SuffixArray
	version    int
	indexStart int
	indexEnd   int
}

func (s *SuffixArrayIterator) HasNext() bool {
	return s.indexStart < s.indexEnd
}

func (s *SuffixArrayIterator) Next() (int, int) {
	if s.version != s.suffix.version {
		panic("version failed")
	}

	if s.indexStart < s.suffix.Length() {
		index := s.indexStart
		currentElement := s.suffix.Get(index)
		s.indexStart++
		return currentElement, index
	}

	s.indexStart = s.suffix.Length()

	return 0, s.indexStart
}

package suffixarray

import (
	"fmt"
	"math"
)

type SampleSuffixArray struct {
	sa         []int
	sampleRate int
	size       int
	length     int
	version    int
}

func NewSampleSuffixArray(size, sampleRate int) Suffix {
	l := int(math.Ceil(float64(size) / float64(sampleRate)))

	suffix := &SampleSuffixArray{
		sampleRate: sampleRate,
		sa:         make([]int, l),
		size:       l,
		length:     size,
	}

	return suffix
}

func (s *SampleSuffixArray) Has(index int) bool {
	if index < 0 || index >= s.Length() {
		panic(fmt.Sprintf("index %v out of range", index))
	}

	return index%s.sampleRate == 0
}

func (s *SampleSuffixArray) Get(index int) int {
	return s.walk(index, 0)
}

func (s *SampleSuffixArray) get(index, count int) int {
	if index < 0 || index >= s.Length() {
		panic(fmt.Sprintf("index %v out of range", index))
	}

	if index%s.sampleRate == 0 {
		i := index / s.sampleRate
		return s.sa[i] + count
	}

	return -1
}

func (s *SampleSuffixArray) walk(i, count int) int {
	if s.Has(i) {
		return s.get(i, count)
	} else {
		return s.walk(i-1, count+1)
	}
}

func (s *SampleSuffixArray) Set(index, value int) {
	if index < 0 || index >= s.Length() {
		panic(fmt.Sprintf("index %v out of range", index))
	}

	if index%s.sampleRate == 0 {
		i := index / s.sampleRate
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
		suffix:       suffix,
		indexStart:   0,
		currentIndex: 0,
		indexEnd:     suffix.size,
		version:      suffix.version,
	}
}

type SampleSuffixArrayIterator struct {
	suffix       *SampleSuffixArray
	version      int
	indexStart   int
	indexEnd     int
	currentIndex int
}

func (s *SampleSuffixArrayIterator) HasNext() bool {
	return s.indexStart < s.indexEnd
}

func (s *SampleSuffixArrayIterator) Next() (int, int) {
	if s.version != s.suffix.version {
		panic("version failed")
	}

	if s.indexStart < s.indexEnd {
		currentIndex := s.currentIndex
		currentElement := s.suffix.Get(currentIndex)
		s.indexStart++
		s.currentIndex += s.suffix.sampleRate
		return currentElement, currentIndex
	}

	s.indexStart = s.indexEnd

	return 0, s.indexStart
}

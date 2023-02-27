package bwt

type SuffixConstraints interface {
	SuffixArray | SampleSuffixArray
}

type Suffix interface {
	Get(index int) int
	Append(s int)
}

func NewSuffix[T SuffixConstraints]() Suffix {
	suffix := new(T)
	switch t := any(suffix).(type) {
	case *SuffixArray:
		t.sa = []int{}
		return t
	case *SampleSuffixArray:
		t.sa = []int{}
		return t
	}
	return nil
}

type SuffixArray struct {
	sa []int
}

func NewSuffixArray() *SuffixArray {
	return &SuffixArray{
		sa: []int{},
	}
}

func (s *SuffixArray) Get(index int) int {
	return s.sa[index]
}

func (s *SuffixArray) Append(index int) {
	s.sa = append(s.sa, index)
}

func NewSampleSuffixArray() *SampleSuffixArray {
	return &SampleSuffixArray{
		sa: []int{},
	}
}

func (s *SampleSuffixArray) Get(index int) int {
	return s.sa[index]
}

func (s *SampleSuffixArray) Append(index int) {
	s.sa = append(s.sa, index)
}

type SampleSuffixArray struct {
	sa []int
}

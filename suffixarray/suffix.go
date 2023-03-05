package suffixarray

type Suffix interface {
	Has(index int) bool
	Get(index int) int
	Set(index, value int)
	Enumerate() SuffixIterator
}

type SuffixIterator interface {
	HasNext() bool
	Next() (int, int)
}

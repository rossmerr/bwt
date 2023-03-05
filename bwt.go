package bwt

import (
	"fmt"
	"sort"
	"strings"

	"github.com/rossmerr/bwt/suffixarray"
)

// Matrix of the BWT
func Matrix(str string, options ...func(*OptionsBwt)) ([][]rune, error) {
	opts := buildOptions(options)

	if strings.Contains(str, opts.ext) {
		err := fmt.Errorf("input string cannot contain EXT character")
		return [][]rune{}, err
	}

	str = str + opts.ext
	size := len(str)

	matrix := make(matrix, size)

	for i := size - 1; i >= 0; i-- {
		matrix[i] = append([]rune(str[i:]), []rune(str[:i])...)
	}
	sort.Sort(&matrix)

	return matrix, nil
}

// Last column of the BWT matrix
// Optimized to only do the last rotation
func Last(str string, options ...func(*OptionsBwt)) ([]rune, error) {
	opts := buildOptions(options)

	appendFirst := func(i int, r rune) {
	}

	set := func(s, o int) {
	}
	last, err := firstLastSuffix(str, appendFirst, set, opts)
	return last, err
}

// First and Last column of the BWT matrix
// Optimized to only do the last rotation
func FirstLast(str string, options ...func(*OptionsBwt)) ([]rune, []rune, error) {
	opts := buildOptions(options)

	size := len(str + opts.ext)
	first := make([]rune, size)

	appendFirst := func(i int, r rune) {
		first[i] = r
	}

	set := func(s, o int) {
	}
	last, err := firstLastSuffix(str, appendFirst, set, opts)
	return first, last, err
}

// First and Last column of the BWT matrix with a SuffixArray
// The SuffixArray returns the offset of the original string relative to the first column of the BWT matrix
// Optimized to only do the last rotation
func FirstLastSuffix(str string, options ...func(*OptionsBwt)) ([]rune, []rune, suffixarray.Suffix, error) {
	opts := buildOptions(options)

	size := len(str + opts.ext)

	sa := suffixarray.NewSampleSuffixArray(size, opts.SampleRate())

	first := make([]rune, size)

	appendFirst := func(i int, r rune) {
		first[i] = r
	}

	last, err := firstLastSuffix(str, appendFirst, sa.Set, opts)
	return first, last, sa, err
}

func firstLastSuffix(str string, appendFirst func(i int, r rune), set func(index, value int), opts *OptionsBwt) ([]rune, error) {
	if strings.Contains(str, opts.ext) {
		err := fmt.Errorf("input string cannot contain EXT character")
		return []rune{}, err
	}

	str = str + opts.ext
	size := len(str)

	suffixes := make([]string, size)
	for i := 0; i < size; i++ {
		suffixes[i] = str[i:]
	}

	sort.Strings(suffixes)

	last := make([]rune, size)
	for i := 0; i < size; i++ {
		appendFirst(i, rune(suffixes[i][0]))
		s := size - len(suffixes[i])
		mod := (s + size - 1) % size
		last[i] = rune(str[mod])
		set(i, s)
	}

	return last, nil
}

// Reverse the BWT transformation, last column of the BWT matrix back to the original text
func Reverse(str string, options ...func(*OptionsBwt)) string {
	opts := buildOptions(options)

	size := len(str)
	table := make([]string, size)
	for range table {
		for i := 0; i < size; i++ {
			table[i] = str[i:i+1] + table[i]
		}
		sort.Strings(table)
	}
	for _, row := range table {
		if strings.HasPrefix(row, opts.ext) {
			return row[1:]
		}
	}
	return ""
}

type matrix [][]rune

func (m matrix) Len() int { return len(m) }
func (m matrix) Less(i, j int) bool {
	for x := range m[i] {
		if m[i][x] == m[j][x] {
			continue
		}
		return m[i][x] < m[j][x]
	}
	return false
}

func (m matrix) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

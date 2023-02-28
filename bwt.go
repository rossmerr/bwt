package bwt

import (
	"fmt"
	"sort"
	"strings"
)

const ext = "\003"

func Bwt(str string) (string, error) {
	appendFirst := func(i int, r rune) {
	}

	set := func(s, o int) {
	}
	last, err := bwtFirstLastSuffix(str, appendFirst, set)
	return last, err
}

func BwtFirstLast(str string) (string, string, error) {
	size := len(str + ext)
	first := make([]rune, size)

	appendFirst := func(i int, r rune) {
		first[i] = r
	}

	set := func(s, o int) {
	}
	last, err := bwtFirstLastSuffix(str, appendFirst, set)
	return string(first), last, err
}

func BwtFirstLastSuffix[T SuffixConstraints](str string, options ...func(*OptionsSuffix)) (string, string, Suffix, error) {

	size := len(str + ext)

	sa := NewSuffix[T](size, options...)

	first := make([]rune, size)

	appendFirst := func(i int, r rune) {
		first[i] = r
	}

	last, err := bwtFirstLastSuffix(str, appendFirst, sa.Set)
	return string(first), string(last), sa, err
}

func bwtFirstLastSuffix(str string, appendFirst func(i int, r rune), set func(index, value int)) (string, error) {
	if strings.Contains(str, ext) {
		err := fmt.Errorf("input string cannot contain EXT character")
		return "", err
	}

	str = str + ext
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

	return string(last), nil
}

func Ibwt(str string) string {
	size := len(str)
	table := make([]string, size)
	for range table {
		for i := 0; i < size; i++ {
			table[i] = str[i:i+1] + table[i]
		}
		sort.Strings(table)
	}
	for _, row := range table {
		if strings.HasPrefix(row, ext) {
			return row[1:]
		}
	}
	return ""
}

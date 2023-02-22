package bwt

import (
	"fmt"
	"sort"
	"strings"
)

const ext = "\003"

func Bwt(str string) (string, error) {
	_, last, _, err := BwtFirstLastSuffix(str)
	return last, err
}

func BwtFirstLast(str string) (string, string, error) {
	first, last, _, err := BwtFirstLastSuffix(str)
	return first, last, err
}

func BwtFirstLastSuffix(str string) (string, string, []int, error) {
	sa := []int{}
	if strings.Contains(str, ext) {
		err := fmt.Errorf("input string cannot contain EXT character")
		return "", "", sa, err
	}

	str = str + ext
	size := len(str)

	suffixes := make([]string, size)
	for i := 0; i < size; i++ {
		suffixes[i] = str[i:]
	}

	sort.Strings(suffixes)

	first := make([]rune, size)
	last := make([]rune, size)
	for i := 0; i < size; i++ {
		first[i] = rune(suffixes[i][0])
		s := size - len(suffixes[i])
		mod := (s + size - 1) % size
		last[i] = rune(str[mod])
		sa = append(sa, s)
	}

	return string(first), string(last), sa, nil
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

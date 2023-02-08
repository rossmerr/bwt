package bwt

import (
	"fmt"
	"sort"
	"strings"
)

const ext = "\003"

func Bwt(str string) (string, error) {
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

	result := make([]rune, size)
	for i := 0; i < size; i++ {
		sa := size - len(suffixes[i])
		mod := (sa + size - 1) % size
		result[i] = rune(str[mod])
	}

	return string(result), nil
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

type RuneRank struct {
	Value rune
	Rank  int
}

type RankedString []RuneRank

func (s RankedString) String() string {
	str := make([]rune, len(s))

	for i, r := range s {
		str[i] = r.Value
	}
	return string(str)
}

func BwtRank(str string) (RankedString, error) {
	size := len(str) + 1

	result := make(RankedString, size)

	if strings.Contains(str, ext) {
		err := fmt.Errorf("input string cannot contain EXT character")
		return result, err
	}

	str = str + ext

	rank := map[byte]int{}

	suffixes := make([]string, size)
	for i := 0; i < size; i++ {
		suffixes[i] = str[i:]
	}

	sort.Strings(suffixes)

	for i := 0; i < size; i++ {
		sa := size - len(suffixes[i])
		mod := (sa + size - 1) % size
		r := str[mod]
		result[i].Value = rune(r)

		if _, ok := rank[r]; ok {
			rank[r] = rank[r] + 1
		} else {
			rank[r] = 0
		}

		result[i].Rank = rank[r]
	}

	return result, nil
}

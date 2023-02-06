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
	Value string
	Rank  int
}

type RankedString []RuneRank

func (p RankedString) Len() int           { return len(p) }
func (p RankedString) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p RankedString) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (s RankedString) String() string {
	str := make([]string, len(s))

	for i, r := range s {
		str[i] = fmt.Sprintf("{%v, %v}", string(r.Value), r.Rank)
	}
	return "[" + strings.Join(str, " ") + "]"
}

func BwtRank(str string) (RankedString, error) {
	size := len(str) + 1

	result := make(RankedString, size)

	if strings.Contains(str, ext) {
		err := fmt.Errorf("input string cannot contain EXT character")
		return result, err
	}

	str = str + ext

	suffixes := make(RankedString, size)
	rank := map[byte]int{}

	for i := 0; i < size; i++ {
		s := str[i:]
		suffixes[i].Value = s
	}

	sort.Sort(suffixes)

	for i := 0; i < size; i++ {
		sa := size - len(suffixes[i].Value)
		mod := (sa + size - 1) % size
		r := str[mod]
		result[i].Value = string(r)

		if _, ok := rank[r]; ok {
			rank[r] = rank[r] + 1
		} else {
			rank[r] = 0
		}

		result[i].Rank = rank[r]
	}

	return result, nil
}

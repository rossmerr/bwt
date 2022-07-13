package bwt

import (
	"fmt"
	"sort"
	"strings"
)

const stx = "\002"
const ext = "\003"

func Bwt(str string) (string, error) {
	if strings.Contains(str, stx) || strings.Contains(str, ext) {
		err := fmt.Errorf("input string cannot contain STX and ETX characters")
		return "", err
	}

	str = stx + str + ext
	size := len(str)
	table := make([]string, size)
	for i := 0; i < size; i++ {
		table[i] = str[i:] + str[:i]
	}

	sort.Strings(table)

	lastBytes := make([]byte, size)
	for i := 0; i < size; i++ {
		lastBytes[i] = table[i][size-1]
	}
	return string(lastBytes), nil

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
		if strings.HasSuffix(row, ext) {
			return row[1 : size-1]
		}
	}
	return ""
}

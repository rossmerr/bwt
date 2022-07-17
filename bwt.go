package bwt

import (
	"context"
	"fmt"
	"strings"

	"github.com/rossmerr/graphblas"
	"github.com/rossmerr/graphblas/sort"
)

const stx = "\002"

func Bwt(str string) (string, error) {
	if strings.Contains(str, stx) {
		err := fmt.Errorf("input string cannot contain STX character")
		return "", err
	}

	str = stx + str
	size := len(str)
	matrix := graphblas.NewDenseMatrix[rune](size, size)

	for i := 0; i < size; i++ {
		tmp := str[i:] + str[:i]
		for j := 0; j < size; j++ {
			matrix.Set(i, j, rune(tmp[j]))
		}
	}

	sorted := sort.BubbleRow(context.Background(), matrix)
	last := sorted.ColumnsAt(size - 1)
	return graphblas.String(context.Background(), last), nil
}

func Ibwt(str string) string {
	size := len(str)
	matrix := graphblas.NewDenseMatrix[rune](size, size)
	for i := size - 1; i >= 0; i-- {
		for j := 0; j < size; j++ {
			tmp := str[j : j+1][0]
			matrix.Set(j, i, rune(tmp))
		}

		matrix = sort.BubbleRow(context.Background(), matrix).(*graphblas.DenseMatrix[rune])
	}

	first := matrix.RowsAt(0)
	result := graphblas.String(context.Background(), first)
	if strings.HasPrefix(result, stx) {
		return result[1:size]
	}

	return ""
}

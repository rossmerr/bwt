package suffixarray_test

import (
	"reflect"
	"testing"

	"github.com/rossmerr/bwt/suffixarray"
)

func TestSampleSuffixArray(t *testing.T) {

	tests := []struct {
		name        string
		arr         []int
		sa          suffixarray.Suffix
		compression int
		wantErr     bool
	}{
		{
			name: "abaaba",

			arr: []int{6, 5, 2, 3, 0, 4, 1},
			sa: func() suffixarray.Suffix {
				return suffixarray.NewSuffix[suffixarray.SuffixArray](7)

			}(),
			compression: 2,
		},
		{
			name: "abaaba",

			arr: []int{6, 5, 2, 3, 0, 4, 1},
			sa: func() suffixarray.Suffix {
				return suffixarray.NewSuffix[suffixarray.SuffixArray](7)

			}(),
			compression: 3,
		},
		{
			name: "abaaba",

			arr: []int{6, 5, 2, 3, 0, 4, 1},
			sa: func() suffixarray.Suffix {
				return suffixarray.NewSuffix[suffixarray.SuffixArray](7)

			}(),
			compression: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := suffixarray.NewSuffix[suffixarray.SampleSuffixArray](len(tt.arr), suffixarray.WithCompression(tt.compression))

			for i, v := range tt.arr {
				sa.Set(i, v)
				tt.sa.Set(i, v)
			}

			for i := 0; i < 7; i += tt.compression {
				r := tt.sa.Get(i)
				result := sa.Get(i)
				if !reflect.DeepEqual(r, result) {
					t.Errorf("BwtFirstLastSuffix() = %v, want %v", r, result)
				}
			}
		})
	}
}

func TestSampleSuffixArrayEnumerate(t *testing.T) {

	tests := []struct {
		name        string
		arr         []int
		sa          suffixarray.Suffix
		compression int
		wantErr     bool
	}{
		{
			name: "abaaba",

			arr: []int{6, 5, 2, 3, 0, 4, 1},
			sa: func() suffixarray.Suffix {
				return suffixarray.NewSuffix[suffixarray.SuffixArray](7)

			}(),
			compression: 2,
		},
		{
			name: "abaaba",

			arr: []int{6, 5, 2, 3, 0, 4, 1},
			sa: func() suffixarray.Suffix {
				return suffixarray.NewSuffix[suffixarray.SuffixArray](7)

			}(),
			compression: 3,
		},
		{
			name: "abaaba",

			arr: []int{6, 5, 2, 3, 0, 4, 1},
			sa: func() suffixarray.Suffix {
				return suffixarray.NewSuffix[suffixarray.SuffixArray](7)

			}(),
			compression: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := suffixarray.NewSuffix[suffixarray.SampleSuffixArray](len(tt.arr), suffixarray.WithCompression(tt.compression))
			for i, v := range tt.arr {
				sa.Set(i, v)
				tt.sa.Set(i, v)
			}

			iterator := sa.Enumerate()
			count := 0
			for iterator.HasNext() {
				i, _ := iterator.Next()

				r := tt.sa.Get(count)
				if !reflect.DeepEqual(r, i) {
					t.Errorf("BwtFirstLastSuffix.Enumerate(%v) = %v, want %v", count, r, i)
				}
				count += tt.compression
			}

		})
	}
}

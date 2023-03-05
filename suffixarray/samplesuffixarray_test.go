package suffixarray_test

import (
	"reflect"
	"testing"

	"github.com/rossmerr/bwt/suffixarray"
)

func TestSampleSuffixArray(t *testing.T) {

	tests := []struct {
		name       string
		arr        []int
		sa         suffixarray.Suffix
		sampleRate int
		wantErr    bool
	}{
		{
			name: "abaaba",

			arr: []int{6, 5, 2, 3, 0, 4, 1},
			sa: func() suffixarray.Suffix {
				return suffixarray.NewSuffixArray(7)

			}(),
			sampleRate: 1,
		},
		{
			name: "abaaba",

			arr: []int{6, 5, 2, 3, 0, 4, 1},
			sa: func() suffixarray.Suffix {
				return suffixarray.NewSuffixArray(7)

			}(),
			sampleRate: 2,
		},
		{
			name: "abaaba",

			arr: []int{6, 5, 2, 3, 0, 4, 1},
			sa: func() suffixarray.Suffix {
				return suffixarray.NewSuffixArray(7)

			}(),
			sampleRate: 3,
		},
		{
			name: "abaaba",

			arr: []int{6, 5, 2, 3, 0, 4, 1},
			sa: func() suffixarray.Suffix {
				return suffixarray.NewSuffixArray(7)

			}(),
			sampleRate: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := suffixarray.NewSampleSuffixArray(len(tt.arr), tt.sampleRate)

			for i, v := range tt.arr {
				sa.Set(i, v)
				tt.sa.Set(i, v)
			}

			for i := 0; i < 7; i += tt.sampleRate {
				r := tt.sa.Get(i)
				result := sa.Get(i)
				if !reflect.DeepEqual(r, result) {
					t.Errorf("SampleSuffixArray() = %v, want %v", r, result)
				}
			}
		})
	}
}

func TestSampleSuffixArrayEnumerate(t *testing.T) {

	tests := []struct {
		name       string
		arr        []int
		sa         suffixarray.Suffix
		sampleRate int
		wantErr    bool
	}{
		{
			name: "abaaba",

			arr: []int{6, 5, 2, 3, 0, 4, 1},
			sa: func() suffixarray.Suffix {
				return suffixarray.NewSuffixArray(7)

			}(),
			sampleRate: 2,
		},
		{
			name: "abaaba",

			arr: []int{6, 5, 2, 3, 0, 4, 1},
			sa: func() suffixarray.Suffix {
				return suffixarray.NewSuffixArray(7)

			}(),
			sampleRate: 3,
		},
		{
			name: "abaaba",

			arr: []int{6, 5, 2, 3, 0, 4, 1},
			sa: func() suffixarray.Suffix {
				return suffixarray.NewSuffixArray(7)

			}(),
			sampleRate: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := suffixarray.NewSampleSuffixArray(len(tt.arr), tt.sampleRate)
			for i, v := range tt.arr {
				sa.Set(i, v)
				tt.sa.Set(i, v)
			}

			iterator := sa.Enumerate()
			for iterator.HasNext() {
				result, i := iterator.Next()

				r := tt.sa.Get(i)
				if !reflect.DeepEqual(r, result) {
					t.Errorf("SampleSuffixArray.Enumerate(%v) = %v, want %v", i, r, result)
				}
			}

		})
	}
}

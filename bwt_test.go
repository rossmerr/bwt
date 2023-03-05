package bwt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rossmerr/bwt"
	"github.com/rossmerr/bwt/suffixarray"
)

func TestMatrx(t *testing.T) {

	tests := []struct {
		name    string
		str     string
		want    [][]rune
		wantErr bool
	}{
		{
			name: "abaaba",
			str:  "abaaba",
			want: [][]rune{
				[]rune("abaaba"),
				[]rune("aabaab"),
				[]rune("aabaab"),
				[]rune("abaaba"),
				[]rune("abaaba"),
				[]rune("baabaa"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bwt.Matrix(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Last() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)

			for i, row := range tt.want {
				if string(row) != string(got[i]) {
					t.Errorf("Last() = %v, want %v", string(got[i]), string(row))
				}
			}
		})
	}
}

func TestLast(t *testing.T) {

	tests := []struct {
		name    string
		str     string
		want    string
		wantErr bool
	}{
		{
			name: "abaaba",
			str:  "abaaba",
			want: "abbaaa",
		},
		{
			name: "banana",
			str:  "banana",
			want: "annbaa",
		},
		{
			name: "quick fox",
			str:  "The quick brown fox jumps over the lazy dog",
			want: "gkynxeserl i hhv otTu c uwd rfm ebp qjoooza",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bwt.Last(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Last() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != tt.want {
				t.Errorf("Last() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "banana",
			str:  "annbaa",
			want: "banana",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bwt.Reverse(tt.str); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstLast(t *testing.T) {

	tests := []struct {
		name    string
		str     string
		first   string
		last    string
		wantErr bool
	}{
		{
			name:  "abaaba",
			str:   "abaaba",
			first: "aaaabb",
			last:  "abbaaa",
		},
		{
			name:  "quick fox",
			str:   "The quick brown fox jumps over the lazy dog",
			first: "        Tabcdeeefghhijklmnoooopqrrstuuvwxyz",
			last:  "gkynxeserl i hhv otTu c uwd rfm ebp qjoooza",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			first, last, err := bwt.FirstLast(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstLast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(first) != tt.first {
				t.Errorf("FirstLast() = %v, want %v", string(first), tt.first)
			}
			if string(last) != tt.last {
				t.Errorf("FirstLast() = %v, want %v", string(last), tt.last)
			}
		})
	}
}

func TestFirstLastSample(t *testing.T) {

	tests := []struct {
		name       string
		str        string
		first      []rune
		last       []rune
		sa         suffixarray.Suffix
		sampleRate int
		wantErr    bool
	}{
		{
			name:       "abaaba",
			str:        "abaaba",
			first:      []rune("aaaabb"),
			last:       []rune("abbaaa"),
			sampleRate: 1,
			sa: func() suffixarray.Suffix {

				sa := suffixarray.NewSampleSuffixArray(7, 1)
				sa.Set(0, 6)
				sa.Set(1, 5)
				sa.Set(2, 2)
				sa.Set(3, 3)
				sa.Set(4, 0)
				sa.Set(5, 4)
				sa.Set(6, 1)
				return sa
			}(),
		},
		{
			name:       "abaaba",
			str:        "abaaba",
			first:      []rune("aaaabb"),
			last:       []rune("abbaaa"),
			sampleRate: 2,
			sa: func() suffixarray.Suffix {

				sa := suffixarray.NewSampleSuffixArray(7, 1)
				sa.Set(0, 6)
				sa.Set(1, 5)
				sa.Set(2, 2)
				sa.Set(3, 3)
				sa.Set(4, 0)
				sa.Set(5, 4)
				sa.Set(6, 1)
				return sa
			}(),
		},
		{
			name:       "abaaba",
			str:        "abaaba",
			first:      []rune("aaaabb"),
			last:       []rune("abbaaa"),
			sampleRate: 3,
			sa: func() suffixarray.Suffix {

				sa := suffixarray.NewSampleSuffixArray(7, 1)
				sa.Set(0, 6)
				sa.Set(1, 5)
				sa.Set(2, 2)
				sa.Set(3, 3)
				sa.Set(4, 0)
				sa.Set(5, 4)
				sa.Set(6, 1)
				return sa
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			first, last, sa, err := bwt.FirstLastSuffix(tt.str, bwt.WithSampleRate(tt.sampleRate))
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstLastSuffix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(first) != string(tt.first) {
				t.Errorf("FirstLastSuffix() = %v, want %v", first, tt.first)
			}
			if !reflect.DeepEqual(last, tt.last) {
				t.Errorf("FirstLastSuffix() = %v, want %v", last, tt.last)
			}
			for i := 0; i < 7; i += tt.sampleRate {
				r := tt.sa.Get(i)
				result := sa.Get(i)

				if !reflect.DeepEqual(r, result) {
					t.Errorf("FirstLastSuffix() = %v, want %v", r, result)
				}
			}
		})
	}
}

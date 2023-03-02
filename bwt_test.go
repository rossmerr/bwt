package bwt_test

import (
	"reflect"
	"testing"

	"github.com/rossmerr/bwt"
	"github.com/rossmerr/bwt/suffixarray"
)

func TestBwt(t *testing.T) {

	tests := []struct {
		name    string
		str     string
		want    string
		wantErr bool
	}{
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
			got, err := bwt.Bwt(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bwt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bwt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIbwt(t *testing.T) {
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
			if got := bwt.Ibwt(tt.str); got != tt.want {
				t.Errorf("Ibwt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBwtFirstLast(t *testing.T) {

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
			first, last, err := bwt.BwtFirstLast(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bwt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if first != tt.first {
				t.Errorf("Bwt() = %v, want %v", first, tt.first)
			}
			if last != tt.last {
				t.Errorf("Bwt() = %v, want %v", last, tt.last)
			}
		})
	}
}

func TestBwtFirstLastSuffix(t *testing.T) {

	tests := []struct {
		name    string
		str     string
		first   []rune
		last    []rune
		sa      suffixarray.Suffix
		wantErr bool
	}{
		{
			name:  "abaaba",
			str:   "abaaba",
			first: []rune("aaaabb"),
			last:  []rune("abbaaa"),
			sa: func() suffixarray.Suffix {
				sa := suffixarray.NewSuffix[suffixarray.SuffixArray](7)
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
			first, last, sa, err := bwt.BwtFirstLastSuffix[suffixarray.SuffixArray](tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bwt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(first) != string(tt.first) {
				t.Errorf("Bwt() = %v, want %v", first, tt.first)
			}
			if !reflect.DeepEqual(last, tt.last) {
				t.Errorf("BwtFirstLastSuffix() = %v, want %v", last, tt.last)
			}
			for i := 0; i < 7; i++ {
				r := tt.sa.Get(i)
				result := sa.Get(i)
				if !reflect.DeepEqual(r, result) {
					t.Errorf("BwtFirstLastSuffix() = %v, want %v", r, result)
				}
			}
		})
	}
}

func TestBwtFirstLastSampleSuffix(t *testing.T) {

	tests := []struct {
		name        string
		str         string
		first       []rune
		last        []rune
		sa          suffixarray.Suffix
		compression int
		wantErr     bool
	}{
		{
			name:        "abaaba",
			str:         "abaaba",
			first:       []rune("aaaabb"),
			last:        []rune("abbaaa"),
			compression: 2,
			sa: func() suffixarray.Suffix {

				sa := suffixarray.NewSuffix[suffixarray.SuffixArray](7)
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
			name:        "abaaba",
			str:         "abaaba",
			first:       []rune("aaaabb"),
			last:        []rune("abbaaa"),
			compression: 3,
			sa: func() suffixarray.Suffix {

				sa := suffixarray.NewSuffix[suffixarray.SuffixArray](7)
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
			first, last, sa, err := bwt.BwtFirstLastSuffix[suffixarray.SampleSuffixArray](tt.str, suffixarray.WithCompression(tt.compression))
			if (err != nil) != tt.wantErr {
				t.Errorf("Bwt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(first) != string(tt.first) {
				t.Errorf("Bwt() = %v, want %v", first, tt.first)
			}
			if !reflect.DeepEqual(last, tt.last) {
				t.Errorf("BwtFirstLastSuffix() = %v, want %v", last, tt.last)
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

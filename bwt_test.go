package bwt_test

import (
	"reflect"
	"testing"

	"github.com/rossmerr/bwt"
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

func TestBwtFirstLastSuffix(t *testing.T) {

	tests := []struct {
		name string
		str  string
		want string
		sa   bwt.Suffix
	}{
		{
			name: "banana",
			str:  "banana",
			want: "annbaa",
			sa: func() bwt.Suffix {
				sa := bwt.NewSuffix[bwt.SuffixArray]()
				sa.Append(6)
				sa.Append(5)
				sa.Append(3)
				sa.Append(1)
				sa.Append(0)
				sa.Append(4)
				sa.Append(2)
				return sa
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, last, sa, _ := bwt.BwtFirstLastSuffix[bwt.SuffixArray](tt.str)

			if !reflect.DeepEqual(last, tt.want) {
				t.Errorf("BwtFirstLastSuffix() = %v, want %v", last, tt.want)
			}
			if !reflect.DeepEqual(sa, tt.sa) {
				t.Errorf("BwtFirstLastSuffix() = %v, want %v", sa, tt.sa)
			}
		})
	}
}

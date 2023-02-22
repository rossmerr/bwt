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
		sa   []int
	}{
		{
			name: "banana",
			str:  "banana",
			want: "annbaa",
			sa:   []int{6, 5, 3, 1, 0, 4, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, sa, _ := bwt.BwtFirstLastSuffix(tt.str)

			if !reflect.DeepEqual(sa, tt.sa) {
				t.Errorf("BwtFirstLastSuffix() = %v, want %v", sa, tt.want)
			}
		})
	}
}

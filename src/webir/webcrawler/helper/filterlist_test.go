package helper

import (
	"reflect"
	"testing"
)

func TestFilterList(t *testing.T) {
	type args struct {
		filter func(string) bool
	}
	tests := []struct {
		name string
		args args
		want func([]string) []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterList(tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterList() = %v, want %v", got, tt.want)
			}
		})
	}
}

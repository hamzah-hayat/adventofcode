package main

import (
	"reflect"
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "13"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example")
	expected := ""

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestSplitLineIntoLists(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "basic",
			args: args{"[1,2,3,4]"},
			want: []string{"1", "2", "3", "4"},
		},
		{
			name: "One nested Left",
			args: args{"[[1],[2,3,4]]"},
			want: []string{"[1]", "[2,3,4]"},
		},
		{
			name: "One nested right",
			args: args{"[1,[2,[3,[4,[5,6,7]]]],8,9]"},
			want: []string{"1", "[2,[3,[4,[5,6,7]]]]", "8", "9"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitLineIntoLists(tt.args.line); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("SplitLineIntoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

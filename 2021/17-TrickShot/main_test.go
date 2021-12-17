package main

import (
	"reflect"
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "45"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example")
	expected := "112"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func Test_createTargetZone(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example grid test",
			args: args{input: "target area: x=20..30, y=-10..-5"},
			want: 66,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := createTargetZone(tt.args.input); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("createTargetZone() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "0"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example")
	expected := "0"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestCreateSnailPair(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want NumberPair
	}{
		{
			name: "basic",
			args: args{"[1,2]"},
			want: NumberPair{X: &NumberPair{value: 1}, Y: &NumberPair{value: 2}},
		},
		{
			name: "One nested Left",
			args: args{"[[1,2],3]"},
			want: NumberPair{X: &NumberPair{X: &NumberPair{value: 1}, Y: &NumberPair{value: 2}}, Y: &NumberPair{value: 3}},
		},
		{
			name: "One nested right",
			args: args{"[0,[6,7]]"},
			want: NumberPair{X: &NumberPair{value: 0}, Y: &NumberPair{X: &NumberPair{value: 6}, Y: &NumberPair{value: 7}}},
		},
		{
			name: "Both nested",
			args: args{"[[9,1],[1,9]]"},
			want: NumberPair{X: &NumberPair{X: &NumberPair{value: 9}, Y: &NumberPair{value: 1}}, Y: &NumberPair{X: &NumberPair{value: 1}, Y: &NumberPair{value: 9}}},
		},
		{
			name: "Complex small",
			args: args{"[[[0,7],4],[[7,8],[6,0]]]"},
			want: NumberPair{X: &NumberPair{X: &NumberPair{X: &NumberPair{value: 0}, Y: &NumberPair{value: 7}}, Y: &NumberPair{value: 4}}, Y: &NumberPair{X: &NumberPair{X: &NumberPair{value: 7}, Y: &NumberPair{value: 8}}, Y: &NumberPair{X: &NumberPair{value: 6}, Y: &NumberPair{value: 0}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateNumberPair(tt.args.input); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("CreateNumberPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExplodePairs(t *testing.T) {
	type args struct {
		numPair *NumberPair
		nested  int
	}
	tests := []struct {
		name string
		args args
		want *NumberPair
	}{
		{
			name: "basic Explode",
			args: args{CreateNumberPair("[[[[[9,8],1],2],3],4]"), 0},
			want: CreateNumberPair("[[[[0,9],2],3],4]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExplodePairs(tt.args.numPair, tt.args.nested); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExplodePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitPair(t *testing.T) {
	type args struct {
		numPair *NumberPair
	}
	tests := []struct {
		name string
		args args
		want *NumberPair
	}{
		{
			name: "Basic Split",
			args: args{&NumberPair{X: &NumberPair{value: 0}, Y: &NumberPair{value: 13}}},
			want: CreateNumberPair("[0,[6,7]]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitPair(tt.args.numPair); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScorePair(t *testing.T) {
	type args struct {
		numPair *NumberPair
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Basic Score - 1",
			args: args{CreateNumberPair("[[9,1],[1,9]]")},
			want: 129,
		},
		{
			name: "Basic Score - 2",
			args: args{CreateNumberPair("[[1,2],[[3,4],5]]")},
			want: 143,
		},
		{
			name: "Basic Score - 3",
			args: args{CreateNumberPair("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")},
			want: 1384,
		},
		{
			name: "Basic Score - 4",
			args: args{CreateNumberPair("[[[[1,1],[2,2]],[3,3]],[4,4]]")},
			want: 445,
		},
		{
			name: "Basic Score - 5",
			args: args{CreateNumberPair("[[[[3,0],[5,3]],[4,4]],[5,5]]")},
			want: 791,
		},
		{
			name: "Basic Score - 6",
			args: args{CreateNumberPair("[[[[5,0],[7,4]],[5,5]],[6,6]]")},
			want: 1137,
		},
		{
			name: "Basic Score - 7",
			args: args{CreateNumberPair("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")},
			want: 3488,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ScorePair(tt.args.numPair); got != tt.want {
				t.Errorf("ScorePair() = %v, want %v", got, tt.want)
			}
		})
	}
}

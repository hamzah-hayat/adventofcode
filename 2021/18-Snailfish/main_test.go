package main

import (
	"reflect"
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "445"

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
		{
			name: "Score example addition",
			args: args{CreateNumberPair("[[[[1,1],[2,2]],[3,3]],[4,4]]")},
			want: 445,
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

func TestExplodePairs(t *testing.T) {
	type args struct {
		numPair *NumberPair
	}
	tests := []struct {
		name  string
		args  args
		want  *NumberPair
		want1 bool
	}{
		{
			name:  "basic Explode - 1",
			args:  args{CreateNumberPair("[[[[[9,8],1],2],3],4]")},
			want:  CreateNumberPair("[[[[0,9],2],3],4]"),
			want1: true,
		},
		// {
		// 	name:  "basic Explode - 2",
		// 	args:  args{CreateNumberPair("[7,[6,[5,[4,[3,2]]]]]")},
		// 	want:  CreateNumberPair("[7,[6,[5,[7,0]]]]"),
		// 	want1: true,
		// },
		// {
		// 	name:  "basic Explode - 3",
		// 	args:  args{CreateNumberPair("[[6,[5,[4,[3,2]]]],1]")},
		// 	want:  CreateNumberPair("[[6,[5,[7,0]]],3]"),
		// 	want1: true,
		// },
		// {
		// 	name:  "basic Explode - 4",
		// 	args:  args{CreateNumberPair("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]")},
		// 	want:  CreateNumberPair("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"),
		// 	want1: true,
		// },
		// {
		// 	name:  "basic Explode - 5",
		// 	args:  args{CreateNumberPair("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]")},
		// 	want:  CreateNumberPair("[[3,[2,[8,0]]],[9,[5,[7,0]]]]"),
		// 	want1: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ExplodePairs(tt.args.numPair, 0, nil)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExplodePairs() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ExplodePairs() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSplitPair(t *testing.T) {
	type args struct {
		numPair *NumberPair
	}
	tests := []struct {
		name  string
		args  args
		want  *NumberPair
		want1 bool
	}{
		{
			name:  "Basic Split",
			args:  args{&NumberPair{X: &NumberPair{value: 0}, Y: &NumberPair{value: 13}}},
			want:  CreateNumberPair("[0,[6,7]]"),
			want1: true,
		},
		{
			name:  "Basic Split",
			args:  args{&NumberPair{X: &NumberPair{value: 0}, Y: &NumberPair{value: 12}}},
			want:  CreateNumberPair("[0,[6,6]]"),
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SplitPair(tt.args.numPair)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitPair() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SplitPair() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

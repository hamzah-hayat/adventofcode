package main

import (
	"reflect"
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "31"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example")
	expected := "54"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestParsePacket(t *testing.T) {
	type args struct {
		currentIndex int64
		inputBinary  string
	}
	tests := []struct {
		name string
		args args
		want Packet
	}{
		{
			name: "simple literal",
			args: args{
				currentIndex: 0,
				inputBinary:  convertHexToBinary("D2FE28"),
			},
			want: Packet{
				version:    6,
				id:         4,
				length:     21,
				lengthId:   0,
				value:      2021,
				subPackets: []Packet{},
			},
		},
		{
			name: "lenghtid 0 operator",
			args: args{
				currentIndex: 0,
				inputBinary:  convertHexToBinary("38006F45291200"),
			},
			want: Packet{
				version:  1,
				id:       6,
				length:   49,
				lengthId: 0,
				value:    0,
				subPackets: []Packet{{
					version:    6,
					id:         4,
					length:     11,
					lengthId:   0,
					value:      10,
					subPackets: []Packet{},
				}, {
					version:    2,
					id:         4,
					length:     16,
					lengthId:   0,
					value:      20,
					subPackets: []Packet{},
				}},
			},
		},
		{
			name: "lenghtid 1 operator",
			args: args{
				currentIndex: 0,
				inputBinary:  convertHexToBinary("EE00D40C823060"),
			},
			want: Packet{
				version:  7,
				id:       3,
				length:   51,
				lengthId: 1,
				value:    0,
				subPackets: []Packet{{
					version:    2,
					id:         4,
					length:     11,
					lengthId:   0,
					value:      1,
					subPackets: []Packet{},
				}, {
					version:    4,
					id:         4,
					length:     11,
					lengthId:   0,
					value:      2,
					subPackets: []Packet{},
				}, {
					version:    1,
					id:         4,
					length:     11,
					lengthId:   0,
					value:      3,
					subPackets: []Packet{},
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParsePacket(tt.args.currentIndex, tt.args.inputBinary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePacket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvaluatePacketRecursive(t *testing.T) {
	type args struct {
		packet Packet
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Sum packet",
			args: args{
				packet: ParsePacket(0, convertHexToBinary("C200B40A82")),
			},
			want: 3,
		}, {
			name: "Product packet",
			args: args{
				packet: ParsePacket(0, convertHexToBinary("04005AC33890")),
			},
			want: 54,
		}, {
			name: "Min packet",
			args: args{
				packet: ParsePacket(0, convertHexToBinary("880086C3E88112")),
			},
			want: 7,
		}, {
			name: "Max packet",
			args: args{
				packet: ParsePacket(0, convertHexToBinary("CE00C43D881120")),
			},
			want: 9,
		}, {
			name: "Less Than packet",
			args: args{
				packet: ParsePacket(0, convertHexToBinary("D8005AC2A8F0")),
			},
			want: 1,
		}, {
			name: "Greater Than packet",
			args: args{
				packet: ParsePacket(0, convertHexToBinary("F600BC2D8F")),
			},
			want: 0,
		}, {
			name: "Equal packet",
			args: args{
				packet: ParsePacket(0, convertHexToBinary("9C005AC2F8F0")),
			},
			want: 0,
		}, {
			name: "Calc packet",
			args: args{
				packet: ParsePacket(0, convertHexToBinary("9C0141080250320F1802104A08")),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvaluatePacketRecursive(tt.args.packet); got != tt.want {
				t.Errorf("EvaluatePacket() = %v, want %v", got, tt.want)
			}
		})
	}
}

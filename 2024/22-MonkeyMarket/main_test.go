package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "37327623"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestSilver_SecretGen(t *testing.T) {

	secrets := make(map[int]int)
	result := CalculateSecretAfterLoop(123, 1)
	secrets[result] += 1
	result = CalculateSecretAfterLoop(result, 1)
	secrets[result] += 1
	result = CalculateSecretAfterLoop(result, 1)
	secrets[result] += 1
	result = CalculateSecretAfterLoop(result, 1)
	secrets[result] += 1
	result = CalculateSecretAfterLoop(result, 1)
	secrets[result] += 1
	result = CalculateSecretAfterLoop(result, 1)
	secrets[result] += 1
	result = CalculateSecretAfterLoop(result, 1)
	secrets[result] += 1
	result = CalculateSecretAfterLoop(result, 1)
	secrets[result] += 1
	result = CalculateSecretAfterLoop(result, 1)
	secrets[result] += 1
	result = CalculateSecretAfterLoop(result, 1)
	secrets[result] += 1

	expected := make(map[int]int)
	expected[15887950] = 1
	expected[16495136] = 1
	expected[527345] = 1
	expected[704524] = 1
	expected[1553684] = 1
	expected[12683156] = 1
	expected[11100544] = 1
	expected[12249484] = 1
	expected[7753432] = 1
	expected[5908254] = 1

	for s, _ := range expected {
		_, exists := secrets[s]
		if !exists {
			t.Error("Did not find", s, "in secrets map")
		}
	}
}

func TestMix(t *testing.T) {
	value := 42 ^ 15
	expected := 37

	if value != expected {
		t.Error("Got ", value, " expected ", expected)
	}
}

func TestPrune(t *testing.T) {
	value := 100000000 % 16777216
	expected := 16113920

	if value != expected {
		t.Error("Got ", value, " expected ", expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example_gold", 2000)
	expected := "23"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold_Small(t *testing.T) {
	value := PartTwo("example_gold_small", 10)
	expected := "6"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

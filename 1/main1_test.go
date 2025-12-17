package main

import "testing"

func TestBuildCombinedString(t *testing.T) {
	result := buildCombinedString(
		42, 052, 0x2A,
		3.14,
		"Golang",
		true,
		complex64(1+2i),
	)

	expected := "4242423.14Golangtrue(1+2i)"

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestAddSaltToRunes(t *testing.T) {
	runes := []rune("abcdef")
	salt := []rune("go-2024")

	result := addSaltToRunes(runes, salt)
	expected := []rune("abcgo-2024def")

	if string(result) != string(expected) {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestHashRunesSHA256(t *testing.T) {
	input := []rune("test-go-2024")
	hash1 := hashRunesSHA256(input)
	hash2 := hashRunesSHA256(input)

	if hash1 != hash2 {
		t.Error("hashes should be equal for same input")
	}
}

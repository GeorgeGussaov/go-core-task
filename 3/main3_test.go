package main

import "testing"

func TestAddAndExists(t *testing.T) {
	m := NewStringIntMap()

	m.Add("one", 1)

	if !m.Exists("one") {
		t.Error("expected key 'one' to exist")
	}
}

func TestGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("answer", 42)

	value, ok := m.Get("answer")

	if !ok {
		t.Error("expected key to exist")
	}

	if value != 42 {
		t.Errorf("expected 42, got %d", value)
	}
}
func TestRemove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("temp", 100)

	m.Remove("temp")

	if m.Exists("temp") {
		t.Error("expected key to be removed")
	}
}
func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Add("b", 2)

	copied := m.Copy()

	if len(copied) != 2 {
		t.Errorf("expected 2 elements, got %d", len(copied))
	}

	if copied["a"] != 1 || copied["b"] != 2 {
		t.Error("copied map has incorrect values")
	}

	m.Remove("a")

	if _, ok := copied["a"]; !ok {
		t.Error("copy should not be affected by changes to original")
	}
}

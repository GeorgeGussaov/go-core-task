package main

import (
	"testing"
)

func Test(t *testing.T) {
	rnd := Random(15)
	if len(rnd) != 15 { //не знаю как еще проверить функцию с рандом значениями :/
		t.Error("wrong len")
	}
}

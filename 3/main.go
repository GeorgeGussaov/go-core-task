package main

import "maps"

type StringIntMap struct {
	storage map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		storage: make(map[string]int),
	}
}

func (sim *StringIntMap) Add(key string, value int) {
	sim.storage[key] = value
}

func (sim *StringIntMap) Remove(key string) {
	delete(sim.storage, key)
}

func (sim *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int)
	maps.Copy(newMap, sim.storage)
	return newMap
}

func (sim *StringIntMap) Exists(key string) bool {
	_, ok := sim.storage[key]
	return ok
}
func (sim *StringIntMap) Get(key string) (int, bool) {
	num, ok := sim.storage[key]
	return num, ok
}

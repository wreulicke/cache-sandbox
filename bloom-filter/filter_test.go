package main

import (
	"math"
	"testing"
)

type BloomFilter struct {
	bits      []bool
	numHashes uint64 // k
}

func New(capacity int, falsePositiveProbability float64) *BloomFilter {
	filterSize := int64(-(float64(capacity) * math.Log(falsePositiveProbability)) / (math.Log(2) * math.Log(2)))

	return &BloomFilter{
		bits:      make([]bool, filterSize),
		numHashes: uint64(float64(filterSize) / float64(capacity) * math.Log(2)),
	}
}

func (b *BloomFilter) Add(v int) {
	hash := uint64(v)
	var i uint64
	for ; i < b.numHashes; i++ {
		index := int(hash+uint64(i)*(hash>>32)) % len(b.bits)

		b.bits[index] = true
	}
}

func (b *BloomFilter) Contains(v int) bool {
	hashTop := uint64(v)
	hashBottom := hashTop >> 32
	var i uint64
	for ; i < b.numHashes; i++ {
		index := int(hashTop+uint64(i)*hashBottom) % len(b.bits)
		if !b.bits[index] {
			return false
		}
	}
	return true
}

func TestAddAndContains(t *testing.T) {
	b := New(10000, 0.1)
	b.Add(1)
	if !b.Contains(1) {
		t.Error("bloom filter does not contains 1")
	}
	if b.Contains(2) {
		t.Error("bloom filter does not contains 2")
	}
}

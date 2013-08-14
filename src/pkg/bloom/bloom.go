package bloom

import (
	"hash/fnv"
	"math/rand"
)

type Bloom uint64

const (
	k int = 3
	m int = 64
)

func NewBloom(str string) Bloom {
	var bloom Bloom = 0
	return bloom.computeBloom(str)
}

func (bloom Bloom) computeBloom(str string) Bloom {
	for i := 0; i < k; i++ {
		bloom |= 1 << generateIndex(str)
	}
	return bloom
}

func generateIndex(str string) uint {
	h := fnv.New32()
	b := []byte(str)
	
	h.Write(b)
	
	rand.Seed(int64(h.Sum32()))
	return uint(rand.Intn(m))
}

type Set struct {
	bloom Bloom
}

func NewSet() *Set {
	return &Set {
		0,
	}
}

func (set *Set) Add(str string) {
	set.bloom = set.bloom.computeBloom(str)
}

func (set *Set) Contains(str string) bool {
	bloom := NewBloom(str)
	for i := uint(0); i < 64; i++ {
		b1 := getBit(bloom, i)
		b2 := getBit(set.bloom, i)
		
		if b1 == 1 && b2 != 1 {
			return false
		}
	}
	return true
}

func getBit(bloom Bloom, index uint) int {
	return int((bloom >> index) & 1)
}


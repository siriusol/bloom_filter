package bloom_filter

import (
	"github/siriusol/bloom_filter/hash"
)

type BloomFilter struct {
	length int64
	filter []byte
	hashes []hash.HashHelper
}

type BloomFilterConfig struct {
	hashes []hash.HashHelper
}

/**
Create a bloom filter and use exact size.
*/
func NewBloomFilter(length int64) *BloomFilter {
	return &BloomFilter{
		length: length,
		filter: make([]byte, length),
	}
}

func (bf *BloomFilter) Init(items []int64, config BloomFilterConfig) {
	bf.hashes = config.hashes
	bf.Add(items...)
}

func (bf *BloomFilter) Add(items ...int64) {
	for _, item := range items {
		for _, helper := range bf.hashes {
			bf.filter[helper.Hash(item)] = 1
		}
	}
}

func (bf *BloomFilter) Contain(item int64) bool {
	for _, helper := range bf.hashes {
		if bf.filter[helper.Hash(item)] == 0 {
			return false
		}
	}
	return true
}

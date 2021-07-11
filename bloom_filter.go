package bloom_filter

import (
	"ther.cool/bloom_filter/hash"
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
	for _, item := range items {
		for _, helper := range bf.hashes {
			bf.filter[helper.Hash(item)] = 1
		}
	}
}

func (bf *BloomFilter) Contain(item int64) bool {
	isContained := true
	for _, helper := range bf.hashes {
		if bf.filter[helper.Hash(item)] == 0 {
			isContained = false
			break
		}
	}
	return isContained
}

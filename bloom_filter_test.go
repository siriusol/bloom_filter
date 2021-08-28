package bloom_filter

import (
	"fmt"
	"github/siriusol/bloom_filter/hash"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	bf := NewBloomFilter(30)
	bfConfig := BloomFilterConfig{
		hashes: []hash.HashHelper{
			hash.NewSimpleAddHashHelper(1, 30),
			hash.NewLinearHashHelper(3, 5, 30),
			hash.NewSquareHashHelper(30),
		},
	}
	testNums := []int64{1, 5, 9}
	bf.Init(testNums, bfConfig)
	bf.Add(5, 19)
	fmt.Println()
	for i := 0; i < 20; i++ {
		if bf.Contain(int64(i)) {
			fmt.Printf("%d\t", i)
		}
	}
}

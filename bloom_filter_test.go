package bloom_filter

import (
	"fmt"
	"testing"
	"ther.cool/bloom_filter/hash"
)

func TestBloomFilter(t *testing.T) {
	 bf := NewBloomFilter(10)
	 bfConfig := BloomFilterConfig{
	 	hashes: []hash.HashHelper{
			hash.NewSimpleAddHashHelper(1, 10),
			hash.NewSimpleAddHashHelper(3, 10),
			hash.NewSimpleAddHashHelper(7, 10),
		},
	 }
	 bf.Init([]int64{1,5,6}, bfConfig)
	for i := 0; i < 20; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()
	for i := 0; i < 20; i++ {
		fmt.Printf("%t\t", bf.Contain(int64(i)))
	}
}

package hash

import (
	"fmt"
	"testing"
)

func TestSimpleAddHashHelper(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()

	sa := NewSimpleAddHashHelper(1, 10)
	for i := 0; i < 20; i++ {
		fmt.Printf("%d\t", sa.Hash(int64(i)))
	}
	fmt.Println()

	sa2 := NewSimpleAddHashHelper(2, 10)
	for i := 0; i < 20; i++ {
		fmt.Printf("%d\t", sa2.Hash(int64(i)))
	}
}

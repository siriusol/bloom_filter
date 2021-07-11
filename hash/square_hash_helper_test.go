package hash

import (
	"fmt"
	"testing"
)

func TestSquareHashHelper(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()

	sa := NewSquareHashHelper(10)
	for i := 0; i < 20; i++ {
		fmt.Printf("%d\t", sa.Hash(int64(i)))
	}
	fmt.Println()
}

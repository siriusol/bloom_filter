package hash

import (
	"fmt"
	"testing"
)

func TestLinearHashHelper(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()

	sa := NewLinearHashHelper(2, 1, 10)
	for i := 0; i < 20; i++ {
		fmt.Printf("%d\t", sa.Hash(int64(i)))
	}
	fmt.Println()

	sa2 := NewLinearHashHelper(5, 0, 10)
	for i := 0; i < 20; i++ {
		fmt.Printf("%d\t", sa2.Hash(int64(i)))
	}
}

package bt_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/prathoss/gods/bt"
)

func TestBinaryTree_Add(t *testing.T) {
	tests := []struct {
		name  string
		items []int
	}{
		{
			name:  "simple",
			items: []int{3, 1, 5},
		},
		{
			name:  "full right",
			items: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "full left",
			items: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bt.NewBinaryTree(bt.CpFunc[int])
			for _, item := range tt.items {
				b.Add(item)
			}
			t.Logf("%v", b)
		})
	}
}

const benchmarkElements = 10_000

func getRandomNumber() int {
	return rand.Intn(benchmarkElements + 1)
}

func BenchmarkBinaryTree_Find(b *testing.B) {
	bt := bt.NewBinaryTree(bt.CpFunc[int])
	// load tree
	for i := 0; i < benchmarkElements; i++ {
		bt.Add(i + 1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ele := getRandomNumber()
		b.ResetTimer()
		bt.Find(ele)
	}
}

func BenchmarkSlice_Find(b *testing.B) {
	slice := make([]int, benchmarkElements)
	for i := 0; i < benchmarkElements; i++ {
		slice[i] = i + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ele := getRandomNumber()
		b.ResetTimer()
		for j := 0; j < len(slice); j++ {
			if slice[j] == ele {
				break
			}
		}
	}
}

func ExampleBinaryTree_Print() {
	binTree := bt.NewBinaryTree(bt.CpFunc[int])
	binTree.Add(3)
	binTree.Add(1)
	binTree.Add(5)
	binTree.Add(0)
	binTree.Add(2)
	binTree.Add(4)
	binTree.Add(6)
	render := binTree.Print(func(i int) string {
		return fmt.Sprint(i)
	})
	for i := 0; i < len(render); i++ {
		fmt.Print("[")
		for j := 0; j < len(render[i]); j++ {
			fmt.Print(render[i][j])
		}
		fmt.Print("]")
		fmt.Print("\n")
	}
	// Output:
	// [			3				]
	// [	1				5		]
	// [0		2		4		6	]
}

func ExampleBinaryTree_Print_Incomplete() {
	binTree := bt.NewBinaryTree(bt.CpFunc[int])
	binTree.Add(3)
	binTree.Add(1)
	binTree.Add(5)
	binTree.Add(4)
	binTree.Add(6)
	render := binTree.Print(func(i int) string {
		return fmt.Sprint(i)
	})
	for i := 0; i < len(render); i++ {
		fmt.Print("[")
		for j := 0; j < len(render[i]); j++ {
			fmt.Print(render[i][j])
		}
		fmt.Print("]")
		fmt.Print("\n")
	}
	// Output:
	// [			3				]
	// [	1				5		]
	// [				4		6	]
}

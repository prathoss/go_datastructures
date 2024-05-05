package bt

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestAdd(t *testing.T) {
	tc := []struct {
		name     string
		elements []int
	}{
		{
			name:     "rotate left",
			elements: []int{1, 2, 3},
		},
		{
			name:     "rotate right",
			elements: []int{3, 2, 1},
		},
		{
			name:     "rotate right left",
			elements: []int{1, 3, 2},
		},
		{
			name:     "rotate left right",
			elements: []int{3, 1, 2},
		},
	}

	for _, tc := range tc {
		t.Run(tc.name, func(t *testing.T) {
			tree := NewBinarySearchTree(CpFunc[int])

			for _, element := range tc.elements {
				tree.Add(element)
			}

			if tree.root.value != 2 {
				t.Fatalf("tree root value should be 2, but is: %v", tree.root.value)
			}
			if tree.root.height != 1 {
				t.Fatalf("tree root height should be 1, but is: %v", tree.root.height)
			}

			if tree.root.leftChild.value != 1 {
				t.Fatalf("tree left child value should be 1, but is: %v", tree.root.leftChild.value)
			}
			if tree.root.leftChild.height != 0 {
				t.Fatalf("tree left child height should be 0, but is: %v", tree.root.leftChild.height)
			}

			if tree.root.rightChild.value != 3 {
				t.Fatalf("tree right child value should be 3, but is: %v", tree.root.rightChild.value)
			}
			if tree.root.rightChild.height != 0 {
				t.Fatalf("tree right child height should be 0, but is: %v", tree.root.rightChild.height)
			}
		})
	}
}

const benchmarkElements = 10_000

func getRandomNumber() int {
	return rand.IntN(benchmarkElements + 1)
}

func TestAddSpeed(t *testing.T) {
	tree := NewBinarySearchTree(CpFunc[int])
	// load tree
	for i := range benchmarkElements {
		tree.Add(i + 1)
		// perct := float32(i) / benchmarkElements * 100
		// if int(perct)%5 == 0 {
		// 	t.Logf("loading progress: %v%%", perct)
		// }
	}

	for range benchmarkElements {
		tree.Find(getRandomNumber())
	}
}

func BenchmarkBinarySearchTree_Find(b *testing.B) {
	tree := NewBinarySearchTree(CpFunc[int])
	// load tree
	for i := range benchmarkElements {
		tree.Add(i)
	}
	elements := make([]int, 0, b.N)
	for range b.N {
		elements = append(elements, getRandomNumber())
	}

	b.ResetTimer()
	for i := range b.N {
		tree.Find(elements[i])
	}
}

func BenchmarkSlice_BinarySearch(b *testing.B) {
	s := make([]int, 0, benchmarkElements)
	for i := range benchmarkElements {
		s = append(s, i)
	}
	searchElements := make([]int, 0, b.N)
	for range b.N {
		searchElements = append(searchElements, getRandomNumber())
	}

	for i := range b.N {
		slices.BinarySearch(s, searchElements[i])
	}
}

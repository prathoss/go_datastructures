# Go datastructures

⚠️ This library is meant for playing with datastructures and golang generics,
it is not supposed to be used in any serious application.

List of packages:
- list (linked list)
- bt
  - [binary tree](./bt/bt.go)
  - [avl tree](./bt/bst.go)

## Benchmarks

Benchmarking is run on 0..10_000 elements. Will be always in order so that
slice binary search could be executed.

Slice binary search is for reference only, it is imported from `std`.

```
BenchmarkBinarySearchTree_Find-10       12373402                90.32 ns/op            0 B/op          0 allocs/op
BenchmarkSlice_BinarySearch-10          18345468                63.92 ns/op            8 B/op          0 allocs/op
BenchmarkBinaryTree_Find-10                26326             45410 ns/op               0 B/op          0 allocs/op
BenchmarkSlice_Find-10                    750846              1571 ns/op               0 B/op          0 allocs/op
```

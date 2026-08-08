package topkfrequentelements

import (
	"container/heap"
	"slices"
)

type kv struct {
	Key   int
	Value int
}

type kvHeap []kv

func (h kvHeap) Len() int {
	return len(h)
}

func (h kvHeap) Less(i, j int) bool {

	return h[i].Value < h[j].Value
}

func (h kvHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *kvHeap) Push(x any) {
	*h = append(*h, x.(kv))
}

func (h *kvHeap) Pop() any {
	i := len(*h) - 1
	v := (*h)[i]
	*h = (*h)[:i]
	return v
}

func topKFrequent(nums []int, k int) []int {

	result := make([]int, k)
	h := make(kvHeap, 0, len(nums))
	counts := make(map[int]int)
	heap.Init(&h)

	for _, v := range nums {
		counts[v]++
	}

	for i, j := range counts {
		heap.Push(&h, kv{Key: i, Value: j})
		if h.Len() > k {
			heap.Pop(&h)
		}
	}

	for i := 0; i < k; i++ {
		result[i] = h[i].Key
	}
	return result

}

func topKFrequentSort(nums []int, k int) []int {

	result := make([]int, k)
	counts := make(map[int]int)

	for _, v := range nums {
		counts[v]++
	}

	var list []kv
	for num, v := range counts {
		list = append(list, kv{num, v})
	}
	slices.SortFunc(list, func(a, b kv) int {
		return b.Value - a.Value
	})

	for i := 0; i < k; i++ {
		result[i] = list[i].Key
	}
	return result

}

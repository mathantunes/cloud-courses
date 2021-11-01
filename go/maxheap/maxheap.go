package maxheap

type MaxHeapArray []int

type MaxHeap struct {
	array MaxHeapArray
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Extract() int {
	max := h.array[0]
	lastIdx := len(h.array) - 1
	h.array[0] = h.array[lastIdx]
	h.array = h.array[:lastIdx]

	h.heapifyDown(0)

	return max
}

func (h *MaxHeap) heapifyUp(index int) {
	for h.array[h.array.parentIndex(index)] < h.array[index] {
		h.swap(h.array.parentIndex(index), index)
		index = h.array.parentIndex(index)
	}
}

func (h *MaxHeap) heapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := h.array.leftChildIndex(index), h.array.rightChildIndex(index)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex || h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = h.array.leftChildIndex(index), h.array.rightChildIndex(index)
		} else {
			return
		}
	}
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func (array *MaxHeapArray) parentIndex(i int) int {
	return (i - 1) / 2
}

func (array *MaxHeapArray) leftChildIndex(i int) int {
	return 2*i + 1
}

func (array *MaxHeapArray) rightChildIndex(i int) int {
	return 2*i + 2
}

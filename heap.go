package tree

type (
	Heap[T Ordered] struct {
		heap []T
	}
	Ordinary interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
	}
	Ordered interface {
		Ordinary
	}
)

func NewHeap[T Ordered](init []T) *Heap[T] {
	var tree = Heap[T]{heap: init}
	if tree.Len() > 0 {
		tree.makeBalance()
	}
	return &tree
}

func (t *Heap[T]) makeBalance() {
	var n = 0
	for idxOfRightChild(n) < t.Len() {
		n = idxOfRightChild(n)
	}
	for ; n >= 0; n-- {
		t.bDown(n)
	}
}

func (t *Heap[T]) Len() int {
	return len(t.heap)
}

func (t *Heap[T]) Put(nodes ...T) {
	for _, node := range nodes {
		heapLen := t.Len()
		t.heap = append(t.heap, node)
		t.bUp(heapLen)
	}
}

func (t *Heap[T]) PopMax() (val T, ok bool) {
	heapLen := t.Len()
	if heapLen == 0 {
		return
	}

	val = t.heap[0]
	t.heap[0] = t.heap[heapLen-1]
	t.heap = t.heap[:heapLen-1]
	t.bDown(0)
	return val, true
}

func (t *Heap[T]) bUp(currentIdx int) {
	for currentIdx > 0 {
		parentIdx := idxOfParent(currentIdx)
		if t.heap[parentIdx] < t.heap[currentIdx] {
			t.swap(currentIdx, parentIdx)
			currentIdx = parentIdx
			continue
		}
		return
	}
}

func (t *Heap[T]) bDown(currentIdx int) {
	for heapLen := t.Len(); currentIdx < heapLen; {
		var (
			leftChildIdx              = idxOfLeftChild(currentIdx)
			rightChildIdx             = idxOfRightChild(currentIdx)
			isLeftChildExists         = leftChildIdx < heapLen
			isRightChildExists        = rightChildIdx < heapLen
			isLeftGreaterThanCurrent  = isLeftChildExists && t.heap[currentIdx] < t.heap[leftChildIdx]
			isRightGreaterThanCurrent = isRightChildExists && t.heap[currentIdx] < t.heap[rightChildIdx]
			isRightGreaterThanLeft    = isRightChildExists && t.heap[leftChildIdx] < t.heap[rightChildIdx]
		)
		switch {

		case isLeftGreaterThanCurrent && !isRightGreaterThanLeft:
			t.swap(currentIdx, leftChildIdx)
			currentIdx = leftChildIdx
			continue

		case isRightGreaterThanCurrent:
			t.swap(currentIdx, rightChildIdx)
			currentIdx = rightChildIdx
			continue
		}
		return
	}
}

func (t *Heap[T]) swap(a, b int) {
	t.heap[a], t.heap[b] = t.heap[b], t.heap[a]
}

func idxOfLeftChild(idx int) int {
	return idx*2 + 1
}

func idxOfRightChild(idx int) int {
	return idx*2 + 2
}

func idxOfParent(idx int) int {
	return (idx - 1) / 2
}

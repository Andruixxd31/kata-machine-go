package heap

type Heap []int

func (h Heap) Insert(item int) Heap {
	h = append(h, item)
	h.bubbleUp(len(h) - 1)
	return h
}

func (h Heap) Delete() Heap {
	if len(h) == 0 {
		return h
	}

	h[0] = h[len(h)-1]
	h = h[:len(h)-1]

	h.bubbleDown(0)

	return h
}

func (h Heap) bubbleDown(idx int) Heap {
	n := len(h)
	left := 2*idx + 1
	right := 2*idx + 2
	smallest := idx

	if left < n && h[left] < h[smallest] {
		smallest = left
	}

	if right < n && h[right] < h[smallest] {
		smallest = right
	}

	if h[idx] < h[smallest] {
		return h
	}

	if smallest != idx {
		h.swap(smallest, idx)
		return h.bubbleDown(smallest)
	}

	return h
}

func (h Heap) bubbleUp(idx int) Heap {
	parent := (idx - 1) / 2

	if h[idx] >= h[parent] {
		return h
	}

	h.swap(idx, parent)
	h.bubbleUp(parent)
	return h
}

func (h Heap) swap(child, parent int) {
	h[child], h[parent] = h[parent], h[child]
}

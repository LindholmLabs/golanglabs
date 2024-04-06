package dsa

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) EnqueueSlice(items []int) {
	q.items = append(q.items, items...)
}

func (q *Queue) Dequeue() (item int, flag bool) {
	if q.IsEmpty() {
		return -1, false
	} else {
		head := q.items[0]
		q.items = q.items[1:]
		return head, true
	}
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

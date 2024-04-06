package dsa

type Queue struct {
	items []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) EnqueueSlice(items []interface{}) {
	for _, item := range items {
		q.Enqueue(item)
	}
}

func (q *Queue) Dequeue() (item interface{}, flag bool) {
	if q.IsEmpty() {
		return nil, false
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

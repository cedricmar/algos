package ds

type Queue[T any] []T

func (q *Queue[T]) Enqueue(e ...T) {
	*q = append(*q, e...)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.Length() == 0 {
		var v T
		return v, false
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v, true
}

func (q Queue[T]) Length() int {
	return len(q)
}

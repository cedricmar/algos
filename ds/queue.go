package ds

type Queuer interface {
	Enqueue(e ...int)
	Dequeue() int
	Length() int
}

type Queue []int

func (q *Queue) Enqueue(e ...int) {
	*q = append(*q, e...)
}

func (q *Queue) Dequeue() int {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

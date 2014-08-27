package list

type Node struct {
    next *Node
    value int
}

type List struct {
    head   Node
    count  int
}

func New() *List {
    return new(List)
}

func (l *List) Add(v int) {
}

func (l *List) Remove(v int) {
}

func (l *List) Contains(v int) bool {
    return false
}

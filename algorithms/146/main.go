type DLinkedNode struct {
        key   int
        value int
        next  *DLinkedNode
        prev  *DLinkedNode
}

type LRUCache struct {
        capacity int
        nodeMap  map[int]*DLinkedNode
        head     *DLinkedNode
        tail     *DLinkedNode
}

func Constructor(capacity int) LRUCache {
        head := &DLinkedNode{}
        tail := &DLinkedNode{}
        head.prev, tail.next = tail, head
        return LRUCache{
                capacity: capacity,
                nodeMap:  map[int]*DLinkedNode{},
                head:     head,
                tail:     tail,
        }
}

func (this *LRUCache) Get(key int) int {
        if node, ok := this.nodeMap[key]; ok && node != nil {
                this.moveToTop(node)
                return node.value
        } else {
                return -1
        }
}

func (this *LRUCache) deleteTail() {
        this.tail.next = this.tail.next.next
        this.tail.next.prev = this.tail
}

func (this *LRUCache) moveToTop(node *DLinkedNode) {
        if node.next != nil && node.prev != nil {
                node.next.prev, node.prev.next = node.prev, node.next
        }
        node.next, node.prev = this.head, this.head.prev
        this.head.prev.next, this.head.prev = node, node
}

func (this *LRUCache) Put(key int, value int) {
        node, ok := this.nodeMap[key]
        if ok {
            node.value = value
        } else {
                node = &DLinkedNode{
                        key:   key,
                        value: value,
                }
        } 
        this.moveToTop(node)
        this.nodeMap[key] = node
        if len(this.nodeMap) > this.capacity {
                delete(this.nodeMap, this.tail.next.key)
                this.deleteTail()
        }
}

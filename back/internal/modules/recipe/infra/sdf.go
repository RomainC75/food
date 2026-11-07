package recipe_infra

import "fmt"

type LRUCache struct {
	Cap           int
	Len           int
	LastNodeIndex int
	Head          *Node
	End           *Node
}

type Node struct {
	Value int
	Key   int
	Next  *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	fmt.Printf("\n-> GET %d", key)
	this.Display()
	if this.Head == nil {
		return -1
	}
	current := this.Head
	for {
		if current.Key == key {
			return current.Value
		}
		if current.Next == nil {
			break
		}
		current = current.Next
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	newNode := &Node{
		Key:   key,
		Value: value,
	}
	if this.Head == nil {
		this.Head = newNode
		this.End = newNode
		this.Len++
	} else if this.Len < this.Cap {
		this.End.Next = newNode
		this.End = newNode
		this.Len++
	} else {
		this.InsertInOldest(key, value)
	}
	fmt.Println("len: ", this.Len)
	this.Display()
}

func (this *LRUCache) Increment() {
	if this.LastNodeIndex == this.Cap-1 {
		this.LastNodeIndex = 0
		return
	}
	this.LastNodeIndex++
}

func (this *LRUCache) InsertInOldest(key int, value int) {
	current := this.Head
	i := 0
	for {

		if i == this.LastNodeIndex {
			current.Key = key
			current.Value = value
			break
		}
		if current.Next == nil {
			break
		}
		i++
		current = current.Next
	}
	this.Increment()
	fmt.Println("-")
}

func (this *LRUCache) Display() {
	if this.Head == nil {
		fmt.Println("[]")
	}
	current := this.Head
	i := 0
	for {
		fmt.Printf(" - [%d,%d]\n", current.Key, current.Value)

		if current.Next == nil {
			break
		}
		i++
		current = current.Next
	}
}

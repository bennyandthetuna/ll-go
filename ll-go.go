package llgo

//linked list implemented for a node container that holds string

type Node struct {
	value string
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) findEndPointer() *Node {
	if ll.head == nil {
		return ll.head
	} else {
		current := ll.head

		for current != nil {
			current = current.next
		}
		//and end of this loop, current should be the nil pointer at the end of the linked list
		return current
	}
}

func (ll *LinkedList) Add(value string) {
	newNode := Node{
		value: value,
		next:  nil,
	}

	endpointer := ll.findEndPointer()

	if endpointer == nil { //may as well double check that we actually have the ending nil pointer
		endpointer = &newNode
	}
}

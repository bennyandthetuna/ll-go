package llgo

import "fmt"

//linked list implemented for a node struct that holds string

type Node struct {
	Value string
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

// loops through entirety of linked list to find the nil pointer at the end of the list
func (ll *LinkedList) findEndItem() *Node {
	if ll.Head == nil {
		fmt.Print("head of list is empty!\n")
		return ll.Head
	} else {
		current := ll.Head

		for current.Next != nil {
			current = current.Next
		}
		//and end of this loop, current should be the nil pointer at the end of the linked list
		return current
	}
}

// find string value within linked list and return position indexed from 0, returns -1 if value not found
func (ll *LinkedList) FindValue(value string) int {

	current := ll.Head
	counter := 0
	for current != nil {
		if current.Value == value {
			return counter
		} else {
			counter++
			current = current.Next
		}
	}
	//if you exit the loop without returning the counter, then return not found.
	return -1
}

// add string value to the end of the linked list
func (ll *LinkedList) AddAtEnd(value string) {
	newNode := &Node{
		Value: value,
	}
	if ll.Head == nil {
		ll.Head = newNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}

}

// deletes the node at the end of a linked list
// returns true if operation was successful
// returns false if operations was unsuccessful
func (ll *LinkedList) DeleteAtEnd() bool {
	if ll.Head == nil {
		return false //no nodes in linked list, operation unsuccessful
	} else {
		current := ll.Head

		for current.Next.Next != nil { //checking not the current next pointer in node, but peaking into the one ahead
			//of current.next.next is nil, then that is  the last node in the linked list
			current = current.Next
		}
		current.Next = nil // by setting to nil, I'm removing all references to the last node, marking it for gc

		return true

	}
}

func (ll *LinkedList) PrintList() {
	current := ll.Head
	fmt.Print("Start of linked list\n")
	for current != nil {
		fmt.Printf("%s\n", current.Value)
		current = current.Next
	}
	fmt.Print("End of linked list\n")
}

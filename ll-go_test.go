package llgo

import (
	"fmt"
	"testing"
)

func TestFindEndPointer(t *testing.T) {
	value := "testing"
	value1 := "also-testing"

	myLinkedList := &LinkedList{}

	myLinkedList.AddAtEnd(value)
	myLinkedList.AddAtEnd(value1)

	ending := myLinkedList.findEndItem()

	if ending.Next != nil {
		fmt.Print("Did not correctly find  the end of the linked list - meaning currentNode.Next is not nil")
	}

}

func TestAddAtEnd(t *testing.T) {
	value := "testing"
	value1 := "also-testing"

	myLinkedList := &LinkedList{}

	myLinkedList.AddAtEnd(value)
	myLinkedList.AddAtEnd(value1)

	if myLinkedList.Head.Value != value || myLinkedList.Head.Next.Value != value1 {
		fmt.Print("AddAtEnd did not correctly put the values in place")
	}
}

func TestDeleteAtEnd(t *testing.T) {
	value := "testing"
	value1 := "also-testing"

	myLinkedList := &LinkedList{}

	myLinkedList.AddAtEnd(value)
	myLinkedList.AddAtEnd(value1)

	result := myLinkedList.DeleteAtEnd()

	if result == false {
		t.Error("node failed to delete")
	}

	if myLinkedList.Head.Value == value && myLinkedList.Head.Next == nil {
		t.Log("Success")

	} else {
		t.Error("AddAtEnd or DeleteAtEnd did not correctly run")
	}
}

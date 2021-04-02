package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func (n Node) setData(d int) {
	n.data = d
}

func (n Node) setNext(next *Node) {
	n.next = next
}

func (n Node) getNext() *Node {
	return n.next
}

func (n Node) getData() int {
	return n.data
}

func addNode(data int, head *Node) {

	if head.next == nil {
		head.next = new(Node)
		head.next.data = data
	} else {

		var temp *Node = head
		for temp.next != nil {
			temp = temp.getNext()
		}

		//on last node now append new node to iyt

		temp.next = new(Node)
		temp.next.data = data
	}

}

func insertAt(val int, position int, head *Node) {

	var parentptr *Node = head
	childptr := head.getNext()

	for i := 0; i < position-1 && parentptr != nil && childptr != nil; i++ {
		parentptr = childptr
		childptr = childptr.getNext()

	}

	parentptr.next = &Node{
		data: val,
		next: childptr,
	}

}

func deleteNode(val int, head *Node) {

	if head == nil {
		return
	}

	var parentptr *Node = head
	childptr := head.getNext()
	//fmt.Println("chile ptr", childptr.getData())
	for childptr != nil {

		if childptr.getData() == val {
			break
		} else {

			parentptr = childptr
			childptr = childptr.getNext()

		}
	}

	parentptr.next = childptr.next
	childptr.next = nil

}

func printList(head *Node) {
	if head == nil || head.getNext() == nil {
		fmt.Print("Link list is empty")
	} else {
		head = head.getNext()
		for head != nil {
			print(head.getData(), "-->")
			head = head.getNext()
		}
	}
}

func main() {

	var header *Node = &Node{
		data: 0,
		next: nil,
	}
	addNode(10, header)
	addNode(20, header)
	addNode(30, header)
	addNode(40, header)

	printList(header)
	fmt.Println("")
	deleteNode(20, header)

	deleteNode(10, header)

	deleteNode(30, header)

	deleteNode(40, header)

	printList(header)

	addNode(10, header)
	addNode(20, header)
	addNode(30, header)
	addNode(40, header)

	//deleteNode(40, header)
	fmt.Println("")

	printList(header)
	fmt.Println("")

	insertAt(22, 2, header)

	printList(header)

	insertAt(100, 100, header)
	insertAt(101, 101, header)

	fmt.Println("")
	printList(header)
}

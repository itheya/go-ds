package main

import "fmt"

type datanode struct {
	data string
	next *datanode
}

type linkeddata struct {
	head       *datanode
	lenofnodes int64
}

func (ld *linkeddata) insertNodeafter(afterData string, newData string) bool {
	pdataNode := datanode{
		data: newData,
		next: nil,
	}
	head := ld.head
	for {
		if ld == nil {
			fmt.Println(" no list List is empty inserting at the first")
			return false
		} else if head == nil {
			ld.head = &pdataNode
			break
		} else if head.data == afterData {
			pdataNode.next = head.next
			head.next = &pdataNode
			ld.lenofnodes++
			break
		} else {
			head = head.next
		}
	}

	//	ld.head = head
	return true
}

func (ld *linkeddata) CreateNewNode(pData string) {
	pdatanode := datanode{
		data: pData,
		next: nil,
	}
	//fmt.Println("Creating a new node")
	//var head *datanode
	head := ld.head

	if ld.head == nil {
		ld.head = &pdatanode
		//fmt.Println("Created  a first node", head.data)
	} else {
		for {
			if head.next == nil {
				head.next = &pdatanode
				break
			} else {
				head = head.next
			}
		}
	}
	//ld.head = head
	ld.lenofnodes++
}

func (ld linkeddata) printlist() {
	head := ld.head
	fmt.Println("Total number of the node in the linked list is", ld.lenofnodes)
	for {
		if head == nil {
			break
		} else {
			fmt.Println("->", head.data)
			head = head.next
		}
	}
}

func main() {
	mld := &linkeddata{
		head:       nil,
		lenofnodes: 0,
	}
	mld.CreateNewNode("a")
	mld.CreateNewNode("e")
	mld.CreateNewNode("g")
	mld.CreateNewNode("z")
	mld.insertNodeafter("a", "c")
	mld.printlist()
}

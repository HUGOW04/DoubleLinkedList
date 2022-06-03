package main

import "fmt"

type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
		linkedList.headNode.previousNode = node
	}
	linkedList.headNode = node
}

func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

func (linkedList *LinkedList) AddAfter(nodeProerty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var nodeWith *Node

	nodeWith = linkedList.NodeWithValue(nodeProerty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
	}
}

func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var lastNode *Node

	lastNode = linkedList.LastNode()

	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func (linkedList *LinkedList) NodeBetweenValues(firstProperty int, secoundProperty int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProperty && node.nextNode.property == secoundProperty {
				nodeWith = node
				break
			}
		}
	}
	return nodeWith
}

func main() {
	var linkedList LinkedList

	linkedList = LinkedList{}

	linkedList.AddToHead(1)
	linkedList.AddToHead(2)
	linkedList.AddToHead(3)
	linkedList.AddToHead(4)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 8)
	fmt.Println(linkedList.headNode.property)

	var node *Node
	node = linkedList.NodeBetweenValues(1, 5)
	fmt.Println(node.property)
}

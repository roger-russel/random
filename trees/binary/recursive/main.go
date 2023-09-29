package main

import (
	"errors"
	"fmt"
)

func main() {
	// bt := &BinaryTreeImpl{}
}

func NewBinaryTree() *BinaryTreeImpl {
	return &BinaryTreeImpl{}
}

type BinaryTree interface {
	Add(index int, data any)
	Remove(index int)
	Search(index int) (data any, err error)
}

type BinaryTreeNodesI interface {
	Add(Node BinaryTreeNodes, index int, data any)
}

var _ BinaryTree = (*BinaryTreeImpl)(nil)

type BinaryTreeNodes struct {
	TreeLeft *BinaryTreeNodes
	TreeRight *BinaryTreeNodes
	Index int
	Data interface{}
}

type BinaryTreeImpl struct {
	nodes *BinaryTreeNodes
}

func (bt *BinaryTreeNodes) Add(Node *BinaryTreeNodes, index int, data any) {
	switch {
		case Node.Index < index:
			if Node.TreeRight == nil {
				Node.TreeRight = &BinaryTreeNodes{
					Index: index,
					Data: data,
				}
				return
			}
			bt.Add(Node.TreeRight, index, data)			
		case Node.Index > index:
			if Node.TreeLeft == nil {
				Node.TreeLeft = &BinaryTreeNodes{
					Index: index,
					Data: data,
				}
				return
			}
			bt.Add(Node.TreeLeft, index, data)
	}
}

func (bt *BinaryTreeImpl) Add(index int, data any) {
	if bt.nodes == nil {
		bt.nodes = &BinaryTreeNodes{
			Index: index,
			Data: data,
		}
		return
	}
	bt.nodes.Add(bt.nodes, index, data)
}

func (bt *BinaryTreeImpl) Remove(index int) {}

var ErrIndexNotFound = errors.New("index was not found")

func (bt *BinaryTreeImpl) Search(index int) (data any, err error) {
	node := bt.nodes
	for {
		switch{
		case node == nil:
			return nil, fmt.Errorf("%w: %v", ErrIndexNotFound, index)
		case node.Index == index:
			return node.Data, nil
		case node.Index < index:
			node = node.TreeRight
		case node.Index > index:
			node = node.TreeLeft
		}
	}
}

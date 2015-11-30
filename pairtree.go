package main

import(

)

type PairTreeNode struct {
	value string
	parent *PairTreeNode
	left *PairTreeNode
	right *PairTreeNode
	length int
}

func NewPairTreeNode(e string) *PairTreeNode {
	return &PairTreeNode{value: e}
}

func (this *PairTreeNode) GetValue() string {
	if this == nil {
		return ""
	}
	return this.value
}

func (this *PairTreeNode) SetValue(e string) {
	this.value = e
}

func (this *PairTreeNode) HasLeft() bool {
	return this.left != nil
}

func (this *PairTreeNode) GetLeft() *PairTreeNode {
	if !this.HasLeft(){
		return nil
	}
	return this.left
}

func (this *PairTreeNode) SetLeft(l *PairTreeNode){
	this.left = l
}

func (this *PairTreeNode) HasRight() bool {
	return this.right != nil
}

func (this *PairTreeNode) GetRight() *PairTreeNode {
	if !this.HasRight(){
		return nil
	}
	return this.right
}

func (this *PairTreeNode) SetRight(r *PairTreeNode){
	this.right = r
}

func (this *PairTreeNode) Car() *PairTreeNode {
	return this.GetRight()
}

func (this *PairTreeNode) Cdr() *PairTreeNode {
	return this.GetLeft()
}

func (this *PairTreeNode) GetLength() int {
	this.SetLength()
	return this.length
}

func (this *PairTreeNode) SetLength() {
	var l int = 1
	var t *PairTreeNode = this
	for t.HasLeft() {
		l = l + 1
		t = t.left
	}
	this.length = l
}
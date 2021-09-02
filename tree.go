package main

type node struct {
	data int
	left *node
	right *node
}


func NewNode(data int) *node {
	return &node{data: data}
}

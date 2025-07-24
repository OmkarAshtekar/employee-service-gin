package model

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type Node struct {
	Emp  *Employee
	Prev *Node
	Next *Node
}

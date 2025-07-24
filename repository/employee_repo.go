package repository

import "employee-service-gin/model"

var head *model.Node
var tail *model.Node

func ClearRepo() {
	head = nil
	tail = nil
}

func AddEmployeeRepo(emp *model.Employee) {
	newNode := &model.Node{Emp: emp}
	if tail == nil {
		head = newNode
		tail = newNode
	} else {
		tail.Next = newNode
		newNode.Prev = tail
		tail = newNode
	}
}

func DeleteEmployeeRepo(id int) bool {
	current := head
	for current != nil {
		if current.Emp.ID == id {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				head = current.Next
			}
			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				tail = current.Prev
			}
			return true
		}
		current = current.Next
	}
	return false
}

func DisplayEmployeesRepo() []*model.Employee {
	var list []*model.Employee
	current := head
	for current != nil {
		list = append(list, current.Emp)
		current = current.Next
	}
	return list
}

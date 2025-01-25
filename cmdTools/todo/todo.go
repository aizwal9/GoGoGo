package main

import (
	"time"
)

type TodoItem struct{
	ID int
	Title string
	Desc string
	DueDate time.Time
	Completed bool
}

type TodoList struct{
	Items []TodoItem
}

func (tl *TodoList) AddItem(todoItem TodoItem){
	tl.Items = append(tl.Items, todoItem)
}

func (tl *TodoList) RemoveItem(id int){
	for i,item := range tl.Items{
		if item.
	}
}
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
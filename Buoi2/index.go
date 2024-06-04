package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New() // Initialize an empty list
	l.PushFront(2)
	l.PushBack(1)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == 5 {
			l.Remove(e)
		}
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}

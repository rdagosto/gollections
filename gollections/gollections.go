package gollections

import "fmt"

func Gollection() {
	fmt.Println("Here is Gollections")
}

type Collection struct {
	items []interface{}
}

func New(items ...interface{}) *Collection {
	return &Collection{items: items}
}

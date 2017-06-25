package main

import (
	"container/list"
	"log"
)

type Test struct {
	b int
}

func (this *Test) Test() {
	log.Println(this)
}

func main() {
	l := list.New()
	l.PushBack(&Test{b: 6})
	l.Remove(&list.Element{})
	log.Println(l)
}

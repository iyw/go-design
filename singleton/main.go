package main

import (
	"fmt"
	"sync"
)

type singleton map[string]string

var (
	once     sync.Once
	instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})
	return instance
}
func main() {
	s1 := New()
	fmt.Println(s1)
	s1["test"] = "tanyawen"
	s2 := New()
	fmt.Println(s2["test"])
}

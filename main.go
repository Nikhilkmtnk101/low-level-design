package main

import (
	"fmt"
	"reflect"
)

type singleton struct {
	counter int
}

func temp(t1, t2 interface{}) {
	t1, ok := t1.([]byte)
	if !ok {
		fmt.Println(reflect.DeepEqual(t1, t2))
	}
}
func main() {
	s1 := new(singleton)
	s2 := new(singleton)
	temp(s1, s2)
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "abcde"

	for i, c := range s {
		fmt.Println(i, reflect.TypeOf(c))
	}

	for i := 0; i < len(s); i++ {
		fmt.Println(reflect.TypeOf(s[i]))
	}
}

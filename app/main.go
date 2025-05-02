package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

func main() {
	nmap := make(map[int]int)

	nmap[2] = 55
	fmt.Println(nmap[1])
	fmt.Println(nmap[2])

	if nmap[3] == 0 {
		nmap[3]++
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Nice to meet you")
	}
	fmt.Println(nmap[3])
}

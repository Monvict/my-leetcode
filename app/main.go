package main

import "fmt"

func main() {
	s := "Hello"
	r := []byte{}
	for _, c := range s {
		charByte := byte(c)
		fmt.Println("charByte = ", charByte)
		if charByte >= 65 && charByte <= 90 {
			charByte = charByte + 32
		}

		r = append(r, charByte)
	}

	fmt.Println(r)
	fmt.Println(string(r))
}

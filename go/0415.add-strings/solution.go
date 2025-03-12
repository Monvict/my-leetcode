// Created by lotus at 2025/03/12 22:48
// leetgo: 1.4.13
// https://leetcode.com/problems/add-strings/

package main

import "fmt"

// @lc code=begin

func addStrings(num1 string, num2 string) string {
	i := len(num1) - 1
	j := len(num2) - 1

	resut := make([]byte, 0)
	carry := byte(0)

	for i >= 0 || j >= 0 || carry > 0 {
		a := byte(0)
		b := byte(0)

		if i >= 0 {
			a = num1[i] - '0'
		}

		if j >= 0 {
			b = num2[j] - '0'
		}

		sum := a + b + carry
		current := sum % 10
		carry = sum / 10

		resut = append(resut, current+'0')
		i--
		j--
	}

	reverse(resut)
	return string(resut)
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// @lc code=end

func main() {
	fmt.Println(addStrings("11", "222"))
}

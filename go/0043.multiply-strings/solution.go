// Created by lotus at 2025/03/12 10:35
// leetgo: 1.4.13
// https://leetcode.com/problems/multiply-strings/

package main

import "fmt"

// @lc code=begin

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	mResult := ""
	for i := len(num1) - 1; i >= 0; i-- {
		// 进位
		carry := byte(0)
		result := make([]byte, 0)
		for j := len(num2) - 1; j >= 0; j-- {
			// 被乘数
			a := num1[i] - '0'
			// 乘数
			b := num2[j] - '0'

			sum := a*b + carry
			current := sum % 10
			carry = sum / 10

			result = append(result, current+'0')
		}

		if carry > 0 {
			result = append(result, carry+'0')
		}

		reverse(result)

		// 补零：被乘数，从第2位起就要补0
		for k := i; k < len(num1)-1; k++ {
			result = append(result, '0')
		}

		// 每次被乘数和乘数得到一个结果后，就相加
		mResult = addStrings(mResult, string(result))
	}
	return mResult
}

// 字符串的加法
func addStrings(num1 string, num2 string) string {
	if len(num1) == 0 {
		return num2
	}

	if len(num2) == 0 {
		return num1
	}

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
	fmt.Println(multiply("123", "456"))
}

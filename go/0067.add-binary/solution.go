// Created by lotus at 2025/03/11 23:06
// leetgo: 1.4.13
// https://leetcode.com/problems/add-binary/

package main

// @lc code=begin
// 解题思路：
//
// 1. 将字符串相加，转变为数字相加 ---> 每个字符-'0'
//
// 2. 加法的思路是
//
// 2.1 sum = 第一个数 + 第二个数 + 进位
//
//	2.2 当前值 = sum % 2（类比 8 + 7 = 15， 当前位= 5 （15 % 10）；进位= 1 （15 / 10）
//	2.3 进位 = sum / 2
//
// 3. 循环遍历的条件是：a 或 b未处理完，或者进位 >= 1
func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := byte(0)
	strByte := make([]byte, 0)

	for i >= 0 || j >= 0 || carry > 0 {
		abit := byte(0)
		bbit := byte(0)

		if i >= 0 {
			abit = a[i] - '0'
		}

		if j >= 0 {
			bbit = b[j] - '0'
		}

		// 求和后，计算当前位、进位
		sum := abit + bbit + carry
		carry = sum / 2
		current := sum % 2

		strByte = append(strByte, '0'+current)

		i--
		j--
	}

	reverse(strByte)
	return string(strByte)
}

func reverse(strBytes []byte) {
	for i, j := 0, len(strBytes)-1; i < j; i, j = i+1, j-1 {
		strBytes[i], strBytes[j] = strBytes[j], strBytes[i]
	}
}

// @lc code=end

func main() {
	addBinary("11", "1")
}

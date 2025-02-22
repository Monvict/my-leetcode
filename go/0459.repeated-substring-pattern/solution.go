// Created by lotus at 2025/02/22 22:28
// leetgo: 1.4.13
// https://leetcode.com/problems/repeated-substring-pattern/

package main

// "bufio"
// "fmt"
// "os"

// . "github.com/j178/leetgo/testutils/go"

// @lc code=begin

func repeatedSubstringPattern(s string) bool {
	for i := 1; i < len(s); i++ {
		if len(s)%i != 0 {
			continue
		}

		if check(s, i) {
			return true
		}
	}

	return false
}

func check(s string, span int) bool {
	reaptStr := s[:span]
	for i := span; i < len(s); i = i + span {
		subStr := s[i : i+span]

		if reaptStr != subStr {
			return false
		}
	}
	return true
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// s := Deserialize[string](ReadLine(stdin))
	// ans := repeatedSubstringPattern(s)

	// fmt.Println("\noutput:", Serialize(ans))

	repeatedSubstringPattern("ababcd")
}

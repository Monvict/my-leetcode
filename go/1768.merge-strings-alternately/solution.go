// Created by lotus at 2025/02/20 21:34
// leetgo: 1.4.13
// https://leetcode.com/problems/merge-strings-alternately/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func mergeAlternately(word1 string, word2 string) string {
	wc1 := []rune(word1)
	wc2 := []rune(word2)
	i, len1 := 0, len(wc1)-1
	j, len2 := 0, len(wc2)-1

	wc := make([]rune, 0)
	for {
		// all done
		if i > len1 && j > len2 {
			break
		}

		// wc2 left
		if i > len1 && j <= len2 {
			wc = append(wc, wc2[j])
			j++
			continue
		}

		// wc1 left
		if j > len2 && i <= len1 {
			wc = append(wc, wc1[i])
			i++
			continue
		}

		// interact insert
		if i <= j {
			wc = append(wc, wc1[i])
			i++
		} else {
			wc = append(wc, wc2[j])
			j++
		}
	}

	return string(wc)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	word1 := Deserialize[string](ReadLine(stdin))
	word2 := Deserialize[string](ReadLine(stdin))
	ans := mergeAlternately(word1, word2)

	fmt.Println("\noutput:", Serialize(ans))

	// s1 := "ac"
	// s2 := "edf"
	// s := mergeAlternately(s1, s2)
	// fmt.Println(s)
}

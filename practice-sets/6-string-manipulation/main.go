package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		j := len(runes) - 1 - i
		if j <= i {
			break
		}
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseString2(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(s string) bool {
	runes := []rune{}
	for _, r := range s {
		if unicode.IsNumber(r) || unicode.IsLetter(r) {
			runes = append(runes, r)
		}
	}
	str := string(runes)

	// return strings.ToLower(str) == strings.ToLower(reverseString(str))
	return strings.EqualFold(str, reverseString(str))
}

func firstNonRepeatingCharacter(s string) string {
	runesMap := make(map[rune]int)
	// runes := []rune(s)
	for _, r := range s {
		if _, ok := runesMap[r]; !ok {
			runesMap[r] = 1
		} else {
			runesMap[r] += 1
		}
	}
	for _, r := range s {
		if val := runesMap[r]; val == 1 {
			return string(r)
		}
	}
	return ""
}

func lengthOfLongestSubstring1(s string) int {
	// My first attempt at it. Second attempt (below) is the better one.
	charMap := make(map[rune]int)
	runes := []rune(s)
	x := 0
	longest := 0
	for y := 0; y < len(runes); y++ {
		r := runes[y]
		if _, ok := charMap[r]; !ok {
			charMap[r] = 1
			longest = max(longest, len(charMap))
		} else {
			// remove runes[x] from map and x++ as long as charMap[r] is true
			for {
				if _, ok := charMap[r]; ok {
					delete(charMap, runes[x])
					x++
				} else {
					break
				}
			}
			y--
		}
	}
	return longest
}

func lengthOfLongestSubstring2(s string) int {
	charMap := make(map[rune]int)
	longest := 0
	left := 0
	for right, r := range s {
		// if curr char r is previously seen and its prev loc falls within curr window (i.e. >= left)
		// then update window size from left side (i.e. left = prev loc + 1)
		// to discard that prev char loc from curr window
		if idx, ok := charMap[r]; ok && idx >= left {
			left = charMap[r] + 1
		}
		charMap[r] = right
		longest = max(longest, right-left+1)
	}
	return longest
}

func stringCompression(s string) string {
	if len(s) == 0 {
		return s
	}

	count := 1
	var compressed strings.Builder

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			compressed.WriteByte(s[i-1])
			compressed.WriteString(strconv.Itoa(count))
			count = 1
		}
	}
	compressed.WriteByte(s[len(s)-1])
	compressed.WriteString(strconv.Itoa(count))
	res := compressed.String()
	if len(res) >= len(s) {
		return s
	}
	return res
}

func main() {
	p := fmt.Println

	p(reverseString("hello"))  // olleh
	p(reverseString2("hello")) // olleh

	p(isPalindrome("hello")) // false
	p(isPalindrome("civic")) // true
	p(isPalindrome("A man, a plan, a canal: Panama"))

	p(firstNonRepeatingCharacter("swiss"))    // "w"
	p(firstNonRepeatingCharacter("abcdcba"))  // "d"
	p(firstNonRepeatingCharacter("aabbccdd")) // ""

	p(lengthOfLongestSubstring1("abcabcbb"))
	p(lengthOfLongestSubstring2("abcabcbb"))

	p(stringCompression("aabcccccaaa")) // a2b1c5a3
	p(stringCompression("aabccccca"))   // a2b1c5a1

}

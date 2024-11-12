package main

import (
	"fmt"
	"regexp"
)

func main() {
	var match bool
	p := fmt.Println

	// 1. Match a string with a pattern
	match, _ = regexp.MatchString("at$", "bat")
	p(match) // true
	match, _ = regexp.MatchString("at$", "bot")
	p(match) // false

	// 2. Compile a pattern and match a string
	r := regexp.MustCompile(".[a-z]t")

	p(r.MatchString("bat")) // true
	p(r.MatchString("bot")) // false

	// 3. Find matched string
	p(r.FindString("bot sat bat"))    // bot
	p(r.FindString("plot slot blot")) // lot

	// 4. Find matched strings's index
	p(r.FindStringIndex("bot sat bat"))    // [0 3]
	p(r.FindStringIndex("plot slot blot")) // [1 4]

	// 5. Find all strings
	p(r.FindAllString("bot sat bat", -1)) // [bot sat bat]
	p(r.FindAllString("bot sat bat", 2))  // [bot sat]

	// 6. Find a string submatch
	r2 := regexp.MustCompile(`p([a-z]+)ch`)
	p(r2.FindStringSubmatch("peach pooch paach"))      // [peach ea]
	p(r2.FindStringSubmatchIndex("peach pooch paach")) // [0 5 1 3]

	// 7. Find all string submatches
	p(r2.FindAllStringSubmatch("peach pooch paach", -1))      // [[peach ea] [pooch oo] [paach aa]]
	p(r2.FindAllStringSubmatchIndex("peach pooch paach", -1)) // [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	// 8. Replace
	p(r2.ReplaceAllString("peach pooch paach", "foo")) // foo foo foo

	// 9. Split
	r3 := regexp.MustCompile(`\s+`)
	p(r3.Split("a b    c", -1)) // [a b c]
	p(r3.Split("a b    c", 3))  // [a b c]

}

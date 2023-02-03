/*
String Concatination
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func testConcat() {
	// Using + to concate
	s1 := "Go"
	s2 := "Programming"
	s3 := "Language"
	s := s1 + s2 + s3
	fmt.Println(s)
	fmt.Printf("s[0] if '%c' and is of type %T\n", s[0], s[0])

	// Using strings.Join to concate
	q := []string{"meetgor.com", "tags", "golang", "string"}
	r := strings.Join(q, "/")
	fmt.Println(r)

	// Using Sprintf method
	name := "peter"
	domain := "telecom"
	service := "ceo"

	email := fmt.Sprintf("%s.%s@%s.com", service, name, domain)
	fmt.Println(email)

	// Using strings.Builder function
	c := []string{"j", "a", "v", "a"}
	var builder strings.Builder
	for _, item := range c {
		builder.WriteString(item)
	}
	fmt.Println("Builder =", builder.String())
	b := []byte{'s', 'c', 'r', 'i', 'p', 't'}
	for _, item := range b {
		builder.WriteByte(item)
	}
	fmt.Println("Builder =", builder.String())
	builder.WriteRune('s')
	fmt.Println("Builder =", builder.String())
	fmt.Println("Builder =", builder)
}

func testSubstr() {
	s := "Japan"
	fmt.Println(s[0])
	fmt.Println(s[:2])
	fmt.Println(s[1:3])
	fmt.Printf("Type of s[1:3] is %T\n", s[1:3])

	// iterating strings
	for i, ch := range s {
		fmt.Printf("%d:%q ", i, ch)
	}
	fmt.Println()
	for i, ch := range "Japan 日本" {
		fmt.Printf("%d:%q ", i, ch)
	}
	fmt.Println()
	// This creates nonsense characters
	s = "Japan 日本"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%q ", s[i])
	}
}

func testSearch() {
	fmt.Println(strings.Contains("Japan", "abc"))     // checks if "abc" exists in "Japan"
	fmt.Println(strings.ContainsAny("Japan", "abc"))  // checks if a or b or c exists in "Japan"
	fmt.Println(strings.Count("Banana", "ana"))       // checks number of non-overlaping ana
	fmt.Println(strings.HasPrefix("Japan", "Ja"))     // check prefix
	fmt.Println(strings.HasSuffix("Japan", "an"))     // check prefix
	fmt.Println(strings.Index("Japan", "ap"))         // returns the first instance of "ap", else return -1
	fmt.Println(strings.Index("Japan", "ape"))        // returns the first instance of "ape", else return -1
	fmt.Println(strings.IndexAny("Japan", "abc"))     // return first instance of "a" or "b" or "c"
	fmt.Println(strings.IndexAny("Japan", "xyz"))     // return first instance of "x" or "y" or "z"
	fmt.Println(strings.LastIndex("Japan", "abc"))    // return the last instance of "abc"
	fmt.Println(strings.LastIndexAny("Japan", "abc")) // return the last instance of "abc"
}

func testReplace() {
	fmt.Println(strings.Replace("foo", "oo", ".", -1)) // replaces all "oo" with "."
	fmt.Println(strings.Replace("foo", "o", ".", 1))   // replaces first occurance of "o" with "."

	f := func(r rune) rune {
		return r + 1
	}
	fmt.Println(strings.Map(f, "abc"))

	fmt.Println(strings.ToUpper("Japan"))
	fmt.Println(strings.ToLower("Japan"))
	fmt.Println(strings.Title("japan country")) // deprecated

	fmt.Println(strings.TrimSpace("         abc   \n \n"))
	fmt.Println(strings.Trim("foodkoodffo", "fo")) // strip leading and trailing f or o's
	fmt.Println(strings.TrimLeft("foo", "f"))
	fmt.Println(strings.TrimRight("foo", "o"))

	fmt.Println(strings.TrimPrefix("foo", "fo"))
	fmt.Println(strings.TrimSuffix("foo", "o"))
}

func testSpilt() {
	fmt.Println(strings.Fields(" abc\t    bcd \n\n"))
	fmt.Println(strings.Split("a,b", ","))
	fmt.Println(strings.SplitAfter("a,b", ",")) // keeps the separator
}

func testConvert() {
	fmt.Println(strconv.Itoa(-42))
	fmt.Println(strconv.FormatInt(255, 32))
}

func main() {
	testSpilt()
}

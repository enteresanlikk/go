package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("contains", s.Contains("test", "es"))
	p("count", s.Count("test", "t"))
	p("hasprefix", s.HasPrefix("test", "te"))
	p("hassuffix", s.HasSuffix("test", "st"))
	p("index", s.Index("test", "e"))
	p("join", s.Join([]string{"t", "e"}, "-"))
	p("repeat", s.Repeat("te", 5))
	p("replace", s.Replace("foo", "o", "a", -1))
	p("replace", s.Replace("foo", "o", "b", -1))
	p("split", s.Split("a-b-c-d-e", "-"))
	p("tolower", s.ToLower("TEST"))
	p("toupper", s.ToUpper("test"))
}

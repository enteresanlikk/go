package main

import (
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	p := Point{1, 2}

	fmt.Printf("struct1: %v\n", p)
	fmt.Printf("struct2: %+v\n", p)
	fmt.Printf("struct3: %#v\n", p)
	fmt.Printf("type: %T\n", p)
	fmt.Printf("bool: %t\n", true)
	fmt.Printf("int %d\n", 123)
	fmt.Printf("bin: %b\n", 2)
	fmt.Printf("char: %c\n", 196)
	fmt.Printf("hex: %x\n", 255)
	fmt.Printf("float1: %f\n", 78.9)
	fmt.Printf("float2: %e\n", 78.9)
	fmt.Printf("float2: %E\n", 78.9)
	fmt.Printf("str1: %s\n", "\"string\"")
	fmt.Printf("str2: %q\n", "\"string\"")
	fmt.Printf("str3: %x\n", "hex this")
	fmt.Printf("pointer: %p\n", &p)
	fmt.Printf("width1: |%6d|%6d|\n", 12, 24)
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 2.4)
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 2.4)
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "bar")
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "bar")

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}

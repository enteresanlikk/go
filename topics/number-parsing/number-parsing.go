package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println(f, reflect.TypeOf(f))

	i, _ := strconv.ParseInt("32312", 0, 64)
	fmt.Println(i, reflect.TypeOf(i))

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d, reflect.TypeOf(d))

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u, reflect.TypeOf(u))

	k, _ := strconv.Atoi("135")
	fmt.Println(k, reflect.TypeOf(k))

	_, e := strconv.Atoi("wat")
	fmt.Println(e, reflect.TypeOf(e))
}

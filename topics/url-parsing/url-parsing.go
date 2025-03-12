package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println("scheme:", u.Scheme)
	fmt.Println("user:", u.User)
	fmt.Println("username:", u.User.Username())

	p, _ := u.User.Password()
	fmt.Println("password:", p)

	fmt.Println("host:", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("real host:", host)
	fmt.Print("port:", port)

	fmt.Println("path:", u.Path)
	fmt.Println("fragment:", u.Fragment)

	fmt.Println("raw query:", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("query:", m)
	fmt.Println(m["k"][0])
}

package main

import "fmt"

func main() {
	q := "ck"
	w := "what am i doing"

	u := x(q, w)
	p(u)
}

func x(y,z string) []byte {
	l := z + y

	return []byte(l)
}

func p(o []byte) {
	for i := 0; i < 25; i++ {
		fmt.Println(o)
	}
}

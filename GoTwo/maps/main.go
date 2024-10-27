package main

import (
	"fmt"
	"maps"
)

func main() {
	a := make(map[string]string)
	a["name"] = "goLang"
	a["area"] = "backend"
	fmt.Println(a["name"], a["area"])
	fmt.Println(a)
	b := make(map[string]int)
	b["age"] = 20
	b["price"] = 99
	delete(b, "age")
	fmt.Println(b["age"], b["phone"], b["price"])
	fmt.Println(len(b))
	clear(b)
	fmt.Println(len(b))

	c := map[string]int{"price": 40, "phones": 3}
	fmt.Println(c)

	k, ok := c["price"]
	if ok {
		fmt.Println("element is present", k)
	} else {
		fmt.Println("element is not present")
	}
	fmt.Println(maps.Equal(c, c))
}

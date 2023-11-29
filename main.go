package main

import "fmt"

func main() {
	var name string
	fmt.Scanf("%s", &name)
	res := hello(name)
	fmt.Println(res)
}

func hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

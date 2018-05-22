package gpDemo

import "fmt"

type Person struct {
	id		int
	name 	string
}

func SayHello(name string) {
	fmt.Printf("hello %s\n", name)
}

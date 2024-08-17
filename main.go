package main

import "fmt"

func SayHello(name string) string {
	msg := fmt.Sprintf("Hello %s", name)
	return msg
}

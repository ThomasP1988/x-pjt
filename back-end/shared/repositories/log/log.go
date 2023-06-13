package log

import "fmt"

func Add(text string) {
	fmt.Printf("text: %v\n", text)
}

func Error(text string, err error) {
	fmt.Printf("text: %v\n", text)
	fmt.Printf("err: %v\n", err)
}

package linewriter

import "fmt"

type LineWriter interface {
	WriteLine(val string)
}

type ConsoleWriter struct{}

func (c *ConsoleWriter) WriteLine(val string) {
	fmt.Println(val)
}

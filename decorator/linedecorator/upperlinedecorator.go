package linedecorator

import (
	"patrones/decorator/linewriter"
	"strings"
)

type UpperLineDecorator struct {
	writer linewriter.LineWriter
}

func NewUpperLineDecorator(lw linewriter.LineWriter) *UpperLineDecorator {
	return &UpperLineDecorator{
		writer: lw,
	}
}

func (u *UpperLineDecorator) WriteLine(val string) {
	u.writer.WriteLine(strings.ToUpper(val))
}

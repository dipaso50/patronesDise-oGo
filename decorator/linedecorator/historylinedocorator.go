package linedecorator

import (
	"fmt"
	"patrones/decorator/linewriter"
	"strings"
	"time"
)

type HistoryLineDecorator struct {
	writer  linewriter.LineWriter
	history map[time.Time]string
}

func NewHistoryLineDecorator(lw linewriter.LineWriter) *HistoryLineDecorator {
	return &HistoryLineDecorator{
		writer:  lw,
		history: make(map[time.Time]string),
	}
}

func (u *HistoryLineDecorator) WriteLine(val string) {
	u.history[time.Now()] = val
	u.writer.WriteLine(strings.ToUpper(val))
}

func (u *HistoryLineDecorator) PrintHistory() {
	for key, value := range u.history {
		fmt.Println("time:", key.Format(time.RFC3339), "Value:", value)
	}
}

package disassemble

import (
	"bufio"
	"fmt"
)

type ConsoleIterationWriter struct {
}

func (ir ConsoleIterationWriter) Write(ic IterationChanges) {

	fmt.Println(ic.ToString())
}

type WriterIterationWriter struct {
	Writer bufio.Writer
}

// Streight order instead of reverse
func (ir WriterIterationWriter) Write(ic IterationChanges) {
	_, err := ir.Writer.WriteString(fmt.Sprintf("%s\n", ic.ToString()))
	if err != nil {
		panic(err)
	}
	ir.Writer.Flush()
}

type StringSliceWriter struct {
	Slice  []string
	String string
}

func (is StringSliceWriter) Write(ic IterationChanges) {
	str := ic.ToString() + "\n"
	is.Slice[0] = str + is.Slice[0]
}

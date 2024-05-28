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

func (ir WriterIterationWriter) Write(ic IterationChanges) {
	_, err := ir.Writer.WriteString(fmt.Sprintf("%s\n", ic.ToString()))
	if err != nil {
		panic(err)
	}
	ir.Writer.Flush()
}

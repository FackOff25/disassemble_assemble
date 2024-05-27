package disassemble

import "fmt"

type ConsoleIterationWriter struct {
}

func (ir ConsoleIterationWriter) Write(ic IterationChanges) {
	fmt.Println(ic.ToString())
}

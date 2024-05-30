package assemble

import (
	"bufio"
	"encoding/json"

	"github.com/FackOff25/disassemble_assemble/disassemble"
)

type StraightScannerIterationReader struct {
	scanner *bufio.Scanner
}

func (r StraightScannerIterationReader) New(_scanner *bufio.Scanner) {
	r.scanner = _scanner
}

func (r StraightScannerIterationReader) ReadNextIteration() IterationChanges {
	if !r.scanner.Scan() {
		return IterationChanges{}
	}
	str := r.scanner.Text()
	var disassembleChanges disassemble.IterationChanges
	json.Unmarshal([]byte(str), &disassembleChanges)
	result := IterationChanges{
		AddedNodes:   disassembleChanges.RemovedNodes,
		RemovedEdges: disassembleChanges.AddedEdges,
	}

	return result
}

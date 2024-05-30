package assemble

import (
	"bufio"
	"encoding/json"
	"io"

	"github.com/FackOff25/disassemble_assemble/disassemble"
)

type StraightScannerIterationReader struct {
	Reader *bufio.Reader
}

func (r StraightScannerIterationReader) ReadNextIteration() IterationChanges {
	str, err := r.Reader.ReadString('\n')
	if err != nil {
		if err != io.EOF {
			panic(err)
		}
	}

	var disassembleChanges disassemble.IterationChanges
	json.Unmarshal([]byte(str), &disassembleChanges)
	result := IterationChanges{
		AddedNodes:   disassembleChanges.RemovedNodes,
		RemovedEdges: disassembleChanges.AddedEdges,
	}

	return result
}

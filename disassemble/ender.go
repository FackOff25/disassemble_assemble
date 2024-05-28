package disassemble

import (
	"github.com/FackOff25/disassemble_assemble/graph"
)

type NodeNumEnder struct {
	NodeNum int
}

func (rc NodeNumEnder) CheckPruningEnd(_graph graph.Graph) bool {
	return !(len(_graph.Nodes) > rc.NodeNum)
}

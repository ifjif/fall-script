package ir

import (
	"fmt"
)

type ExportStmt struct {
	Name     string
	GlobalId int
}

func (es *ExportStmt) stmtNode() {}
func (es *ExportStmt) String() string {
	return fmt.Sprintf("globalId: %d name: %s", es.GlobalId, es.Name)
}

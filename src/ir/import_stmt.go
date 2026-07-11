package ir

import "fmt"

type ImportStmt struct {
	GlobalId int
	From     string
	Imported string
	Local    string
}

func (is *ImportStmt) stmtNode() {}

func (is *ImportStmt) String() string {
	return fmt.Sprintf("globalId: %d from: %s imported: %s local: %s", is.GlobalId, is.From, is.Imported, is.Local)
}

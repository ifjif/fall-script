package ir

import (
	"bytes"
	"fmt"

	"zzc/fall-script/src/object"
)

type StructDeclStmt struct {
	Name     *Symbol
	MetaData *object.StructMeta
}

func (sd *StructDeclStmt) stmtNode() {}

func (sd *StructDeclStmt) String() string {
	var buf bytes.Buffer
	buf.WriteString("struct ")
	name := fmt.Sprintf("(name: %s pos: %d scope: %s captured: %t)", sd.Name.Name, sd.Name.Pos, sd.Name.Scope, sd.Name.Captured)
	buf.WriteString(name)
	buf.WriteString("{\n")
	buf.WriteString("fields:\n")
	for _, f := range sd.MetaData.Fields {
		msg := fmt.Sprintf("name: %s index: %d isEmbed: %t targetMeta: %v\n", f.Name, f.Index, f.IsEmbed, f.TargetStructMeta)
		buf.WriteString(msg)
	}

	buf.WriteString("\nmethods:\n")
	for _, m := range sd.MetaData.Methods {
		msg := fmt.Sprintf("name: %s targetMetaIndex: %d index: %d\n", m.Name, m.TargetStructMeta, m.Index)
		buf.WriteString(msg)
	}
	buf.WriteString("}")

	return buf.String()
}

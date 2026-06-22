package binarychunck

import (
	"bytes"
	"encoding/binary"

	"zzc/fall-script/src/object"
)

func Dump(module *object.Module) []byte {
	var buf bytes.Buffer
	writeHeader(&buf)
	writeModule(&buf, module)

	return buf.Bytes()
}

func writeHeader(buf *bytes.Buffer) {
	buf.WriteString(SIGNATURE)
	buf.WriteByte(MAJOR)
	buf.WriteByte(MINOR)
	buf.WriteByte(PATCH)
}

/*
*

	Name      string
	GlobalNum int
	Imports   []*ImportRef
	Exports   []*ExportRef
	Cf        *CompiledFunction

*
*/
func writeModule(buf *bytes.Buffer, module *object.Module) {
	writeModuleName(buf, module.Name)
	writeModuleGlobalNum(buf, module.GlobalNum)
	writeModuleImports(buf, module.Imports)
	writeModuleExports(buf, module.Exports)
	writeFunction(buf, module.Cf)
}

func writeModuleName(buf *bytes.Buffer, name string) {
	length := len(name)
	writeUint32(buf, length)
	buf.WriteString(name)
}

func writeModuleGlobalNum(buf *bytes.Buffer, globalNum int) {
	writeUint16(buf, globalNum)
}

func writeModuleImports(buf *bytes.Buffer, imports []*object.ImportRef) {
	num := len(imports)
	writeUint16(buf, num)
	for _, imp := range imports {
		writeModuleImport(buf, imp)
	}
}

/*
*

	type ImportRef struct {
		From     int
		Imported int
		Local    int
	}
*/
func writeModuleImport(buf *bytes.Buffer, imp *object.ImportRef) {
	writeUint16(buf, imp.From)
	writeUint16(buf, imp.Imported)
	writeUint16(buf, imp.Local)
}

/*
*

	type ExportRef struct {
		Name     int
		GlobalId int
	}
*/
func writeModuleExports(buf *bytes.Buffer, exports []*object.ExportRef) {
	num := len(exports)
	writeUint16(buf, num)
	for _, exp := range exports {
		writeModuleExport(buf, exp)
	}
}

func writeModuleExport(buf *bytes.Buffer, exp *object.ExportRef) {
	writeUint16(buf, exp.Name)
	writeUint16(buf, exp.GlobalId)
}

func writeFunction(buf *bytes.Buffer, cf *object.CompiledFunction) {
	writeMaxStackDepth(buf, cf.StackDepth)
	writeLocalVarNum(buf, cf.LocalsNum)
	writeConstants(buf, cf.Constants)
	writeInstructions(buf, cf.Instructions)
}

func writeMaxStackDepth(buf *bytes.Buffer, depth int) {
	buf.WriteByte(byte(depth))
}

func writeLocalVarNum(buf *bytes.Buffer, num int) {
	buf.WriteByte(byte(num))
}

func writeConstants(buf *bytes.Buffer, consts []object.Object) {
	num := len(consts)
	writeUint16(buf, num)
	for _, con := range consts {
		writeConstant(buf, con)
	}
}

func writeConstant(buf *bytes.Buffer, con object.Object) {
	switch con := con.(type) {
	case *object.Integer:
		buf.WriteByte(I64)
		binary.Write(buf, binary.BigEndian, con.Value)
	case *object.String:
		length := len(con.Value)
		buf.WriteByte(STR)
		writeUint32(buf, length)
		buf.WriteString(con.Value)
	case *object.CompiledFunction:
		buf.WriteByte(CF)
		writeFunction(buf, con)
	}
}

func writeInstructions(buf *bytes.Buffer, instructions []byte) {
	length := len(instructions)
	writeUint32(buf, length)
	buf.Write(instructions)
}

func writeUint16(buf *bytes.Buffer, i int) {
	binary.Write(buf, binary.BigEndian, uint16(i))
}

func writeUint32(buf *bytes.Buffer, i int) {
	binary.Write(buf, binary.BigEndian, uint32(i))
}

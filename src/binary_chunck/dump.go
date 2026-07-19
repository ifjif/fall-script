package binarychunck

import (
	"zzc/fall-script/src/object"
)

func Dump(module *object.Module) []byte {
	var buf FlWriter
	writeHeader(&buf)
	writeModule(&buf, module)

	return buf.Bytes()
}

func writeHeader(buf *FlWriter) {
	buf.WriteStr(SIGNATURE)
	buf.WriteUint8(MAJOR)
	buf.WriteUint8(MINOR)
	buf.WriteUint8(PATCH)
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
func writeModule(buf *FlWriter, module *object.Module) {
	writeModuleName(buf, module.Name)
	writeModuleGlobalNum(buf, module.GlobalNum)
	writeModuleImports(buf, module.Imports)
	writeModuleExports(buf, module.Exports)
	writeModuleStructMetaConstantPosition(buf, module.Structs)
	writeFunction(buf, module.Cf)
}

func writeModuleName(buf *FlWriter, name string) {
	buf.WriteStr(name)
}

func writeModuleGlobalNum(buf *FlWriter, globalNum int) {
	buf.WriteUint16(uint16(globalNum))
}

func writeModuleImports(buf *FlWriter, imports []*object.ImportRef) {
	num := len(imports)
	buf.WriteUint16(uint16(num))
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
func writeModuleImport(buf *FlWriter, imp *object.ImportRef) {
	buf.WriteUint16(uint16(imp.From))
	buf.WriteUint16(uint16(imp.Imported))
	buf.WriteUint16(uint16(imp.Local))
}

/*
*

	type ExportRef struct {
		Name     int
		GlobalId int
	}
*/
func writeModuleExports(buf *FlWriter, exports []*object.ExportRef) {
	num := len(exports)
	buf.WriteUint16(uint16(num))
	for _, exp := range exports {
		writeModuleExport(buf, exp)
	}
}

func writeModuleExport(buf *FlWriter, exp *object.ExportRef) {
	buf.WriteUint16(uint16(exp.Name))
	buf.WriteUint16(uint16(exp.GlobalId))
}

// structMeta 在常量池中的位置
func writeModuleStructMetaConstantPosition(buf *FlWriter, exports map[int]*object.StructMeta) {
	count := len(exports)
	buf.WriteVarint(uint64(count))
	for idx := range exports {
		buf.WriteVarint(uint64(idx))
	}
}

func writeFunction(buf *FlWriter, cf *object.CompiledFunction) {
	writeMaxStackDepth(buf, cf.StackDepth)
	writeLocalVarNum(buf, cf.LocalsNum)
	writeConstants(buf, cf.Constants)
	writeInstructions(buf, cf.Instructions)
}

func writeMaxStackDepth(buf *FlWriter, depth int) {
	buf.WriteUint8(byte(depth))
}

func writeLocalVarNum(buf *FlWriter, num int) {
	buf.WriteUint8(byte(num))
}

/*
*

	Name        string
	FieldCount  int
	Fields      map[string]*FieldInfo
	Methods     map[string]*MethodRef

*
*/
func writeStructMeta(buf *FlWriter, sm *object.StructMeta) {
	buf.WriteStr(sm.Name)
	buf.WriteVarint(uint64(sm.FieldCount))
	writeStructMetaFields(buf, sm.Fields)
	writeStructMetaMethods(buf, sm.Methods)
}

func writeStructMetaFields(buf *FlWriter, fields map[string]*object.FieldInfo) {
	count := len(fields)
	buf.WriteVarint(uint64(count))

	for _, field := range fields {
		writeStructMetaField(buf, field)
	}
}

/*
*
	Name             string
	IsEmbed          bool
	Index            int
	TargetStructMeta int
*
*/

func writeStructMetaField(buf *FlWriter, field *object.FieldInfo) {
	buf.WriteStr(field.Name)
	buf.WriteBool(field.IsEmbed)
	buf.WriteVarint(uint64(field.Index))
	buf.WriteVarint(uint64(field.TargetStructMeta))
}

func writeStructMetaMethods(buf *FlWriter, methods map[string]*object.MethodRef) {
	count := len(methods)
	buf.WriteVarint(uint64(count))
	for _, method := range methods {
		writeStructMetaMethod(buf, method)
	}
}

/*
*

	TargetStructMeta int
	Name             string
	Index            int
*/
func writeStructMetaMethod(buf *FlWriter, method *object.MethodRef) {
	buf.WriteVarint(uint64(method.TargetStructMeta))
	buf.WriteStr(method.Name)
	buf.WriteVarint(uint64(method.Index))
}

func writeConstants(buf *FlWriter, consts []object.Object) {
	num := len(consts)
	buf.WriteUint16(uint16(num))
	for _, con := range consts {
		writeConstant(buf, con)
	}
}

func writeConstant(buf *FlWriter, con object.Object) {
	switch con := con.(type) {
	case *object.Integer:
		buf.WriteUint8(I64)
		buf.WriteVarint(uint64(con.Value))
	case *object.String:
		buf.WriteUint8(STR)
		buf.WriteStr(con.Value)
	case *object.CompiledFunction:
		buf.WriteUint8(CF)
		writeFunction(buf, con)
	case *object.StructMeta:
		buf.WriteUint8(STRUCT_META)
		writeStructMeta(buf, con)
	}
}

func writeInstructions(buf *FlWriter, instructions []byte) {
	length := len(instructions)
	buf.WriteVarint(uint64(length))
	buf.WriteBytes(instructions)
}

package binarychunck

import (
	"zzc/fall-script/src/object"
)

func Undump(data []byte) *object.Module {
	reader := &FlReader{data: data}
	readHeader(reader)
	mo := readModule(reader)
	return mo
}

func readHeader(r *FlReader) {
	signature := r.ReadStr()
	major := r.ReadUint8()
	minor := r.ReadUint8()
	patch := r.ReadUint8()
	validateHeader(signature, int(major), int(minor), int(patch))
	// fmt.Println(signature)
	// fmt.Println(major)
	// fmt.Println(minor)
	// fmt.Println(patch)
}

func validateHeader(signature string, major, minor, patch int) {
}

func readModule(r *FlReader) *object.Module {
	mo := &object.Module{}
	mo.Name = readModuleName(r)
	mo.GlobalNum = readModuleGlobalNum(r)
	mo.Imports = readModuleImports(r)
	mo.Exports = readModuleExports(r)
	smCidxs := readModuleStructMetaConstantPosition(r)
	mo.Cf = readFunction(r)
	consts := mo.Cf.Constants
	structs := map[int]*object.StructMeta{}
	for _, v := range smCidxs {
		structs[v] = consts[v].(*object.StructMeta)
	}
	mo.Structs = structs

	return mo
}

func readModuleName(r *FlReader) string {
	name := r.ReadStr()
	return name
}

func readModuleGlobalNum(r *FlReader) int {
	n := r.ReadUint16()
	return int(n)
}

func readModuleImports(r *FlReader) []*object.ImportRef {
	num := int(r.ReadUint16())
	imports := make([]*object.ImportRef, num)
	for i := range num {
		imp := readModuleImport(r)
		imports[i] = imp
	}

	return imports
}

func readModuleImport(r *FlReader) *object.ImportRef {
	imp := &object.ImportRef{}
	imp.From = int(r.ReadUint16())
	imp.Imported = int(r.ReadUint16())
	imp.Local = int(r.ReadUint16())
	return imp
}

func readModuleExports(r *FlReader) []*object.ExportRef {
	num := int(r.ReadUint16())
	exports := make([]*object.ExportRef, num)
	for i := range num {
		exp := readModuleExport(r)
		exports[i] = exp
	}

	return exports
}

func readModuleExport(r *FlReader) *object.ExportRef {
	exp := &object.ExportRef{}
	exp.Name = int(r.ReadUint16())
	exp.GlobalId = int(r.ReadUint16())
	return exp
}

func readModuleStructMetaConstantPosition(r *FlReader) []int {
	count := r.ReadVarint()
	cIdxs := make([]int, count)
	for i := range count {
		cIdxs[i] = int(r.ReadVarint())
	}

	return cIdxs
}

func readFunction(r *FlReader) *object.CompiledFunction {
	cf := &object.CompiledFunction{}
	cf.StackDepth = readMaxStackDepth(r)
	cf.LocalsNum = readLocalVarNum(r)
	cf.Constants = readConstants(r)
	cf.Instructions = readInstructions(r)
	return cf
}

func readMaxStackDepth(r *FlReader) int {
	d := r.ReadUint8()
	return int(d)
}

func readLocalVarNum(r *FlReader) int {
	d := r.ReadUint8()
	return int(d)
}

/*
*

	Name        string
	FieldCount  int
	Fields      map[string]*FieldInfo
	Methods     map[string]*MethodRef

*
*/
func readStructMeta(r *FlReader) *object.StructMeta {
	sm := &object.StructMeta{}
	sm.Name = r.ReadStr()
	sm.FieldCount = int(r.ReadVarint())
	sm.Fields = readStructMetaFields(r)
	sm.Methods = readStructMetaMethods(r)

	return sm
}

func readStructMetaFields(r *FlReader) map[string]*object.FieldInfo {
	count := r.ReadVarint()
	fields := map[string]*object.FieldInfo{}
	for range count {
		field := readStructMetaField(r)
		fields[field.Name] = field
	}

	return fields
}

/*
*

	Name             string
	IsEmbed          bool
	Index            int
	TargetStructMeta int

*
*/
func readStructMetaField(r *FlReader) *object.FieldInfo {
	field := &object.FieldInfo{}
	field.Name = r.ReadStr()
	field.IsEmbed = r.ReadBool()
	field.Index = int(r.ReadVarint())
	field.TargetStructMeta = int(r.ReadVarint())

	return field
}

func readStructMetaMethods(r *FlReader) map[string]*object.MethodRef {
	count := r.ReadVarint()
	methods := map[string]*object.MethodRef{}
	for range count {
		method := readStructMetaMethod(r)
		methods[method.Name] = method
	}

	return methods
}

/*
*

	TargetStructMeta int
	Name             string
	Index            int
*/
func readStructMetaMethod(r *FlReader) *object.MethodRef {
	method := &object.MethodRef{}
	method.TargetStructMeta = int(r.ReadVarint())
	method.Name = r.ReadStr()
	method.Index = int(r.ReadVarint())

	return method
}

func readConstants(r *FlReader) []object.Object {
	num := int(r.ReadUint16())
	consts := make([]object.Object, num)

	for i := range consts {
		consts[i] = readConstant(r)
	}

	return consts
}

func readConstant(r *FlReader) object.Object {
	typeTag := r.ReadUint8()

	switch typeTag {
	case I64:
		v := r.ReadVarint()
		return &object.Integer{Value: int64(v)}
	case STR:
		str := r.ReadStr()
		return &object.String{Value: string(str)}
	case CF:
		return readFunction(r)
	case STRUCT_META:
		return readStructMeta(r)
	}

	panic("unknown type constant")
}

func readInstructions(r *FlReader) []byte {
	length := r.ReadVarint()
	data := r.ReadBytes(length)
	return data
}

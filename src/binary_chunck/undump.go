package binarychunck

import (
	"encoding/binary"

	"zzc/fall-script/src/object"
)

type Reader struct {
	data []byte
}

func (r *Reader) readUint8() uint8 {
	d := r.data[0]
	r.data = r.data[1:]
	return d
}

func (r *Reader) readUint16() uint16 {
	d := binary.BigEndian.Uint16(r.data)
	r.data = r.data[2:]
	return d
}

func (r *Reader) readUint32() uint32 {
	d := binary.BigEndian.Uint32(r.data)
	r.data = r.data[4:]
	return d
}

func (r *Reader) readUint64() uint64 {
	d := binary.BigEndian.Uint64(r.data)
	r.data = r.data[8:]
	return d
}

func (r *Reader) readBytes(num int) []byte {
	b := r.data[:num]
	r.data = r.data[num:]
	return b
}

func Undump(data []byte) *object.Module {
	reader := &Reader{data: data}
	readHeader(reader)
	mo := readModule(reader)
	return mo
}

func readHeader(r *Reader) {
	signature := string(r.readBytes(len(SIGNATURE)))
	major := r.readUint8()
	minor := r.readUint8()
	patch := r.readUint8()
	validateHeader(signature, int(major), int(minor), int(patch))
	// fmt.Println(signature)
	// fmt.Println(major)
	// fmt.Println(minor)
	// fmt.Println(patch)
}

func validateHeader(signature string, major, minor, patch int) {
}

func readModule(r *Reader) *object.Module {
	mo := &object.Module{}
	mo.Name = readModuleName(r)
	mo.GlobalNum = readModuleGlobalNum(r)
	mo.Imports = readModuleImports(r)
	mo.Exports = readModuleExports(r)
	mo.Cf = readFunction(r)

	return mo
}

func readModuleName(r *Reader) string {
	length := r.readUint32()
	name := r.readBytes(int(length))
	return string(name)
}

func readModuleGlobalNum(r *Reader) int {
	n := r.readUint16()
	return int(n)
}

func readModuleImports(r *Reader) []*object.ImportRef {
	num := int(r.readUint16())
	imports := make([]*object.ImportRef, num)
	for i := range num {
		imp := readModuleImport(r)
		imports[i] = imp
	}

	return imports
}

func readModuleImport(r *Reader) *object.ImportRef {
	imp := &object.ImportRef{}
	imp.From = int(r.readUint16())
	imp.Imported = int(r.readUint16())
	imp.Local = int(r.readUint16())
	return imp
}

func readModuleExports(r *Reader) []*object.ExportRef {
	num := int(r.readUint16())
	exports := make([]*object.ExportRef, num)
	for i := range num {
		exp := readModuleExport(r)
		exports[i] = exp
	}

	return exports
}

func readModuleExport(r *Reader) *object.ExportRef {
	exp := &object.ExportRef{}
	exp.Name = int(r.readUint16())
	exp.GlobalId = int(r.readUint16())
	return exp
}

func readFunction(r *Reader) *object.CompiledFunction {
	cf := &object.CompiledFunction{}
	cf.StackDepth = readMaxStackDepth(r)
	cf.LocalsNum = readLocalVarNum(r)
	cf.Constants = readConstants(r)
	cf.Instructions = readInstructions(r)
	return cf
}

func readMaxStackDepth(r *Reader) int {
	d := r.readUint8()
	return int(d)
}

func readLocalVarNum(r *Reader) int {
	d := r.readUint8()
	return int(d)
}

func readConstants(r *Reader) []object.Object {
	num := int(r.readUint16())
	consts := make([]object.Object, num)

	for i := range consts {
		consts[i] = readConstant(r)
	}

	return consts
}

func readConstant(r *Reader) object.Object {
	typeTag := r.readUint8()

	switch typeTag {
	case I64:
		v := r.readUint64()
		return &object.Integer{Value: int64(v)}
	case STR:
		num := int(r.readUint32())
		str := r.readBytes(num)
		return &object.String{Value: string(str)}
	case CF:
		return readFunction(r)
	}

	panic("unknown type constant")
}

func readInstructions(r *Reader) []byte {
	length := r.readUint32()
	data := r.readBytes(int(length))
	return data
}

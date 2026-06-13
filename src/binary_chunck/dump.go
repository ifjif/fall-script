package binarychunck

import (
	"bytes"
	"encoding/binary"

	"zzc/fall-script/src/object"
)

func Dump(cf *object.CompiledFunction) []byte {
	var buf bytes.Buffer
	writeHeader(&buf)
	writeFunction(&buf, cf)

	return buf.Bytes()
}

func writeHeader(buf *bytes.Buffer) {
	buf.WriteString(SIGNATURE)
	buf.WriteByte(MAJOR)
	buf.WriteByte(MINOR)
	buf.WriteByte(PATCH)
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
	binary.Write(buf, binary.BigEndian, uint32(num))
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
		buf.WriteByte(Str)
		binary.Write(buf, binary.BigEndian, uint32(length))
		buf.WriteString(con.Value)
	case *object.CompiledFunction:
		buf.WriteByte(CF)
		writeFunction(buf, con)
	}
}

func writeInstructions(buf *bytes.Buffer, instructions []byte) {
	length := len(instructions)
	binary.Write(buf, binary.BigEndian, uint32(length))
	buf.Write(instructions)
}

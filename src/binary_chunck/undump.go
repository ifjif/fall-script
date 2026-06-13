package binarychunck

import (
	"encoding/binary"
	"fmt"

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

func Undump(data []byte) *object.CompiledFunction {
	reader := &Reader{data: data}
	readHeader(reader)
	cf := readFunction(reader)
	return cf
}

func readHeader(r *Reader) {
	signature := string(r.readBytes(len(SIGNATURE)))
	fmt.Println(signature)
	major := r.readUint8()
	fmt.Println(major)
	minor := r.readUint8()
	fmt.Println(minor)
	patch := r.readUint8()
	fmt.Println(patch)
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
	num := int(r.readUint32())
	consts := make([]object.Object, num)

	for i := 0; i < num; i++ {
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
	case Str:
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

package base

import "encoding/binary"

type ByteReader struct {
	code []byte
	pc   int
}

func (br *ByteReader) Pc() int {
	return br.pc
}

func (br *ByteReader) Reset(code []byte, pc int) {
	br.code = code
	br.pc = pc
}

func (br *ByteReader) ReadUint8() uint8 {
	data := br.code[br.pc]
	br.pc++

	return data
}

func (br *ByteReader) ReadUint16() uint16 {
	data := binary.BigEndian.Uint16(br.code[br.pc:])
	br.pc += 2
	return data
}

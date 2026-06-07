package base

import "encoding/binary"

type ByteReader struct {
	code []byte
	pc   int
}

func NewByteReader() *ByteReader {
	return &ByteReader{}
}

func (br *ByteReader) PC() int {
	return br.pc
}

func (br *ByteReader) Reset(code []byte, pc int) {
	br.code = code
	br.pc = pc
}

func (br *ByteReader) ReadUint8() uint8 {
	i := br.code[br.pc]
	br.pc++
	return i
}

func (br *ByteReader) ReadUint16() uint16 {
	v := binary.BigEndian.Uint16(br.code[br.pc:])
	br.pc += 2
	return v
}

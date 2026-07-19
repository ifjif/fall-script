package binarychunck

import (
	"bytes"
	"encoding/binary"
)

type FlWriter struct {
	bytes.Buffer
}

func (fw *FlWriter) WriteBool(b bool) {
	v := 0
	if b {
		v = 1
	}
	fw.WriteByte(byte(v))
}

func (fw *FlWriter) WriteUint8(d byte) {
	fw.WriteByte(d)
}

func (fw *FlWriter) WriteUint16(d uint16) {
	binary.Write(fw, binary.BigEndian, d)
}

func (fw *FlWriter) WriteUint32(d uint32) {
	binary.Write(fw, binary.BigEndian, d)
}

func (fw *FlWriter) WriteUint64(d uint64) {
	binary.Write(fw, binary.BigEndian, d)
}

func (fw *FlWriter) WriteBytes(data []byte) {
	_, err := fw.Write(data)
	if err != nil {
		panic(err)
	}
}

func (fw *FlWriter) WriteStr(str string) {
	length := len(str)
	fw.WriteVarint(uint64(length))
	_, err := fw.WriteString(str)
	if err != nil {
		panic(err)
	}
}

// 最多 10字节
func (fw *FlWriter) WriteVarint(d uint64) {
	var buf []byte

	for d >= 0x80 {
		buf = append(buf, byte(d)|0x80)
		d >>= 7
	}

	buf = append(buf, byte(d))

	for _, d := range buf {
		fw.WriteByte(d)
	}
}

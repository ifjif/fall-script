package binarychunck

import "encoding/binary"

type FlReader struct {
	data []byte
}

func (fr *FlReader) Bytes() []byte {
	return fr.data
}

func (fr *FlReader) InitData(data []byte) {
	if len(fr.data) > 0 {
		return
	}
	fr.data = data
}

func (fr *FlReader) ReadBool() bool {
	d := fr.ReadUint8()
	if d == 0 {
		return false
	}
	return true
}

func (fr *FlReader) ReadUint8() uint8 {
	d := fr.data[0]
	fr.data = fr.data[1:]
	return d
}

func (fr *FlReader) ReadUint16() uint16 {
	d := binary.BigEndian.Uint16(fr.data)
	fr.data = fr.data[2:]
	return d
}

func (fr *FlReader) ReadUint32() uint32 {
	d := binary.BigEndian.Uint32(fr.data)
	fr.data = fr.data[4:]
	return d
}

func (fr *FlReader) ReadUint64() uint64 {
	d := binary.BigEndian.Uint64(fr.data)
	fr.data = fr.data[8:]
	return d
}

func (fr *FlReader) ReadBytes(num uint64) []byte {
	b := fr.data[:num]
	fr.data = fr.data[num:]
	return b
}

func (fr *FlReader) ReadStr() string {
	length := fr.ReadVarint()
	return string(fr.ReadBytes(length))
}

// 0000000 0000000
// 最多 10字节
func (fr *FlReader) ReadVarint() uint64 {
	var x uint64
	var shift uint

	b := fr.ReadUint8()

	for b >= 0x80 {
		x |= uint64(b&0x7f) << shift
		shift += 7
		b = fr.ReadUint8()
	}

	if shift >= 64 {
		panic("varint overflow")
	}

	x |= uint64(b&0x7f) << shift

	return x
}

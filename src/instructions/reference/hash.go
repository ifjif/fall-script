package reference

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type Hash struct {
	kvNum int
}

func (h *Hash) FetchOperand(br *base.ByteReader) {
	kvNum := br.ReadUint16()
	h.kvNum = int(kvNum)
}

func (h *Hash) Execute(frame *rt.Frame) {
	kv := frame.PopStacks(h.kvNum)

	pairs := make(map[object.HashKey]object.HashPair)
	for i := 0; i < len(kv); i = i + 2 {
		k := kv[i]

		hashKey, ok := k.(object.HashTableKey)
		if !ok {
			panic("Error: not usable hash key!")
		}
		v := kv[i+1]

		pair := object.HashPair{Key: k, Value: v}

		pairs[hashKey.HashKey()] = pair
	}

	hash := &object.Hash{Pairs: pairs}

	frame.PushStack(hash)
}

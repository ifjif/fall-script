package base

import (
	"zzc/fall-script/src/vm/rt"
)

type Nop struct{}

func (n *Nop) FetchOperand(br *ByteReader) {}

func (n *Nop) Execute(frame *rt.Frame) {}

package vm

import (
	"zzc/fall-script/src/code"
	"zzc/fall-script/src/instructions"
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/vm/rt"
)

func interpreter(thread *rt.Thread) {
	loop(thread)
}

func loop(thread *rt.Thread) {
	br := base.NewByteReader()
	for {
		frame := thread.CurrentFrame()
		thread.Pc = frame.NextPc
		br.Reset(frame.Code, frame.NextPc)
		opcode := br.ReadUint8()
		inst := instructions.NewInstruction(code.OpCode(opcode))
		inst.FetchOperand(br)
		frame.NextPc = br.PC()
		inst.Execute(frame)
	}
}

package vm

import (
	"zzc/fall-script/src/instructions"
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/vm/rt"
)

func interpreter(thread *rt.Thread) {
	defer catchError()
	loop(thread)
}

func catchError() {
	if r := recover(); r != nil {
		// fmt.Println(r)
		panic(r)
	}
}

func loop(thread *rt.Thread) {
	br := &base.ByteReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPc()
		thread.SetPc(pc)
		code := frame.Code()
		br.Reset(code, pc)
		inst := instructions.NewInstruction(br.ReadUint8())
		inst.FetchOperand(br)
		frame.SetNextPc(br.Pc())
		inst.Execute(frame)

		if thread.IsEmpty() {
			break
		}
	}
}

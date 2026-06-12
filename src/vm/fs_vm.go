package vm

import (
	"zzc/fall-script/src/builtin"
	"zzc/fall-script/src/compiler"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

const (
	MAX_FRAMS   = 1024
	GLOBAL_SIZE = 65535
)

type FsVM struct {
	builtins   []*builtin.BuiltinDef
	globals    []object.Object
	mainThread *rt.Thread
}

func NewFsVM(c *compiler.Compiler) *FsVM {
	locals := c.SymbolTable.MaxNum()
	insts := c.CurrentInstructions()
	consts := c.CurrentConstant()
	stackDepth := c.CurrentStackDepth()
	globals := make([]object.Object, GLOBAL_SIZE)

	compileFn := &object.CompiledFunction{
		Instructions: insts,
		Constants:    consts,
		LocalsNum:    locals,
		ParamsNum:    0,
		StackDepth:   stackDepth,
	}

	mainClosure := &object.Closure{
		Fn:   compileFn,
		Free: []object.Object{},
	}

	fv := &FsVM{
		globals:  globals,
		builtins: c.Builtins,
	}

	mainThread := fv.NewThread()
	mainFrame := mainThread.NewFrame(mainClosure)
	mainThread.PushFrame(mainFrame)

	fv.mainThread = mainThread
	return fv
}

func (fv *FsVM) NewThread() *rt.Thread {
	return rt.NewThread(MAX_FRAMS, fv.globals, fv.builtins)
}

func (fv *FsVM) Run() {
	interpreter(fv.mainThread)
}

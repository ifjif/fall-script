package vm

import (
	"path/filepath"

	"zzc/fall-script/src/builtin"
	"zzc/fall-script/src/builtin/vmb"
	"zzc/fall-script/src/module"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

const (
	MAX_FRAMS = 1024
)

type FsVM struct {
	builtins       []*builtin.BuiltinDef
	mainThread     *rt.Thread
	loader         *module.Loader
	moduleRegister module.ModuleRegister
}

func newFsVM() *FsVM {
	loader := module.NewLoader()
	moduleRegister := module.ModuleRegister{}
	fv := &FsVM{
		builtins:       vmb.Builtins,
		loader:         loader,
		moduleRegister: moduleRegister,
	}

	return fv
}

func NewFsVMWithFile(file string) *FsVM {
	fp := module.ResolveImportPath(".", file)
	startDir := filepath.Dir(fp)
	fv := newFsVM()
	module := fv.loader.LoadFile(fp)
	fv.initMain(module, startDir)
	return fv
}

func NewFsVMWithText(input []byte) *FsVM {
	fv := newFsVM()
	module := fv.loader.LoadText(input, ".")
	fv.initMain(module, ".")
	return fv
}

func (fv *FsVM) initMain(mo *object.Module, dir string) {
	cmo := fv.load(mo)
	fv.link(mo, cmo, dir)
	mainClosure := cmo.Closure()
	mainThread := fv.NewThread()
	mainFrame := mainThread.NewFrame(mainClosure)
	mainThread.PushFrame(mainFrame)
	cmo.Status = object.Initializing

	fv.mainThread = mainThread
}

func (fv *FsVM) NewThread() *rt.Thread {
	return rt.NewThread(MAX_FRAMS, fv.builtins)
}

func (fv *FsVM) Run() {
	interpreter(fv.mainThread)
}

func (fv *FsVM) load(mo *object.Module) *object.CompiledModule {
	cf := mo.Cf
	globals := make([]object.Object, mo.GlobalNum)

	consts := cf.Constants
	exports := make(map[string]int)
	for _, eref := range mo.Exports {
		nidx := consts[eref.Name]
		if nidx.Type() == object.STRING_OBJ {
			str := nidx.(*object.String)
			exports[str.Value] = eref.GlobalId
		}
	}

	module := &object.CompiledModule{
		Globals: globals,
		Cf:      cf,
		Exports: exports,
		Status:  object.Uninitialized,
	}
	// 为所有 struct设置 module
	for _, st := range mo.Structs {
		st.OwnerModule = module
	}
	fv.moduleRegister[mo.Name] = module

	return module
}

func (fv *FsVM) link(mo *object.Module, cmo *object.CompiledModule, dir string) {
	consts := mo.Cf.Constants
	for i, iref := range mo.Imports {
		relativePath := consts[iref.From].Inspect()
		source := module.ResolveImportPath(filepath.Join(dir, mo.Name), relativePath)
		target, ok := fv.moduleRegister[source]
		if !ok {
			nmod := fv.loader.LoadFile(source)
			target = fv.load(nmod)
			fv.link(nmod, target, dir)
		}

		imported := consts[iref.Imported].Inspect()
		targetIdx := target.Exports[imported]
		globalRef := &object.GlobalRef{
			TargetModule: target,
			TargetIdx:    targetIdx,
		}
		cmo.Globals[i] = globalRef
	}
	// todo struct promoted
}

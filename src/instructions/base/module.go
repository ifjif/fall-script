package base

import (
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

// 通过 bool 判断是否在初始
// 如果 在外部使用
// targetModule.Status = object.Initializing
// 需要在外部额外判断是否 运行的是当前模块
// 即： frame.Closure().OwnerModule != targetModule
// 模块的 顶部指令 都是 设置 structMeta信息
func InitModule(frame *rt.Frame, targetModule *object.CompiledModule) (*object.CompiledModule, error) {
	if targetModule.Status == object.Uninitialized {
		closure := targetModule.Closure()
		nframe := frame.Thread().NewFrame(closure)
		frame.Thread().PushFrame(nframe)
		targetModule.Status = object.Initializing
		// 恢复pc
		frame.SetNextPc(frame.Thread().GetPc())
		return targetModule, ErrAwaitModuleInitialization
	}
	return targetModule, nil
}

package base

import (
	"fmt"

	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

func GetTargetStructMeta(frame *rt.Frame, curStructMeta *object.StructMeta, targetStructMeta int) (*object.StructMeta, error) {
	// 从 struct 所在的 module找它的 组合的struct
	ownerModule, awaitInit := InitModule(frame, curStructMeta.OwnerModule)
	if awaitInit != nil {
		return nil, awaitInit
	}
	meta := ownerModule.Globals[targetStructMeta]
	if globalRef, ok := meta.(*object.GlobalRef); ok {
		targetModule, awaitInit := InitModule(frame, globalRef.TargetModule)
		if awaitInit != nil {
			return nil, awaitInit
		}
		meta = targetModule.Globals[globalRef.TargetIdx]
	}

	target, ok := meta.(*object.StructMeta)
	if !ok {
		return nil, fmt.Errorf("expected type is %s, got %s", object.STRUCT_META_OBJ, meta.Type())
	}

	return target, nil
}

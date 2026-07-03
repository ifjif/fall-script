package object

type ObjectType string

const (
	BYTE_OBJ              = "BYTE"
	INTEGER_OBJ           = "INTEGER"
	BOOLEAN_OBJ           = "BOOLEAN"
	STRING_OBJ            = "STRING"
	ARRAY_OBJ             = "ARRAY"
	HASH_OBJ              = "HASH"
	BOX_OBJ               = "BOX"
	FUNCTION_OBJ          = "FUNCTION"
	COMPILED_FUNCTION_OBJ = "COMPILED_FUNCTION"
	BUILTIN_FUNCTION_OBJ  = "BUILTIN_FUNCTION"
	CLOSURE_OBJ           = "CLOSURE"
	RETURN_OBJ            = "RETURN"
	NULL_OBJ              = "NULL"
	ERROR_OBJ             = "ERROR"
	MACRO_OBJ             = "MACRO"
	QUOTE_OBJ             = "QUOTE"
	MODULE_OBJ            = "MODULE"
	COMPILED_MODULE_OBJ   = "COMPILED_MODULE"
	IMPORT_REF_OBJ        = "IMPORT_REF"
	EXPORT_REF_OBJ        = "EXPORT_REF"
	GLOBAL_REF_OBJ        = "GLOBAL_REF"
)

const SliceOmitted = -999999

type BuiltinFunction func(args ...Object) Object

type Object interface {
	Type() ObjectType
	Inspect() string
}

type HashTableKey interface {
	HashKey() HashKey
}

type Sliceable interface {
	Object
	Len() int
	Cap() int
	Slice(start, end, capc int) Sliceable
	SliceCopy(start, end, step int) Sliceable
}

func calcNewLen(start, end, step int) int {
	if step > 0 {
		return (end - start + step - 1) / step
	}

	return (start - end - step - 1) / -step
}

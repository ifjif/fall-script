package object

type ObjectType string

const (
	INTEGER_OBJ           = "INTEGER"
	BOOLEAN_OBJ           = "BOOLEAN"
	STRING_OBJ            = "STRING"
	ARRAY_OBJ             = "ARRAY"
	HASH_OBJ              = "HASH"
	FUNCTION_OBJ          = "FUNCTION"
	COMPILED_FUNCTION_OBJ = "COMPILED_FUNCTION"
	BUILTIN_FUNCTION_OBJ  = "BUILTIN_FUNCTION"
	CLOSURE_OBJ           = "CLOSURE"
	RETURN_OBJ            = "RETURN"
	NULL_OBJ              = "NULL"
	ERROR_OBJ             = "ERROR"
)

type BuiltinFunction func(args ...Object) Object

type Object interface {
	Type() ObjectType
	Inspect() string
}

type HashTableKey interface {
	HashKey() HashKey
}

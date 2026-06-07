package object

type ObjectType string

const (
	INTEGER_OBJ  = "INTEGER"
	BOOLEAN_OBJ  = "BOOLEAN"
	STRING_OBJ   = "STRING"
	ARRAY_OBJ    = "ARRAY"
	HASH_OBJ     = "HASH"
	FUNCTION_OBJ = "FUNCTION"
	RETURN_OBJ   = "RETURN"
	NULL_OBJ     = "NULL"
	ERROR_OBJ    = "ERROR"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type HashTableKey interface {
	HashKey() HashKey
}

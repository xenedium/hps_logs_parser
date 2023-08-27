package types

type field struct {
	Length int64
	Value  string
	Raw    []byte
}

type Field = field

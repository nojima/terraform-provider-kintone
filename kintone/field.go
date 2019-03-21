package kintone

type FieldType string

func (ft FieldType) String() string {
	return string(ft)
}

type FieldCode string

func (fc FieldCode) String() string {
	return string(fc)
}

type Field interface {
	Type() FieldType
	Code() FieldCode
	Label() string
}

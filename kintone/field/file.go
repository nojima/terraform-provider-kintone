package field

import "github.com/naruta/terraform-provider-kintone/kintone"

type File struct {
	code  kintone.FieldCode
	label string
}

func NewFile(code kintone.FieldCode, label string) *File {
	return &File{
		code:  code,
		label: label,
	}
}

func (s *File) Type() kintone.FieldType {
	return "FILE"
}

func (s *File) Code() kintone.FieldCode {
	return s.code
}

func (s *File) Label() string {
	return s.label
}

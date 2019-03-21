package field

import "github.com/naruta/terraform-provider-kintone/kintone"

type Date struct {
	code  kintone.FieldCode
	label string
}

func NewDate(code kintone.FieldCode, label string) *Date {
	return &Date{
		code:  code,
		label: label,
	}
}

func (s *Date) Type() kintone.FieldType {
	return "DATE"
}

func (s *Date) Code() kintone.FieldCode {
	return s.code
}

func (s *Date) Label() string {
	return s.label
}

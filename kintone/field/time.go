package field

import "github.com/naruta/terraform-provider-kintone/kintone"

type Time struct {
	code  kintone.FieldCode
	label string
}

func NewTime(code kintone.FieldCode, label string) *Time {
	return &Time{
		code:  code,
		label: label,
	}
}

func (s *Time) Type() kintone.FieldType {
	return "TIME"
}

func (s *Time) Code() kintone.FieldCode {
	return s.code
}

func (s *Time) Label() string {
	return s.label
}

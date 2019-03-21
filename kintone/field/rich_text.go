package field

import "github.com/naruta/terraform-provider-kintone/kintone"

type RichText struct {
	code  kintone.FieldCode
	label string
}

func NewRichText(code kintone.FieldCode, label string) *RichText {
	return &RichText{
		code:  code,
		label: label,
	}
}

func (s *RichText) Type() kintone.FieldType {
	return "RICH_TEXT"
}

func (s *RichText) Code() kintone.FieldCode {
	return s.code
}

func (s *RichText) Label() string {
	return s.label
}

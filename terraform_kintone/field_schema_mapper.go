package terraform_kintone

import (
	"github.com/naruta/terraform-provider-kintone/kintone"
	"github.com/naruta/terraform-provider-kintone/kintone/field"
	"github.com/pkg/errors"
)

type fieldFactory func(code kintone.FieldCode, label string, fieldMap map[string]interface{}) (kintone.Field, error)

var fieldFactories = map[string]fieldFactory {
	"SINGLE_LINE_TEXT": func(code kintone.FieldCode, label string, fieldMap map[string]interface{}) (kintone.Field, error) {
		return field.NewSingleLineText(code, label), nil
	},
	"MULTI_LINE_TEXT": func(code kintone.FieldCode, label string, fieldMap map[string]interface{}) (kintone.Field, error) {
		return field.NewMultiLineText(code, label), nil
	},
	"RICH_TEXT": func(code kintone.FieldCode, label string, fieldMap map[string]interface{}) (kintone.Field, error) {
		return field.NewRichText(code, label), nil
	},
	"NUMBER": func(code kintone.FieldCode, label string, fieldMap map[string]interface{}) (kintone.Field, error) {
		return field.NewNumber(code, label), nil
	},
	"DATE": func(code kintone.FieldCode, label string, fieldMap map[string]interface{}) (kintone.Field, error) {
		return field.NewDate(code, label), nil
	},
}

func validFieldTypes() []string {
	var fieldTypes []string
	for fieldType := range fieldFactories {
		fieldTypes = append(fieldTypes, fieldType)
	}
	return fieldTypes
}

type fieldSchemaMapper struct{}

func (m *fieldSchemaMapper) SchemaToField(fieldMap map[string]interface{}) (kintone.Field, error) {
	fieldType := fieldMap["type"].(string)
	code := kintone.FieldCode(fieldMap["code"].(string))
	label := fieldMap["label"].(string)

	factory, ok := fieldFactories[fieldType]
	if !ok {
		return nil, errors.Errorf("unknown field type: %s", fieldType)
	}

	return factory(code, label, fieldMap)
}

func (m *fieldSchemaMapper) FieldToSchema(f kintone.Field) map[string]interface{} {
	return map[string]interface{}{
		"code":  f.Code().String(),
		"label": f.Label(),
		"type":  f.Type().String(),
	}
}

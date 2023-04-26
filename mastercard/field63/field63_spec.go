package field63

import (
	"reflect"

	"github.com/moov-io/iso8583/field"
)

type Field63Spec struct {
	Name   string
	Fields map[int]field.Field
}

// Creates a map with new instances of Fields (Field interface)
// based on the field type in Field54Spec.
func (s *Field63Spec) CreateField54Fields() map[int]field.Field {

	fields := map[int]field.Field{}

	for k, specField := range s.Fields {
		fields[k] = createField63(specField)
	}

	return fields
}

func createField63(specField field.Field) field.Field {
	fieldType := reflect.TypeOf(specField).Elem()

	// create new field and convert it to field.Field interface
	fl := reflect.New(fieldType).Interface().(field.Field)
	fl.SetSpec(specField.Spec())

	return fl
}

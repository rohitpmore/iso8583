package field63

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strconv"

	"github.com/moov-io/iso8583/field"
)

var _ json.Marshaler = (*Field63)(nil)
var _ json.Unmarshaler = (*Field63)(nil)

const (
	bitmapIdx = 0
)

type Field63 struct {
	spec      *Field63Spec
	data      interface{}
	dataValue *reflect.Value

	// stores all fields according to the spec
	fields map[int]field.Field
	// tracks which fields were set
	fieldsMap map[int]struct{}
}

func NewField63(spec *Field63Spec) *Field63 {
	fields := spec.CreateField54Fields()

	return &Field63{
		fields:    fields,
		spec:      spec,
		fieldsMap: map[int]struct{}{},
	}
}

func (m *Field63) Data() interface{} {
	return m.data
}

func (m *Field63) SetData(data interface{}) error {
	m.data = data

	if m.data == nil {
		return nil
	}

	dataStruct := reflect.ValueOf(data)
	if dataStruct.Kind() == reflect.Ptr || dataStruct.Kind() == reflect.Interface {
		// get the struct
		dataStruct = dataStruct.Elem()
	}

	if reflect.TypeOf(dataStruct).Kind() != reflect.Struct {
		return fmt.Errorf("failed to set data as struct is expected, got: %s", reflect.TypeOf(dataStruct).Kind())
	}

	m.dataValue = &dataStruct
	return nil
}

func (m *Field63) GetSpec() *Field63Spec {
	return m.spec
}

func (m *Field63) Field(id int, val string) error {
	if f, ok := m.fields[id]; ok {
		m.fieldsMap[id] = struct{}{}
		return f.SetBytes([]byte(val))
	}
	return fmt.Errorf("failed to set field %d. ID does not exist", id)
}

func (m *Field63) BinaryField(id int, val []byte) error {
	if f, ok := m.fields[id]; ok {
		m.fieldsMap[id] = struct{}{}
		return f.SetBytes(val)
	}
	return fmt.Errorf("failed to set binary field %d. ID does not exist", id)
}

func (m *Field63) GetString(id int) (string, error) {
	if f, ok := m.fields[id]; ok {
		m.fieldsMap[id] = struct{}{}
		return f.String()
	}
	return "", fmt.Errorf("failed to get string for field %d. ID does not exist", id)
}

func (m *Field63) GetBytes(id int) ([]byte, error) {
	if f, ok := m.fields[id]; ok {
		m.fieldsMap[id] = struct{}{}
		return f.Bytes()
	}
	return nil, fmt.Errorf("failed to get bytes for field %d. ID does not exist", id)
}

func (m *Field63) GetField(id int) field.Field {
	return m.fields[id]
}

// Fields returns the map of the set fields
func (m *Field63) GetFields() map[int]field.Field {
	fields := map[int]field.Field{}
	for i := range m.fieldsMap {
		fields[i] = m.GetField(i)
	}
	return fields
}

func (m *Field63) Pack() ([]byte, error) {
	packed := []byte{}

	ids, err := m.setPackableDataFields()
	if err != nil {
		return nil, fmt.Errorf("failed to pack message: %w", err)
	}

	// pack fields
	for _, i := range ids {
		field, ok := m.fields[i]
		if !ok {
			return nil, fmt.Errorf("failed to pack field %d: no specification found", i)
		}
		packedField, err := field.Pack()
		if err != nil {
			return nil, fmt.Errorf("failed to pack field %d (%s): %w", i, field.Spec().Description, err)
		}
		packed = append(packed, packedField...)
	}

	return packed, nil
}

func (m *Field63) Unpack(src []byte) error {
	var off int

	off = 0
	fmt.Println("src: ", src, " string src:", string(src))
	for off < len(src) {
		for i := 1; i <= 3; i++ {
			fl, ok := m.fields[i]
			if !ok {
				return fmt.Errorf("failed to unpack field %d: no specification found", i)
			}

			if m.dataValue != nil {
				if err := m.setUnpackableDataField(i); err != nil {
					return err
				}
			}

			m.fieldsMap[i] = struct{}{}
			read, err := fl.Unpack(src[off:])
			if err != nil {
				return fmt.Errorf("failed to unpack field %d (%s): %w", i, fl.Spec().Description, err)
			}

			flValue, _ := fl.String()
			flBytes, _ := fl.Bytes()
			fmt.Printf("Field 54.%v - %v - %v - %v\n", i, fl.Spec().Description, flValue, flBytes)
			off += read

		}
	}

	return nil
}

func (m *Field63) MarshalJSON() ([]byte, error) {
	// by packing the message we will generate bitmap
	// create HEX representation
	// and validate message against the spec
	if _, err := m.Pack(); err != nil {
		return nil, err
	}

	fieldMap := m.GetFields()
	strFieldMap := map[string]field.Field{}
	for k, v := range fieldMap {
		// we don't wish to populate the bitmap in the final
		// JSON since it is dynamically generated when packing
		// and unpacking anyways.
		if k == bitmapIdx {
			continue
		}
		strFieldMap[fmt.Sprint(k)] = v
	}

	// get only fields that were set
	return json.Marshal(field.OrderedMap(strFieldMap))
}

func (m *Field63) UnmarshalJSON(b []byte) error {
	var data map[string]json.RawMessage
	json.Unmarshal(b, &data)

	for id, rawMsg := range data {
		i, err := strconv.Atoi(id)
		if err != nil {
			return fmt.Errorf("failed to unmarshal field %v: could not convert to int", i)
		}

		field, ok := m.fields[i]
		if !ok {
			return fmt.Errorf("failed to unmarshal field %d: no specification found", i)
		}

		if m.dataValue != nil {
			if err := m.setUnpackableDataField(i); err != nil {
				return err
			}
		}

		m.fieldsMap[i] = struct{}{}
		if err := json.Unmarshal(rawMsg, field); err != nil {
			return fmt.Errorf("failed to unmarshal field %v: %w", id, err)
		}
	}

	return nil
}

func (m *Field63) setPackableDataFields() ([]int, error) {
	// Index  1 represent bitmap.
	// It is assumed to be always populated.
	populatedFieldIDs := []int{1}

	for id, field := range m.fields {
		// represents the bitmap
		if id == 1 {
			continue
		}

		// These fields are set using the typed API
		if m.dataValue != nil {
			dataField := m.dataFieldValue(id)
			// no non-nil data field was found with this name
			if dataField == (reflect.Value{}) || dataField.IsNil() {
				continue
			}
			if err := field.SetData(dataField.Interface()); err != nil {
				return nil, fmt.Errorf("failed to set data for field %d: %w", id, err)
			}

			// mark field as set
			m.fieldsMap[id] = struct{}{}
		}

		// These fields are set using the untyped API
		_, ok := m.fieldsMap[id]
		// We don't wish set the MTI again, hence we ignore the 0
		// index
		if (ok || m.dataValue != nil) && id != 0 {
			populatedFieldIDs = append(populatedFieldIDs, id)
		}
	}
	sort.Ints(populatedFieldIDs)

	return populatedFieldIDs, nil
}

func (m *Field63) setUnpackableDataField(id int) error {
	specField, ok := m.fields[id]
	if !ok {
		return fmt.Errorf("failed to unpack field %d: no specification found", id)
	}

	dataField := m.dataFieldValue(id)
	// no data field was found with this name
	if dataField == (reflect.Value{}) {
		return nil
	}

	isNil := dataField.IsNil()
	if isNil {
		dataField.Set(reflect.New(dataField.Type().Elem()))
	}
	if err := specField.SetData(dataField.Interface()); err != nil {
		return fmt.Errorf("failed to set data for field %d: %w", id, err)
	}

	return nil
}

func (m *Field63) dataFieldValue(id int) reflect.Value {
	return m.dataValue.FieldByName(fmt.Sprintf("F%d", id))
}

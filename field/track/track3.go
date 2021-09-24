package track

import (
	"errors"
	"fmt"
	"strings"

	"github.com/moov-io/iso8583/field"
)

var _ field.Field = (*Track3)(nil)

type Track3 struct {
	spec                 *field.Spec `json:"-"`
	FormatCode           string      `json:"format_code,omitempty"`
	PrimaryAccountNumber string      `json:"primary_account_number,omitempty"`
	DiscretionaryData    string      `json:"discretionary_data,omitempty"`
}

func NewTrack3(spec *field.Spec) (*Track3, error) {
	return &Track3{
		spec: spec,
	}, nil
}

func NewTrack3Value(val []byte, fixedLength bool) (*Track3, error) {
	track := &Track3{}
	err := track.parse(val)
	if err != nil {
		return nil, errors.New("invalid track data")
	}
	return track, nil
}

func (f *Track3) Spec() *field.Spec {
	return f.spec
}

func (f *Track3) SetSpec(spec *field.Spec) {
	f.spec = spec
}

func (f *Track3) SetBytes(b []byte) error {
	if err := f.parse(b); err != nil {
		return nil
	}
	return nil
}

func (f *Track3) Bytes() ([]byte, error) {
	return f.serialize()
}

func (f *Track3) String() (string, error) {
	b, err := f.serialize()
	if err != nil {
		return "", fmt.Errorf("failed to encode string: %v", err)
	}
	return string(b), nil
}

func (f *Track3) Pack() ([]byte, error) {
	data, err := f.serialize()
	if err != nil {
		return nil, err
	}

	if f.spec.Pad != nil {
		data = f.spec.Pad.Pad(data, f.spec.Length)
	}

	packed, err := f.spec.Enc.Encode(data)
	if err != nil {
		return nil, fmt.Errorf("failed to encode content: %v", err)
	}

	packedLength, err := f.spec.Pref.EncodeLength(f.spec.Length, len(packed))
	if err != nil {
		return nil, fmt.Errorf("failed to encode length: %v", err)
	}

	return append(packedLength, packed...), nil
}

// returns number of bytes was read
func (f *Track3) Unpack(data []byte) (int, error) {
	dataLen, prefBytes, err := f.spec.Pref.DecodeLength(f.spec.Length, data)
	if err != nil {
		return 0, fmt.Errorf("failed to decode length: %v", err)
	}

	raw, read, err := f.spec.Enc.Decode(data[prefBytes:], dataLen)
	if err != nil {
		return 0, fmt.Errorf("failed to decode content: %v", err)
	}

	if f.spec.Pad != nil {
		raw = f.spec.Pad.Unpad(raw)
	}

	if len(raw) > 0 {
		err = f.parse(raw)
		if err != nil {
			return 0, err
		}
	}

	return read + prefBytes, nil
}

func (f *Track3) SetData(data interface{}) error {
	if data == nil {
		return nil
	}

	track, ok := data.(*Track3)
	if !ok {
		return fmt.Errorf("data does not match required *Track type")
	}

	f.FormatCode = track.FormatCode
	f.PrimaryAccountNumber = track.PrimaryAccountNumber
	f.DiscretionaryData = track.DiscretionaryData

	return nil
}

func (f *Track3) parse(raw []byte) error {
	if raw == nil {
		return errors.New("invalid track data")
	}

	matches := track3Regex.FindStringSubmatch(string(raw))
	for index, val := range matches {
		value := strings.TrimSpace(val)
		if len(value) == 0 || value == "=" {
			continue
		}

		switch index {
		case 1: // Format Code
			f.FormatCode = value
		case 2: // Payment card number (PAN)
			f.PrimaryAccountNumber = value
		case 3: // Security Data + Additional Data
			f.DiscretionaryData = value
		}
	}

	return nil
}

func (f *Track3) serialize() ([]byte, error) {
	raw := fmt.Sprintf(track3Format, f.FormatCode, f.PrimaryAccountNumber, f.DiscretionaryData)
	return []byte(raw), nil
}

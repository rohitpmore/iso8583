package track

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/moov-io/iso8583/field"
)

var _ field.Field = (*Track1)(nil)

type Track1 struct {
	spec                 *field.Spec `json:"-"`
	FixedLength          bool        `json:"fixed_length,omitempty"`
	FormatCode           string      `json:"format_code,omitempty"`
	PrimaryAccountNumber string      `json:"primary_account_number,omitempty"`
	Name                 string      `json:"name,omitempty"`
	ExpirationDate       *time.Time  `json:"expiration_date,omitempty"`
	ServiceCode          string      `json:"service_code,omitempty"`
	DiscretionaryData    string      `json:"discretionary_data,omitempty"`
}

const (
	expiryDateFormat = "0601"
	track1Format     = `%s%s^%s^%s%s%s`
	track2Format     = `%s=%s%s%s`
	track3Format     = `%s%s=%s`
)

var (
	track1Regex = regexp.MustCompile(`^([A-Z]{1})([0-9]{1,19})\^([^\^]{2,26})\^([0-9]{4}|\^)([0-9]{3}|\^)([^\?]+)$`)
	track2Regex = regexp.MustCompile(`^([0-9]{1,19})\=([0-9]{4}|\=)([0-9]{3}|\=)([^\?]+)$`)
	track3Regex = regexp.MustCompile(`^([0-9]{2})([0-9]{1,19})\=([^\?]+)$`)
)

func NewTrack1(spec *field.Spec) (*Track1, error) {
	return &Track1{
		spec: spec,
	}, nil
}

func NewTrack1Value(val []byte, fixedLength bool) (*Track1, error) {
	track := &Track1{
		FixedLength: fixedLength,
	}
	err := track.parse(val)
	if err != nil {
		return nil, errors.New("invalid track data")
	}
	return track, nil
}

func (f *Track1) Spec() *field.Spec {
	return f.spec
}

func (f *Track1) SetSpec(spec *field.Spec) {
	f.spec = spec
}

func (f *Track1) SetBytes(b []byte) error {
	return f.parse(b)
}

func (f *Track1) Bytes() ([]byte, error) {
	return f.serialize()
}

func (f *Track1) String() (string, error) {
	b, err := f.serialize()
	if err != nil {
		return "", fmt.Errorf("failed to encode string: %v", err)
	}
	return string(b), nil
}

func (f *Track1) Pack() ([]byte, error) {
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
func (f *Track1) Unpack(data []byte) (int, error) {
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

func (f *Track1) SetData(data interface{}) error {
	if data == nil {
		return nil
	}

	track, ok := data.(*Track1)
	if !ok {
		return fmt.Errorf("data does not match required *Track type")
	}

	f.FixedLength = track.FixedLength
	f.FormatCode = track.FormatCode
	f.PrimaryAccountNumber = track.PrimaryAccountNumber
	f.Name = track.Name
	f.ExpirationDate = track.ExpirationDate
	f.ServiceCode = track.ServiceCode
	f.DiscretionaryData = track.DiscretionaryData

	return nil
}

func (f *Track1) parse(raw []byte) error {
	if raw == nil {
		return errors.New("invalid track data")
	}

	matches := track1Regex.FindStringSubmatch(string(raw))
	for index, val := range matches {
		value := strings.TrimSpace(val)
		if len(value) == 0 || value == "^" {
			continue
		}

		switch index {
		case 1: // Format Code
			f.FormatCode = value
		case 2: // Payment card number (PAN)
			f.PrimaryAccountNumber = value
		case 3: // Name (NM)
			f.Name = value
		case 4: // Expiration Date (ED)
			expiredTime, timeErr := time.Parse(expiryDateFormat, value)
			if timeErr != nil {
				return errors.New("invalid expired time")
			}
			f.ExpirationDate = &expiredTime
		case 5: // Service Code (SC)
			f.ServiceCode = value
		case 6: // Discretionary data (DD)
			f.DiscretionaryData = value
		}
	}

	return nil
}

func (f *Track1) serialize() ([]byte, error) {

	name := f.Name
	if len(f.Name) > 1 && f.FixedLength {
		name = fmt.Sprintf("%-26.26s", f.Name)
	}
	expired := "^"
	if f.ExpirationDate != nil {
		expired = f.ExpirationDate.Format(expiryDateFormat)
	}
	code := "^"
	if len(f.ServiceCode) > 0 {
		code = f.ServiceCode
	}

	raw := fmt.Sprintf(track1Format, f.FormatCode, f.PrimaryAccountNumber, name, expired, code, f.DiscretionaryData)
	return []byte(raw), nil
}
package field62

import (
	"fmt"

	"github.com/moov-io/iso8583/field"
	"github.com/moov-io/iso8583/utils"
)

const bitmapLength = 8 // 64 bit, 8 bytes, or 16 hex digits

var _ field.Field = (*F62Bitmap)(nil)

// NOTE: Bitmap does not support JSON encoding or decoding.
type F62Bitmap struct {
	spec   *field.Spec
	bitmap *utils.Bitmap
	data   *F62Bitmap
}

func NewBitmap(spec *field.Spec) *F62Bitmap {
	return &F62Bitmap{
		spec:   spec,
		bitmap: utils.NewBitmap(bitmapLength),
	}
}

func (f *F62Bitmap) Spec() *field.Spec {
	return f.spec
}

func (f *F62Bitmap) SetSpec(spec *field.Spec) {
	f.spec = spec
}

func (f *F62Bitmap) SetBytes(b []byte) error {
	f.bitmap = utils.NewBitmapFromData(b)
	if f.data != nil {
		*(f.data) = *f
	}
	return nil
}

func (f *F62Bitmap) Bytes() ([]byte, error) {
	return f.bitmap.Bytes(), nil
}

func (f *F62Bitmap) String() (string, error) {
	return f.bitmap.String(), nil
}

func (f *F62Bitmap) Pack() ([]byte, error) {
	// here we have max possible bytes for the bitmap 8*maxBitmaps
	data, err := f.Bytes()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve bytes: %v", err)
	}

	// data = data[0 : 8*count]

	packed, err := f.spec.Enc.Encode(data)
	if err != nil {
		return nil, fmt.Errorf("failed to encode content: %v", err)
	}

	return packed, nil
}

// Unpack of the Bitmap field returns data of varied length
// if there is only primary bitmap (bit 1 is not set) we return only 8 bytes (or 16 for hex encoding)
// if secondary bitmap presents (bit 1 is set) we return 16 bytes (or 32 for hex encoding)
// and so on for maxBitmaps
func (f *F62Bitmap) Unpack(data []byte) (int, error) {
	minLen, _, err := f.spec.Pref.DecodeLength(bitmapLength, data)
	if err != nil {
		return 0, fmt.Errorf("failed to decode length: %v", err)
	}

	rawBitmap := make([]byte, 0)
	read := 0

	decoded, readDecoded, err := f.spec.Enc.Decode(data[read:], minLen)
	if err != nil {
		return 0, fmt.Errorf("failed to decode content for bitmap: %v", err)
	}
	read += readDecoded

	rawBitmap = append(rawBitmap, decoded...)

	if err := f.SetBytes(rawBitmap); err != nil {
		return 0, fmt.Errorf("failed to set bytes: %w", err)
	}

	return read, nil
}

func (f *F62Bitmap) SetData(data interface{}) error {
	if data == nil {
		return nil
	}

	bmap, ok := data.(*F62Bitmap)
	if !ok {
		return fmt.Errorf("data does not match required *Bitmap type")
	}

	f.data = bmap
	if bmap.bitmap != nil {
		f.bitmap = bmap.bitmap
	}
	return nil
}

func (f *F62Bitmap) Reset() {
	f.bitmap = utils.NewBitmap(bitmapLength)
}

func (f *F62Bitmap) Set(i int) {
	f.bitmap.Set(i)
}

func (f *F62Bitmap) IsSet(i int) bool {
	return f.bitmap.IsSet(i)
}

func (f *F62Bitmap) Len() int {
	return f.bitmap.Len()
}

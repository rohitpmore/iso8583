package prefix

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinaryVarPrefixer_EncodeLengthDigitsValidation(t *testing.T) {
	_, err := Binary.LL.EncodeLength(999, 123)

	require.Contains(t, err.Error(), "number of digits in length: 123 exceeds: 2")
}

func TestBinaryVarPrefixer_EncodeLengthMaxLengthValidation(t *testing.T) {
	_, err := Binary.LL.EncodeLength(20, 22)

	require.Contains(t, err.Error(), "field length: 22 is larger than maximum: 20")
}

func TestBinaryVarPrefixer_DecodeLengthMaxLengthValidation(t *testing.T) {
	_, _, err := Binary.LLL.DecodeLength(20, []byte{0b100010})

	require.Contains(t, err.Error(), "length mismatch: want to read 3 bytes, get only 1")
}

func TestBinaryVarPrefixer_LHelpers(t *testing.T) {
	tests := []struct {
		pref      Prefixer
		bytesRead int
		maxLen    int
		in        int
		out       []byte
	}{
		{Binary.L, 1, 5, 3, []byte{0b11}},
		{Binary.L, 1, 20, 2, []byte{0b10}},
		{Binary.L, 1, 20, 12, []byte{0b1100}},
		{Binary.LL, 2, 340, 2, []byte{0b0, 0b10}},
		{Binary.LL, 2, 340, 200, []byte{0b0, 0b11001000}},
		{Binary.LL, 2, 9999, 1234, []byte{0b00000100, 0b11010010}},
	}

	// test encoding
	for _, tt := range tests {
		t.Run(tt.pref.Inspect()+"_EncodeLength", func(t *testing.T) {
			got, err := tt.pref.EncodeLength(tt.maxLen, tt.in)
			require.NoError(t, err)
			require.Equal(t, tt.out, got)
		})
	}

	// test decoding
	for _, tt := range tests {
		t.Run(tt.pref.Inspect()+"_DecodeLength", func(t *testing.T) {
			got, read, err := tt.pref.DecodeLength(tt.maxLen, tt.out)
			require.NoError(t, err)
			require.Equal(t, tt.in, got)
			require.Equal(t, tt.bytesRead, read)
		})
	}
}

func TestBinaryFixedPrefixer(t *testing.T) {
	pref := binaryFixedPrefixer{}

	// Fixed prefixer returns empty byte slice as
	// size is not encoded into field
	data, err := pref.EncodeLength(8, 8)

	require.NoError(t, err)
	require.Equal(t, 0, len(data))

	// Fixed prefixer returns configured len
	// rather than read it from data
	dataLen, read, err := pref.DecodeLength(8, []byte("1234"))

	require.NoError(t, err)
	require.Equal(t, 8, dataLen)
	require.Equal(t, 0, read)
}

func TestBinaryFixedPrefixer_EncodeLengthValidation(t *testing.T) {
	pref := binaryFixedPrefixer{}

	_, err := pref.EncodeLength(8, 12)

	require.Contains(t, err.Error(), "field length: 12 should be fixed: 8")
}

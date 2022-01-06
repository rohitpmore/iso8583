package encoding

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEBCDIC(t *testing.T) {
	t.Run("Decode", func(t *testing.T) {
		res, read, err := EBCDIC.Decode([]byte{0x12, 0x34}, 2)

		require.NoError(t, err)
		require.Equal(t, []byte{0x12, 0x94}, res)
		require.Equal(t, 2, read)

	})

	t.Run("Encode", func(t *testing.T) {
		res, err := EBCDIC.Encode([]byte{0x12, 0x94})

		require.NoError(t, err)
		require.Equal(t, []byte{0x12, 0x34}, res)

	})
}

func TestMSDI(t *testing.T) {
	t.Run("Decode", func(t *testing.T) {
		res, read, err := EBCDIC.Decode([]byte{0b11010101}, 1)
		fmt.Println("MSDI: ", string(res))

		require.NoError(t, err)
		require.Equal(t, 1, read)

	})
}

func TestMC_F48(t *testing.T) {
	t.Run("Decode", func(t *testing.T) {
		res, read, err := EBCDIC.Decode([]byte{0b00110110, 0b00110001}, 2)
		fmt.Println("F48.61 : ", string(res))

		require.NoError(t, err)
		require.Equal(t, 2, read)

	})
}
func TestField54_WG_70067(t *testing.T) {
	t.Run("Decode", func(t *testing.T) {
		res, read, err := EBCDIC.Decode([]byte{0b11110000, 0b11110000}, 2)
		fmt.Println("Account Type: ", string(res))
		res, read, err = EBCDIC.Decode([]byte{0b11110100, 0b11100010}, 2)
		fmt.Println("Amount Type: ", string(res))
		res, read, err = EBCDIC.Decode([]byte{0b11111000, 0b11110100, 0b11110000}, 3)
		fmt.Println("Currency Code: ", string(res))
		res, read, err = EBCDIC.Decode([]byte{0b11000011}, 1)
		fmt.Println("Amount Sign: ", string(res))
		res, read, err = EBCDIC.Decode([]byte{0b11110000, 0b11110000, 0b11110000, 0b11110000, 0b11110000, 0b11110000, 0b11110000, 0b11110000, 0b11110000, 0b11110100, 0b11110001, 0b11110010}, 12)
		fmt.Println("Amount: ", string(res))

		require.NoError(t, err)
		require.Equal(t, 1, read)
	})
}

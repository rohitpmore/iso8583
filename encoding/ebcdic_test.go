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

package field48_test

import (
	"fmt"
	"testing"

	"github.com/moov-io/iso8583/mastercard/field48"
	"github.com/moov-io/iso8583/specs"
)

func Test_Field48(t *testing.T) {
	fmt.Println("Welcome to the testing of field48")
	f48data := []byte{217, 246, 241, 240, 245, 240, 240, 241, 240, 240}
	f62 := field48.NewField48(specs.Spec87MCField48)
	err := f62.Unpack(f48data)
	if err != nil {
		fmt.Println(err)
	}
}

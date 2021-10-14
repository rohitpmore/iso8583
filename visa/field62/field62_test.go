package field62_test

import (
	"fmt"
	"testing"

	"github.com/moov-io/iso8583/specs"
	"github.com/moov-io/iso8583/visa/field62"
)

func Test_Field62(t *testing.T) {
	fmt.Println("Welcome to the testing of field62")
	f62data := []byte{64, 0, 0, 0, 0, 0, 0, 0, 9, 89, 131, 3, 133, 73, 66, 51}
	fmt.Println("Field62: ", f62data)
	fmt.Printf("Field 62 Binary: %08b\n", f62data)
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err := f62.Unpack(f62data)
	if err != nil {
		fmt.Println(err)
	}
}

func Test_Field62_CVS24(t *testing.T) {
	fmt.Println("Welcome to the testing of field62")
	f62data := []byte{208, 0, 16, 0, 0, 0, 0, 0, 197, 4, 97, 40, 39, 117, 144, 0, 149, 213, 32, 0, 16, 0, 32}
	fmt.Println("Field62: ", f62data)
	fmt.Printf("Field 62 Binary: %08b\n", f62data)
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err := f62.Unpack(f62data)
	if err != nil {
		fmt.Println(err)
	}
}

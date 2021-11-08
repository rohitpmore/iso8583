package field54_test

import (
	"fmt"
	"testing"

	"github.com/moov-io/iso8583/specs"
	"github.com/moov-io/iso8583/visa/field54"
)

func Test_Field54(t *testing.T) {
	fmt.Println("Welcome to the testing of field54")
	f54data := []byte{64, 0, 0, 0, 0, 0, 0, 0, 9, 89, 131, 3, 133, 73, 66, 51}
	fmt.Println("Field54: ", f54data)
	fmt.Printf("Field 54 Binary: %08b\n", f54data)
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err := f54.Unpack(f54data)
	if err != nil {
		fmt.Println(err)
	}
}

func Test_Field54_CVS24(t *testing.T) {
	fmt.Println("Welcome to the testing of field54")
	f54data := []byte{208, 0, 16, 0, 0, 0, 0, 0, 197, 4, 97, 40, 39, 117, 144, 0, 149, 213, 32, 0, 16, 0, 32}
	fmt.Println("Field54: ", f54data)
	fmt.Printf("Field 54 Binary: %08b\n", f54data)
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err := f54.Unpack(f54data)
	if err != nil {
		fmt.Println(err)
	}
}

func Test_Field54_WG_70067(t *testing.T) {
	fmt.Println("Welcome to the testing of field54")
	f54data := []byte{240, 240, 244, 226, 248, 244, 240, 195, 240, 240, 240, 240, 240, 240, 240, 240, 240, 244, 241, 242, 240, 240, 244, 244, 248, 244, 240, 195, 240, 240, 240, 240, 240, 240, 240, 240, 240, 244, 241, 242}
	fmt.Println("Field54: ", f54data)
	fmt.Printf("Field 54 Binary: %08b\n", f54data)
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err := f54.Unpack(f54data)
	if err != nil {
		fmt.Println(err)
	}
}

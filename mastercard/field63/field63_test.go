package field63_test

import (
	"fmt"
	"testing"

	"github.com/moov-io/iso8583/mastercard/field63"
	"github.com/moov-io/iso8583/specs"
)

func Test_Field63(t *testing.T) {
	fmt.Println("Welcome to the testing of field54")
	f63data := []byte{212, 226, 240, 244, 248, 241, 247, 247, 248, 248, 246, 249}
	fmt.Println("Field63: ", f63data)
	fmt.Printf("Field 63 Binary: %08b\n", f63data)
	f63 := field63.NewField63(specs.Spec87MCField63)
	err := f63.Unpack(f63data)
	if err != nil {
		fmt.Println(err)
	}
}

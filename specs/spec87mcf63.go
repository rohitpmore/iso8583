package specs

import (
	"github.com/moov-io/iso8583/encoding"
	"github.com/moov-io/iso8583/field"
	"github.com/moov-io/iso8583/mastercard/field63"
	"github.com/moov-io/iso8583/prefix"
)

var Spec87MCField63 *field63.Field63Spec = &field63.Field63Spec{
	Name: "ISO 8583 v1987 Mastercard Field 54",

	Fields: map[int]field.Field{
		1: field.NewString(&field.Spec{
			Length:      2,
			Description: "Financial Network Code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		2: field.NewString(&field.Spec{
			Length:      1,
			Description: "Interchange Rate Indicator",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		3: field.NewString(&field.Spec{
			Length:      9,
			Description: "Network Reference Number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		4: field.NewString(&field.Spec{
			Length:      9,
			Description: "Banknet Reference Number",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
	},
}

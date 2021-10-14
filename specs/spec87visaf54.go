package specs

import (
	"github.com/moov-io/iso8583/encoding"
	"github.com/moov-io/iso8583/field"
	"github.com/moov-io/iso8583/prefix"
	"github.com/moov-io/iso8583/visa/field54"
)

var Spec87VisaField54 *field54.Field54Spec = &field54.Field54Spec{
	Name: "ISO 8583 v1987 VISA Field 54",

	Fields: map[int]field.Field{
		1: field.NewString(&field.Spec{
			Length:      2,
			Description: "Account Type",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		2: field.NewString(&field.Spec{
			Length:      2,
			Description: "Amount Type",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		3: field.NewString(&field.Spec{
			Length:      3,
			Description: "Currency Code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		4: field.NewString(&field.Spec{
			Length:      1,
			Description: "Amount, Sign",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		5: field.NewString(&field.Spec{
			Length:      12,
			Description: "Amount",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
	},
}

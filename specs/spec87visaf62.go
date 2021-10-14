package specs

import (
	"github.com/moov-io/iso8583/encoding"
	"github.com/moov-io/iso8583/field"
	"github.com/moov-io/iso8583/padding"
	"github.com/moov-io/iso8583/prefix"
	"github.com/moov-io/iso8583/visa/field62"
)

var Spec87VisaField62 *field62.Field62Spec = &field62.Field62Spec{
	Name: "ISO 8583 v1987 VISA Field 62",

	Fields: map[int]field.Field{
		0: field62.NewBitmap(&field.Spec{
			Length:      8,
			Description: "Field62 Bitmap",
			Enc:         encoding.Binary,
			Pref:        prefix.Binary.Fixed,
		}),
		1: field.NewString(&field.Spec{
			Length:      1,
			Description: "Authorization Characteristics Indicator (Bitmap Format)",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		2: field.NewString(&field.Spec{
			Length:      15,
			Description: "Transaction Identifier(Bitmap Format)",
			Enc:         encoding.BCD,
			Pref:        prefix.BCD.Fixed,
			Pad:         padding.Left('0'),
		}),
		3: field.NewString(&field.Spec{
			Length:      4,
			Description: "Validation Code (Bitmap Format)",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		4: field.NewString(&field.Spec{
			Length:      1,
			Description: "Market-Specific Data Identifier",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		5: field.NewString(&field.Spec{
			Length:      2,
			Description: "Duration",
			Enc:         encoding.BCD,
			Pref:        prefix.BCD.Fixed,
		}),
		6: field.NewString(&field.Spec{
			Length:      1,
			Description: "Reserved",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		7: field.NewString(&field.Spec{
			Length:      26,
			Description: "Purchase Identifier",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		8: field.NewString(&field.Spec{
			Length:      6,
			Description: "Auto Rental Check-Out Date, Lodging Check-In Date",
			Enc:         encoding.BCD,
			Pref:        prefix.BCD.Fixed,
		}),
		9: field.NewString(&field.Spec{
			Length:      1,
			Description: "No Show Indicator",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		10: field.NewString(&field.Spec{
			Length:      6,
			Description: "Extra Charges",
			Enc:         encoding.BCD,
			Pref:        prefix.BCD.Fixed,
		}),
		11: field.NewString(&field.Spec{
			Length:      2,
			Description: "Multiple Clearing Sequence Number",
			Enc:         encoding.BCD,
			Pref:        prefix.BCD.Fixed,
		}),
		12: field.NewString(&field.Spec{
			Length:      2,
			Description: "Multiple Clearing Sequence Count",
			Enc:         encoding.BCD,
			Pref:        prefix.BCD.Fixed,
		}),
		13: field.NewString(&field.Spec{
			Length:      1,
			Description: "Restricted Ticket Indicator",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		14: field.NewString(&field.Spec{
			Length:      12,
			Description: "Total Amount Authorized",
			Enc:         encoding.BCD,
			Pref:        prefix.BCD.Fixed,
		}),
		15: field.NewString(&field.Spec{
			Length:      1,
			Description: "Requested Payment Service",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		17: field.NewString(&field.Spec{
			Length:      15,
			Description: "Gateway Transaction Identifier",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		18: field.NewString(&field.Spec{
			Length:      1,
			Description: "Excluded Transaction Identifier Reason Code",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		19: field.NewString(&field.Spec{
			Length:      2,
			Description: "Electronic Commerce Goods Indicator",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		20: field.NewString(&field.Spec{
			Length:      5,
			Description: "Merchant Verification Value",
			Enc:         encoding.Binary,
			Pref:        prefix.Binary.Fixed,
		}),
		21: field.NewString(&field.Spec{
			Length:      4,
			Description: "Online Risk Assessment Risk Score and Reason Codes",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		22: field.NewString(&field.Spec{
			Length:      6,
			Description: "Online Risk Assessment Condition Codes",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		23: field.NewString(&field.Spec{
			Length:      2,
			Description: "Product ID",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		24: field.NewString(&field.Spec{
			Length:      6,
			Description: "Program Identifier",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		25: field.NewString(&field.Spec{
			Length:      1,
			Description: "Spend Qualified Indicator",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
		26: field.NewString(&field.Spec{
			Length:      1,
			Description: "Account Status",
			Enc:         encoding.EBCDIC,
			Pref:        prefix.EBCDIC.Fixed,
		}),
	},
}

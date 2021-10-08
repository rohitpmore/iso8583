package iso8583_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/encoding"
	"github.com/moov-io/iso8583/field"
	"github.com/moov-io/iso8583/padding"
	"github.com/moov-io/iso8583/prefix"
	"github.com/yerden/go-util/bcd"
)

func TestZendaISO(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	spec := &iso8583.MessageSpec{
		Fields: map[int]field.Field{
			0: field.NewString(&field.Spec{
				Length:      4,
				Description: "Message Type Indicator",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			1: field.NewBitmap(&field.Spec{
				Description: "Bitmap",
				Enc:         encoding.Binary,
				Pref:        prefix.Binary.Fixed,
			}),
			2: field.NewString(&field.Spec{
				Length:      11,
				Description: "PAN",
				Enc:         encoding.Binary,
				Pref:        prefix.Binary.L,
			}),
			3: field.NewString(&field.Spec{
				Length:      6,
				Description: "Processing Code",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			4: field.NewString(&field.Spec{
				Length:      12,
				Description: "Amount, Transaction",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			5: field.NewString(&field.Spec{
				Length:      12,
				Description: "Amount, Settlement",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			6: field.NewString(&field.Spec{
				Length:      12,
				Description: "Amount, Cardholder Billing",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			7: field.NewString(&field.Spec{
				Length:      10,
				Description: "Transmission Date and Time",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			8: field.NewString(&field.Spec{
				Length:      8,
				Description: "Amount, Cardholder Billing Fee",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			9: field.NewString(&field.Spec{
				Length:      8,
				Description: "Conversion Rate, Settlement",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			10: field.NewString(&field.Spec{
				Length:      8,
				Description: "Conversion Rate, Cardholder Billing",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			11: field.NewString(&field.Spec{
				Length:      6,
				Description: "System Trace Audit Number",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			12: field.NewString(&field.Spec{
				Length:      6,
				Description: "Time, Local Transaction",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			13: field.NewString(&field.Spec{
				Length:      4,
				Description: "Date, Local Transaction",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			14: field.NewString(&field.Spec{
				Length:      4,
				Description: "Date, Expiration",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			15: field.NewString(&field.Spec{
				Length:      4,
				Description: "Date, Settlement",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			16: field.NewString(&field.Spec{
				Length:      4,
				Description: "Date, Conversion",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			17: field.NewString(&field.Spec{
				Length:      4,
				Description: "Date, Capture",
				Enc:         encoding.Binary,
				Pref:        prefix.Binary.Fixed,
			}),
			18: field.NewString(&field.Spec{
				Length:      4,
				Description: "Merchant Type",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			19: field.NewString(&field.Spec{
				Length:      3,
				Description: "Acquiring Institution Country Code",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
				Pad:         padding.Left('0'),
			}),
			20: field.NewString(&field.Spec{
				Length:      3,
				Description: "PAN Extended, Country Code",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
				Pad:         padding.Left('0'),
			}),
			22: field.NewString(&field.Spec{
				Length:      4,
				Description: "Point-of-Service Entry Mode Code",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			23: field.NewString(&field.Spec{
				Length:      4,
				Description: "Card Sequence Number",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			24: field.NewString(&field.Spec{
				Length:      4,
				Description: "Network International Identifier",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
				Pad:         padding.Left('0'),
			}),
			25: field.NewString(&field.Spec{
				Length:      2,
				Description: "POS Condition Code",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			26: field.NewString(&field.Spec{
				Length:      2,
				Description: "POS PIN Capture Code",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			28: field.NewString(&field.Spec{
				Length:      9,
				Description: "Amount, Transaction Fee",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
			}),
			32: field.NewString(&field.Spec{
				Length:      11,
				Description: "Acquiring Institution Identification Code",
				Enc:         encoding.BCD,
				Pref:        prefix.Binary.L,
				Pad:         padding.Left('0'),
			}),
			33: field.NewString(&field.Spec{
				Length:      11,
				Description: "Forwarding Institution Identification Code",
				Enc:         encoding.BCD,
				Pref:        prefix.Binary.L,
				Pad:         padding.Left('0'),
			}),
			34: field.NewString(&field.Spec{
				Length:      15,
				Description: "PAN, Extended",
				Enc:         encoding.Binary,
				Pref:        prefix.Binary.L,
			}),
			35: field.NewString(&field.Spec{
				Length:      20,
				Description: "Track 2 Data",
				Enc:         encoding.Binary,
				Pref:        prefix.Binary.L,
			}),
			36: field.NewString(&field.Spec{
				Length:      53,
				Description: "Track 3 Data",
				Enc:         encoding.Binary,
				Pref:        prefix.Binary.L,
			}),
			37: field.NewString(&field.Spec{
				Length:      12,
				Description: "Retrieval Reference Number",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
			}),
			38: field.NewString(&field.Spec{
				Length:      6,
				Description: "Authorization Identification Response",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
			}),
			39: field.NewString(&field.Spec{
				Length:      2,
				Description: "Response Code",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
			}),
			41: field.NewString(&field.Spec{
				Length:      8,
				Description: "Card Acceptor Terminal Identification",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
			}),
			42: field.NewString(&field.Spec{
				Length:      15,
				Description: "Card Acceptor Identification Code",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
			}),
			43: field.NewString(&field.Spec{
				Length:      40,
				Description: "Card Acceptor Name/Location",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
			}),
			44: field.NewString(&field.Spec{
				Length:      26,
				Description: "Additional Response Data",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.Binary.L,
			}),
			45: field.NewString(&field.Spec{
				Length:      77,
				Description: "Track 1 Data",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.Binary.L,
			}),
			46: field.NewString(&field.Spec{
				Length:      256,
				Description: "Amounts, Fees",
				Enc:         encoding.Binary,
				Pref:        prefix.Binary.L,
			}),
			47: field.NewString(&field.Spec{
				Length:      256,
				Description: "Additional Data—National",
				Enc:         encoding.Binary,
				Pref:        prefix.Binary.L,
			}),
			48: field.NewString(&field.Spec{
				Length:      256,
				Description: "Additional Data—Private",
				Enc:         encoding.Binary,
				Pref:        prefix.Binary.L,
			}),
			49: field.NewString(&field.Spec{
				Length:      3,
				Description: "Currency Code, Transaction",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
				Pad:         padding.Left('0'),
			}),
			50: field.NewString(&field.Spec{
				Length:      3,
				Description: "Currency Code, Settlement",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
				Pad:         padding.Left('0'),
			}),
			51: field.NewString(&field.Spec{
				Length:      3,
				Description: "Currency Code, Cardholder Billing",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
				Pad:         padding.Left('0'),
			}),
			52: field.NewString(&field.Spec{
				Length:      8,
				Description: "Personal Identification Number (PIN) Data",
				Enc:         encoding.Binary,
				Pref:        prefix.Binary.Fixed,
			}),
			53: field.NewString(&field.Spec{
				Length:      16,
				Description: "Security-Related Control Information",
				Enc:         encoding.BCD,
				Pref:        prefix.BCD.Fixed,
			}),
			54: field.NewString(&field.Spec{
				Length:      121,
				Description: "Additional Amounts",
				Enc:         encoding.Binary,
				Pref:        prefix.Binary.L,
			}),
			55: field.NewString(&field.Spec{
				Length:      256,
				Description: "Integrated Circuit Card (ICC) Related Data",
				Enc:         encoding.Binary,
				Pref:        prefix.Binary.L,
			}),
			56: field.NewString(&field.Spec{
				Length:      256,
				Description: "Payment Account Reference Data",
				Enc:         encoding.Binary,
				Pref:        prefix.Binary.L,
			}),
			57: field.NewString(&field.Spec{
				Length:      256,
				Description: "Reserved Nationa",
				Enc:         encoding.Binary,
				Pref:        prefix.Binary.L,
			}),
			58: field.NewString(&field.Spec{
				Length:      256,
				Description: "Reserved National",
				Enc:         encoding.Binary,
				Pref:        prefix.Binary.L,
			}),
			59: field.NewString(&field.Spec{
				Length:      15,
				Description: "National Point-of-Service Geographic Data",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.Binary.L,
			}),
			60: field.NewString(&field.Spec{
				Length:      6,
				Description: "Additional POS Information",
				Enc:         encoding.Binary,
				Pref:        prefix.Binary.L,
			}),
			61: field.NewString(&field.Spec{
				Length:      19,
				Description: "Other Amounts",
				Enc:         encoding.Binary,
				Pref:        prefix.Binary.L,
			}),
			62: field.NewString(&field.Spec{
				Length:      256,
				Description: "Custom Payment Service Fields Bitmap",
				Enc:         encoding.Binary,
				Pref:        prefix.Binary.L,
			}),
			// 62: field.NewComposite(&field.Spec{
			// 	Length:      256,
			// 	Description: "Custom Payment Service Fields Bitmap",
			// 	Pref:        prefix.Binary.L,
			// 	Tag: &field.TagSpec{
			// 		Sort: sort.StringsByInt,
			// 	},
			// 	Subfields: map[string]field.Field{
			// 		"0": field.NewBitmap(&field.Spec{
			// 			Length:      8,
			// 			Description: "Bitmap",
			// 			Enc:         encoding.Binary,
			// 			Pref:        prefix.Binary.Fixed,
			// 		}),
			// 		"1": field.NewString(&field.Spec{
			// 			Length:      1,
			// 			Description: "Authorization Characteristics Indicator",
			// 			Enc:         encoding.EBCDIC,
			// 			Pref:        prefix.EBCDIC.Fixed,
			// 		}),
			// 		"2": field.NewString(&field.Spec{
			// 			Length:      15,
			// 			Description: "Transaction Identifier",
			// 			Enc:         encoding.BCD,
			// 			Pref:        prefix.BCD.Fixed,
			// 			Pad:         padding.Left('0'),
			// 		}),
			// 	},
			// }),
			63: field.NewString(&field.Spec{
				Length:      256,
				Description: "SMS Private-Use Fields",
				Enc:         encoding.Binary,
				Pref:        prefix.Binary.L,
			}),
		},
	}
	isoFromG := "FgECALwSNFYSNFYAAAAAAAAAAAAAAAIAMmBkgQjwgDYAAAAAAAAAEQAJIgcDQGEAAABJERZCIwhAAAAABldgI/Hy9vXw9/T58fHx9vn19/L18PP05vOZ46TCh6n2x5TYqdL15omVo4WZhoWTk0DCgYOSQNSBmZKFo0DT40DTwdLFQMPJ40BA5OPk4gVAQEBA8QhACvT58PDw+PTx8vEGAAAAAAAAEEAAAAAAAAAACVmDA4VJQjMFgAAAAAI="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(spec)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Field 62:", f62bytes)
}

func TestDecodeBCD(t *testing.T) {
	fmt.Println("let's test BCD decoding.")
	dec := bcd.NewDecoder(bcd.Standard)
	dst := make([]byte, 8)
	_, err := dec.Decode(dst, []byte{6, 87, 96, 35, 241, 242, 246, 245, 240, 247, 244, 249, 241, 241, 241, 246, 249, 245, 247, 242,
		245, 240, 243, 244, 230, 243, 153, 227, 164, 194, 135, 169, 246, 199, 148, 216, 169, 210, 245, 230, 137, 149, 163, 133,
		153, 134, 133, 147, 147, 64, 194, 129, 131, 146, 64, 212, 129, 153, 146, 133, 163, 64, 211, 227, 64, 211, 193, 210,
		197, 64, 195, 201, 227, 64, 64, 228, 227, 228, 226, 5, 64, 64, 64, 64, 241, 8, 64, 10, 244, 249, 240, 240, 240, 248,
		244, 241, 242, 241, 6, 0, 0, 0, 0, 0, 0, 16, 64, 0, 0, 0, 0, 0, 0, 0, 9, 89, 131, 3, 133, 73, 66, 51, 5, 128, 0, 0, 0, 2})
	if err != nil {
		fmt.Println("comes here?")
		t.Error("TestDecodeBCD: ", err.Error())
	}
	fmt.Println(string(dst))

}

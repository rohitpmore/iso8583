package iso8583_test

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"testing"

	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/encoding"
	"github.com/moov-io/iso8583/mastercard/field48"
	mcf54 "github.com/moov-io/iso8583/mastercard/field54"
	"github.com/moov-io/iso8583/specs"
	"github.com/moov-io/iso8583/visa/field54"
	"github.com/moov-io/iso8583/visa/field62"
	"github.com/yerden/go-util/bcd"
)

func TestZendaISO(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECALwSNFYSNFYAAAAAAAAAAAAAAAIAMmBkgQjwgDYAAAAAAAAAEQAJIgcDQGEAAABJERZCIwhAAAAABldgI/Hy9vXw9/T58fHx9vn19/L18PP05vOZ46TCh6n2x5TYqdL15omVo4WZhoWTk0DCgYOSQNSBmZKFo0DT40DTwdLFQMPJ40BA5OPk4gVAQEBA8QhACvT58PDw+PTx8vEGAAAAAAAAEEAAAAAAAAAACVmDA4VJQjMFgAAAAAI="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}
}

func TestZendaMCISO(t *testing.T) {
	fmt.Println("Welcome to Zenda MC ISO test.")
	isoFromG := "8PLw8D77RgGI4eIK8PDw8PDw8PDw8PDw8PDw9PX28PDw8PDw8PDw9PX28PDw8PDw8PDw9PX28fHw9/H38fL08Pbx8PDw8PDw9vHw8PDw8PD39fn5+PXw+fHy8/jx8fD38fHw+PHx8Pb1+fHy8PXx8PDw8Pn49PDw9PTy9/Hx8Pnw8PDw8PD18Pbx8fD38Pnx8vL49PPz8Pn18/bx+PD39/Dw+fXz9kBAQEBAQMPl4mHXyMHZ1EDw+fXz9mBg9PnwQNRA4oGVQNmBlJaVQEBAQEBAw8Hw8fDZ9vHw9fHw8PDw+PTw+PTw+PTw8fH2nwIGAAAAAARWnyYIRTBvIJuez6aCAhgAnzYCAEGfJwGAnxAHBgESA6CgAJUFgIAEgABfKgIIQJoDIREHnAEAnzcEi8CGp58JApYAnzQDQgAAhAegAAAAmAhAnx4IODQyODYwNTefMwPg+MifNQEinxoCCEDw8vbw8PDw8PDw8PDw8/Dw+PTw+fT1+PNAQEBAQPDx8tTi8PL38/j2+fn0+A=="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	// fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Maestro)
	err = message.Unpack(rawISO)
	if err != nil {
		fmt.Println(err)
	}
	// f62bytes, _ := message.GetField(62).Bytes()
	// fmt.Println("Let's unpack Field 62 now!")
	// f62 := field62.NewField62(specs.Spec87VisaField62)
	// err = f62.Unpack(f62bytes)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func TestZendaISO_CVS24(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAStJNSEAAAADEABCjBAYAwJaAAEANmZmgQjwojYAAAAAAAAABWQAAAAABWQQCSEzEGEAAAACAzUmEhAQWRIIQAUQAAAABkE3RvHy+PLx+fDy8PPz9vD18/bw8PH49PT09fD39/Dw+fXz9kBAw+XiYdfIwdnUwcPoQHvw+fXz9kBAQEDiwdVA2cHU1tVAQEBAw8Hk4gVAQEBA8ghACEBZAQBWnzMD4PjIlQWAgAiAAJ83BEYdfpOfEAcGARIDoDAAnyYI3PRuZte0iLOfNgIAAoICGACcAQCfGgIIQJoDIRAJnwIGAAAAAAVkXyoCCECEB6AAAACYCEAK8Pbw8PD59PX48wZFAAAQAAEX0AAQAAAAAADFBGEoJ3WQAJXVIAAQACAFgAAAAAI="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

}

func TestZendaISO_CVS24_Advice(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAUFJNQUAAAACEABAyRD/AgAAAAIgPvtkgQ5h4jYAAAAAAAAABWQAAAAABWQAAAAABWQQEAAAAGEAAABhAAAAAUF0AAAAEAkQEBAQWRIIQAUAAAZBN0bx8vjz8vnw8PHx8/fw9vH09/Pw8PT09PXw9/fw8Pn18/ZAQMPl4mHXyMHZ1MHD6EB78Pn18/ZAQEBA4sHVQNnB1NbVQEBAQMPB5OIJ5fFA8PDw8PDwCEAIQAhARwEARJ9uBEBAQECfMwPg+MiVBYCACIAAnzcERh1+k58mCMTDxvT2xfb2nzYCAAKCAhgAnAEAnxoCCECaAyEQCZ9bBUBAQEBACvD28PDw+fT1+PMEBQAAEB3ABhAAAAAAAOcEYSgndZAAlQAAAAAFZEAgABAAIBaVICAAApEB6UBAQEBAQBAGIyfR8/P2"

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

}

func TestZendaISO_CVS25(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAStJNQkAAAADEABCjBAYAwGdAAEANmZmgQjwojYAAAAAAAAABFYAAAAABFYQCSE0JWEAAAACE1gmEhAQWRIIQAUQAAAABkE3RvHy+PLx9/Dy8fP1+fD18/bw8PH49PT09fD39/Dw+fXz9kBAw+XiYdfIwdnUwcPoQHvw+fXz9kBAQEDiwdVA2cHU1tVAQEBAw8Hk4gVAQEBA8ghACEBZAQBWnzMD4PjIlQWAgAiAAJ83BPq6OtWfEAcGARIDoLAAnyYI4ABcZeen0uyfNgIAA4ICGACcAQCfGgIIQJoDIRAJnwIGAAAAAARWXyoCCECEB6AAAACYCEAK8Pbw8PD59PX48wZFAAAQAAEX0AAQAAAAAADFBGEoJ3ZlYQHVIAAQACAFgAAAAAI="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}
}

func TestZendaISO_CVS25_Advice(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAUFJNSMAAAACEABAyRD/AgAAAAIgPvtkgQ5h4jYAAAAAAAAABFYAAAAABFYAAAAABFYQEAAAAGEAAABhAAAAAUGCAAAAEAkQEBAQWRIIQAUAAAZBN0bx8vjz8vLw8PHx8/fw+PH19/Xw8PT09PXw9/fw8Pn18/ZAQMPl4mHXyMHZ1MHD6EB78Pn18/ZAQEBA4sHVQNnB1NbVQEBAQMPB5OIJ5fFA8PDw8PDwCEAIQAhARwEARJ9uBEBAQECfMwPg+MiVBYCACIAAnzcE+ro61Z8mCMXw8PD1w/b1nzYCAAOCAhgAnAEAnxoCCECaAyEQCZ9bBUBAQEBACvD28PDw+fT1+PMEBQAAEB3ABhAAAAAAAOcEYSgndmVhAQAAAAAEVkAgABAAIBaVICAAApEB6UBAQEBAQBAGIyfR8/P2"

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}
}

func TestZendaISO_WG26(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAStJNREAAAABEABCjBAYARAoAAEANmZmgQjwojYAAAAAAAAABBIAAAAABBIQCSFXMWEAAAAAmAYmEhAQWRIIQAUQAAAABkRFAPHy+PLx9vDw+fjw9/Dw+fDw8PLx9PT09fHw8PH28Pnw80BA5sHTx9nFxdXiQHvx9vD58EBAQEBAQEDiwdVA2cHU1tVAQEBAw8Hk4gVAQEBA8ghACEBZAQBWnzMD4EjIlQWAAAiAAJ83BLrRnnqfEAcGARIDoLAAnyYIDFR0+ueVu6efNgIABoICGACcAQCfGgIIQJoDIRAJnwIGAAAAAAQSXyoCCECEB6AAAACYCEAK8Pbw8PD59PX48gZFAAAQAAEX0AAQAAAAAADFAwEoJ5BRZmLVIAAYAAUFgAAAAAI="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

}

func TestZendaISO_WG26_Advice(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAUFJNQUAAAACEABAyRD/AgAAAAIgPvtkgQ5h4jYAAAAAAAAABBIAAAAABBIAAAAABBIQEAAAAGEAAABhAAAABWk2AAAAEAkQEBAQWRIIQAUAAAZERQDx8vjz8vTw8PD29fbw+PT59vLw8PT09PXx8PDx9vD58PNAQObB08fZxcXV4kB78fbw+fBAQEBAQEBA4sHVQNnB1NbVQEBAQMPB5OIJ5fFA8PDw8PDwCEAIQAhARwEARJ9uBEBAQECfMwPgSMiVBYAACIAAnzcEutGeep8mCPDD9fT39MbBnzYCAAaCAhgAnAEAnxoCCECaAyEQCZ9bBUBAQEBACvD28PDw+fT1+PIEBQAAEB3ABhAAAAAAAOcDASgnkFFmYgAAAAAEEkAgABgABRaVICAAApEB6UBAQEBAQBAAA5jw5dfX"

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

}

func TestZendaISO_WG27(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAStJNSAAAAABEABCjBAYARAnAAEANmZmgQjwojYAAAAAAAAABgEAAAAABgEQCSFYEmEAAAABaVQmEhAQWRIIQAUQAAAABkRFAPHy+PLx9vDx9vn19fDw+fDw8PLx9PT09fHw8PH28Pnw80BA5sHTx9nFxdXiQHvx9vD58EBAQEBAQEDiwdVA2cHU1tVAQEBAw8Hk4gVAQEBA8ghACEBZAQBWnzMD4EjIlQWAAAiAAJ83BDu0xDSfEAcGARIDoCAAnyYIQWqAc/Bi3qyfNgIAB4ICGACcAQCfGgIIQJoDIRAJnwIGAAAAAAYBXyoCCECEB6AAAACYCEAK8Pbw8PD59PX48gZFAAAQAAEX0AAQAAAAAADFAwEoJ5CSJzLVIAAYAAUFgAAAAAI="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}
}

func TestZendaISO_WG27_Advice(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAUFJNQQAAAACEABAyRD/AgAAAAIgPvtkgQ5h4jYAAAAAAAAABgEAAAAABgEAAAAABgEQEAAAAGEAAABhAAAABWlEAAAAEAkQEBAQWRIIQAUAAAZERQDx8vjz8vjw8PD29fbw+PDy8Pnw8PT09PXx8PDx9vD58PNAQObB08fZxcXV4kB78fbw+fBAQEBAQEBA4sHVQNnB1NbVQEBAQMPB5OIJ5fFA8PDw8PDwCEAIQAhARwEARJ9uBEBAQECfMwPgSMiVBYAACIAAnzcEO7TENJ8mCPTx9sH48PfznzYCAAeCAhgAnAEAnxoCCECaAyEQCZ9bBUBAQEBACvD28PDw+fT1+PIEBQAAEB3ABhAAAAAAAOcDASgnkJInMgAAAAAGAUAgABgABRaVICAAApEB6UBAQEBAQBAAA5jw5dfX"

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}
}

func TestZendaISO_WG28(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAStJNQgAAAADEABCjBAYAwLCAAEANmZmgQjwojYAAAAAAAAAAnEAAAAAAnEQCSFYRGEAAAACIREmEhAQWRIIQAUQAAAABkRFAPHy+PLx9vDy8vHx8vDw+fDw8PLx9PT09fHw8PH28Pnw80BA5sHTx9nFxdXiQHvx9vD58EBAQEBAQEDiwdVA2cHU1tVAQEBAw8Hk4gVAQEBA8ghACEBZAQBWnzMD4EjIlQWAAAiAAJ83BO0XMEWfEAcGARIDoCAAnyYIMOppOzUaolefNgIACIICGACcAQCfGgIIQJoDIRAJnwIGAAAAAAJxXyoCCECEB6AAAACYCEAK8Pbw8PD59PX48gZFAAAQAAEX0AAQAAAAAADFBGEoJ5EkiJjVIAAYAAUFgAAAAAI="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}
}

func TestZendaISO_WG28_Advice(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAUFJNQMAAAACEABAyRD/AgAAAAIgPvtkgQ5h4jYAAAAAAAAAAnEAAAAAAnEAAAAAAnEQEAAAAGEAAABhAAAABWlRAAAAEAkQEBAQWRIIQAUAAAZERQDx8vjz8vPw8PD29fbw8/j19fXw8PT09PXx8PDx9vD58PNAQObB08fZxcXV4kB78fbw+fBAQEBAQEBA4sHVQNnB1NbVQEBAQMPB5OIJ5fFA8PDw8PDwCEAIQAhARwEARJ9uBEBAQECfMwPgSMiVBYAACIAAnzcE7RcwRZ8mCPPwxcH2+fPCnzYCAAiCAhgAnAEAnxoCCECaAyEQCZ9bBUBAQEBACvD28PDw+fT1+PIEBQAAEB3ABhAAAAAAAOcEYSgnkSSImAAAAAACcUAgABgABRaVICAAApEB6UBAQEBAQBAAA5jw5dfX"

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}
}

func TestZendaISO_WG_70005_Auth(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAStJNQMAAAAGEABCjBAiBlwiAAEANmZmgQjwojYAAAAAAAAAB5AAAAAAB5AQJgJCIWEAAAACQ4kmEhAmWRIIQAUQAAAABkRFAPHy+fnx9vDy9PP58PDw+fDw8PXx9PT09fHw8PH28Pnw80BA5sHTx9nFxdXiQHvx9vD58EBAQEBAQEDiwdVA2cHU1tVAQEBAw8Hk4gVAQEBA8ghACEBZAQBWnzMD4EjIlQWAAAiAAJ83BMUC7LCfEAcGARIDoCAAnyYIt19DtgAjbN6fNgIAHIICGACcAQCfGgIIQJoDIRAmnwIGAAAAAAeQXyoCCECEB6AAAACYCEAK8Pbw8PD59PX48gZFAAAQAAEX0AAQAAAAAADFBYEpkJdBJInVIAAYAAUFgAAAAAI="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}
}

func TestZendaISO_WG_70004_Auth(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAStJNRQAAAADEABCjBAiA1JpAAEANmZmgQjwojYAAAAAAAAACGkAAAAACGkQJgJBRWEAAAACI1AmEhAmWRIIQAUQAAAABkRFAPHy+fnx9vDy8vP18fDw+fDw8PXx9PT09fHw8PH28Pnw80BA5sHTx9nFxdXiQHvx9vD58EBAQEBAQEDiwdVA2cHU1tVAQEBAw8Hk4gVAQEBA8ghACEBZAQBWnzMD4EjIlQWAAAiAAJ83BAxpltyfEAcGARIDoCAAnyYIRMEwAiuBpVifNgIAG4ICGACcAQCfGgIIQJoDIRAmnwIGAAAAAAhpXyoCCECEB6AAAACYCEAK8Pbw8PD59PX48gZFAAAQAAEX0AAQAAAAAADFBGEpkJcFWUXVIAAYAAUFgAAAAAI="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}
}

func TestZendaISO_WG_70004_Advice(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAUFJNRcAAAACEABAyRD/AgAAAAIgPvtkgQ5h4jYAAAAAAAAACGkAAAAACGkAAAAACGkQJgAAAGEAAABhAAAABTQhAAAAECUQJhAmWRIIQAUAAAZERQDx8vn58vHw8PD38Pfw9Pnz9PHw8PT09PXx8PDx9vD58PNAQObB08fZxcXV4kB78fbw+fBAQEBAQEBA4sHVQNnB1NbVQEBAQMPB5OIJ5fFA8PDw8PDwCEAIQAhARwEARJ9uBEBAQECfMwPgSMiVBYAACIAAnzcEDGmW3J8mCPT0w/Hz8PDynzYCABuCAhgAnAEAnxoCCECaAyEQJp9bBUBAQEBACvD28PDw+fT1+PIEBQAAEB3ABhAAAAAAAOcEYSmQlwVZRQAAAAAIaUAgABgABRaVICAAApEB6UBAQEBAQBAAA5jw5dfX"

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}
}

func TestZendaISO_WG_70065_Auth(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAUBJNQUAAAABEABCjBARAb48AAEANmZmgQjwpjYAAAAAAAAAATkAAAAAATkRBRgFI2EAAAABKIkmEhEGWRIIQAUQAAAABkRFAPHz8Pnx9vDx8vj58PDw+fDw8PLy9PT09fHw8PH28Pnw80BA5sHTx9nFxdXiQHvx9vD58EBAQEBAQEDiwdVA2cHU1tVAQEBAw8Hk4gVAQEBA8ghACEAU8PD04vj08MPw8PDw8PDw8PDw8PBZAQBWnzMD4EjIlQWAAAiAAJ83BAmrXEKfEAcGARIDoCAAnyYILqvazxH5mzWfNgIAO4ICGACcAQCfGgIIQJoDIREFnwIGAAAAAAE5XyoCCECEB6AAAACYCEAK8Pbw8PD59PX48gZFAAAQAAEX0AAQAAAAAADFAwEwllEjJFfUIAAYAAUFgAAAAAI="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}
}

func TestZendaISO_WG_70066_Auth(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAUBJNSAAAAABEABCjBARARk3AAEANmZmgQjwpjYAAAAAAAAABVEAAAAABVERBRgGEWEAAAACBYEmEhEGWRIIQAUQAAAABkRFAPHz8Pnx9vDy8PX48vDw+fDw8PLy9PT09fHw8PH28Pnw80BA5sHTx9nFxdXiQHvx9vD58EBAQEBAQEDiwdVA2cHU1tVAQEBAw8Hk4gVAQEBA8ghACEAU8PD04vj08MPw8PDw8PDw8PDw8PBZAQBWnzMD4EjIlQWAAAiAAJ83BHk4tnqfEAcGARIDoCAAnyYIMuOegNEFn7yfNgIAPIICGACcAQCfGgIIQJoDIREFnwIGAAAAAAVRXyoCCECEB6AAAACYCEAK8Pbw8PD59PX48gZFAAAQAAEX0AAQAAAAAADFAwEwllFxFTDUIAAYAAUFgAAAAAI="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}
}

func TestZendaISO_WG_70064_Auth(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAUBJNRUAAAAGEABCjBARBr49AAEANmZmgQjwpjYAAAAAAAAABlEAAAAABlERBRgEMGEAAAAChyQmEhEGWRIIQAUQAAAABkRFAPHz8Pnx9vDy+Pfy9fDw+fDw8PLy9PT09fHw8PH28Pnw80BA5sHTx9nFxdXiQHvx9vD58EBAQEBAQEDiwdVA2cHU1tVAQEBAw8Hk4gVAQEBA8ghACEAU8PD04vj08MPw8PDw8PDw8PDw8PBZAQBWnzMD4EjIlQWAAAiAAJ83BIF5AlWfEAcGARIDoKAAnyYIqZy2BzP79EKfNgIAOoICGACcAQCfGgIIQJoDIREFnwIGAAAAAAZRXyoCCECEB6AAAACYCEAK8Pbw8PD59PX48gZFAAAQAAEX0AAQAAAAAADFBYEwllBwCSLUIAAYAAUFgAAAAAI="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}
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

func TestZendaISO_WG_70067_Auth(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAUBJNQIAAAACEABCjBAUAhk5AAEANmZmgQjwpjYAAAAAAAAABBIAAAAABBIRBRlCOGEAAAABMQQmEhEGWRIIQAUQAAAABkRFAPHz8Pnx9vDx8/Hw9fDw+fDw8PLx9PT09fHw8PH28Pnw80BA5sHTx9nFxdXiQHvx9vD58EBAQEBAQEDiwdVA2cHU1tVAQEBAw8Hk4gVAQEBA8ghACEAU8PD04vj08MPw8PDw8PDw8PD08fJZAQBWnzMD4EjIlQWAAAiAAJ83BDXSitafEAcGARIDoCAAnyYIQIrod+84SXifNgIAPYICGACcAQCfGgIIQJoDIREFnwIGAAAAAAQSXyoCCECEB6AAAACYCEAK8Pbw8PD59PX48gZFAAAQAAEX0AAQAAAAAADFA4EwlwlYUxbUIAAYAAUFgAAAAAI="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)

}

func TestZendaISO_WG_70094_Auth(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAThJNQgAAAABEABCjBAeAQZPAAEANn5mgQjwojYAAAAAAAAABogAAAAABogRCQFUEWEAAACRUocXVAARCCYREQlUEQhABRAAAAAGQjFo8fPx8vDx+fH18vj3+fn5+fn5+fnw8PD09fnz+flAQEBAQEDiwcbF5sHoQHvw+fH5QEBAQEBAQEBAQOLB2cHj1sfBQEBAQEDDweTiBUBAQEDyCEAIQGIBAF+fMwPg+MiVBYAACIAAnzcE46/CGZ8QBwYBEgOgoACfJgiKn531WgzQfZ82AgALggIYAJwBAJ8aAghAmgMhEQifAgYAAAAABohfKgIIQJ8DBgAAAAAAAIQHoAAAAJgIQArw9vDw8Pn18PfwBnUAABAAARbAABAAAAAAAMUDATEwaFGQJyAAFgAABYAAAAAC"

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_Waqar_70061(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAWRJNQgAAAACEABCyBAeAgbPAAIAPv9mgQjw6jYAQAAAAAAABlQAAAAABlQAAAAABlQRBQBBNmEAAABhAAAAkjVgF0EAEQQmEREFEQVUEQhABRAAAAAGMVeH8fPw+PLw+fLz9fbw9PX59vfz+fPw8PDw8PDw8PD09fn29/PiwcbF5sHoQHvz8vXxQEBAQEBAQEBAQMPk18XZ48nV1kBAQEDDweTiCUBAQEDyQEBA8ghACEAIQCABAQEAAAAAcQEAbl8wAgIBnwYHoAAAAJgIQJ8zA+D4yJUFgAAEgACfNwQ1vfkunxAHBgESA6CgAJ8mCP4uOaoAXLcUnzYCAAiCAhgAnAEAnxoCCECaAyERBJ8CBgAAAAAGVF8qAghAnwMGAAAAAAAAhAegAAAAmAhACvD28PDw+fXw8fQGBQAAEAABFUAAEAAAAAAAA4EwkCSWGXQgABYAAAmAICAAA+jlydk="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_Waqar_70094(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAThJNQgAAAABEABCjBAeAQZPAAEANn5mgQjwojYAAAAAAAAABogAAAAABogRCQFUEWEAAACRUocXVAARCCYREQlUEQhABRAAAAAGQjFo8fPx8vDx+fH18vj3+fn5+fn5+fnw8PD09fnz+flAQEBAQEDiwcbF5sHoQHvw+fH5QEBAQEBAQEBAQOLB2cHj1sfBQEBAQEDDweTiBUBAQEDyCEAIQGIBAF+fMwPg+MiVBYAACIAAnzcE46/CGZ8QBwYBEgOgoACfJgiKn531WgzQfZ82AgALggIYAJwBAJ8aAghAmgMhEQifAgYAAAAABohfKgIIQJ8DBgAAAAAAAIQHoAAAAJgIQArw9vDw8Pn18PfwBnUAABAAARbAABAAAAAAAMUDATEwaFGQJyAAFgAABYAAAAAC"

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_Waqar_70104(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAThJNQkAAAAGEABCjBAgBgWGAAEANn5mgQjwojYAAAAAAAAACVgAAAAACVgREwI0NGEAAACQAkUYNAAREiYRERNUEQhABRAAAAAGQjFo8fPx9vDy+fDw8vT1+fn5+fn5+fnw8PD09fnz+flAQEBAQEDiwcbF5sHoQHvw+fH5QEBAQEBAQEBAQOLB2cHj1sfBQEBAQEDDweTiBUBAQEDyCEAIQGIBAF+fMwPg+MiVBYAACIAAnzcEFxSDUZ8QBwYBEgOgoACfJgidF6Uxf9fPr582AgATggIYAJwBAJ8aAghAmgMhERKfAgYAAAAACVhfKgIIQJ8DBgAAAAAAAIQHoAAAAJgIQArw9vDw8Pn18PfwBnUAABAAARbAABAAAAAAAMUFgTFwknRWIiAAFgAABYAAAAAC"

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_Waqar_70113(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAT9JNQkAAAACEABCjBAeAgYXAAEANn5mgQjwojYAAAAAAAAABlQAAAAABlQRFAFJJGEAAABnKGcXSQAREyYRERRUEQhABxAAAAAGQjFo8fPx9/Dx9vfy+Pb3+fn5+fn5+fnw8PD09fnz+flAQEBAQEDiwcbF5sHoQHvw+fH5QEBAQEBAQEBAQOLB2cHj1sfBQEBAQEDDweTiBUBAQEDyCEAIQGkBAGafbgQgcAAAnzMD4PjIlQUAAAAAAJ83BJXS3TWfEAcGARIDoAAAnyYICBEKa3962myfNgIAFYICAACcAQCfGgIIQJoDIRETnwIGAAAAACUAXyoCCECfAwYAAAAAAACEB6AAAACYCEAK8Pbw8PD59fD38AZ1AAAQAAEWwAAQAAAAAADFA4ExgGVkdYQgABYAAAWAAAAAAg=="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_70124(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAQlJNREAAAADEABCjBAXAwISAAEAtn5kgYjwoDYAAAAAAAAAJAAAAAAAAAAClgAAAAAClhEVIglWYQAAAAGWdBcJVhEVJhERFlMQCEABAFkGRDEGBkQxBvHz8fnz8PT39vj38fn5+fn5+fn58/T18fbw8vTw+PjyQEBA48HZx8XjS8PW1EBAQEBAQEBAQEBAQEDC2dbW0tPo1UDXwdnS1NXk4gtAQEBAQEBAQEBA1AhACEAK8vfw8PD19fT09QVQAAAABxbAABAAAAAAAOYEYTGXl5YCdiAABQAABYAAAAACFWYAEsAJ+fXw9/D08vT3zwXy8fLx9Q4AQAAAAAAAAPHxQPX59Q=="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_70125(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAPNJNQUAAAACEABCjBAXAgJXAAQAtn5kgQzwoDYAAABAAAAAAAAAAAAAAAAClgAAAAAClhEVIglWYQAAAAGWdBcRBhEVJhERFlMQCEABAFkGRDEG8fPx+fPw9Pf2+Pfx8PHx8vD3+fn5+fn5+fnz9PXx9vDy9PD4+PJAQEDjwdnHxeNLw9bUQEBAQEBAQEBAQEBAQMLZ1tbS0+jVQNfB2dLU1eTiAfcIQAhACvL38PDw9fX09PUFUAAAAAcWwAAQAAAAAADmBGExl5eWAnYgAAUAAAegAAAAAiUCAQABlnQRFSIJVgAABEMQYAAAAAAA"

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_Waqar_70087(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAURJNRQAAAABEABCjBAcAVNGAAEANmZmgYjwpjYAAAAAAAAAIwUAAAAAIwURCABQQmEAAAABJkkmEREIUxAIQAUQAAAABkOZAAZERQDx8/Hy8vHw8fL29fDz8fTy9/D48/T09PXw8PP08fTy9/ZAQOPB2cfF40DjYPH08vdAQEBAQEBAQEBA4sHVQNHW4sVAQEBAQMPB5OIFQEBAQPIIQAhAFPDw9OL49PDD8PDw8PDw8PDw9fP2WQEAVp8zA+D4yJUFgAAIgACfNwQNUYI5nxAHBgESA6CgAJ8mCPk9V1xrYOVmnzYCAAqCAhgAnAEAnxoCCECaAyERB58CBgAAAABRAF8qAghAhAegAAAAmAhACvD28PDw+fXx8vkGRQAAEAABF9AAEAAAAAAAxQMBMSAwQjN11CAABQAABYAAAAAC"

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_Waqar_70086(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAQ1JNRcAAAABEABCjBAaAaZ/AAEAtmZkgYjxoDYAAAAAAAAAJAAAAAAAAAAohwAAAAAohxEHIzQTYQAAAGMoNiYREQhTAAhAAQBZBkaSFgZGkhbx8/Hx8/f28/L48/b5+fn5+fn5+fHz8ffz8PDw8PHy9/f28ebm5kDD1uLjw9ZAw9bUQEBAQEBAQEBA+PDwYPn19WDy8vnyQObB5OILQEBAQEBAQEBAQNQIXNNs0/Ly8/YIQAhACvXz8PDw+fjw8vcFUAAAAAcWwAAQAAAAAADmAwExGEhTQQggAAQAAAWAAAAAAhVmABLACfn18Pfw9PL0988F8vHy8fUOAEAAAAAAAADx8UD1+fU="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_Waqar_70127(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAQJJNSMAAAAGEABCjBANBgNQAAEAtn5kgYjwoDYAAAAAAAAAJAAAAAAAAAAxIgAAAAAxIhEXFwgDYQAAAAMFRBIIAxEXJhERGFQRCEABAFkGQjFoBkIxaPHz8vHz8fD58fD29vn5+fn5+fn58/P09vP08fH3+Pj0QEBA4sHGxebB6EvD1tRAe/D58flAQEBAQED49/f18PX08PTwQEBAw8Hk4gtAQEBAQEBAQEBA1AhACEAO8Pbw8PD59fD38PDw8PAFUAAAAAcWwAAQAAAAAADmBYEyFhaDJVQgABYAAAWAAAAAAgpmAAfABfn18PfwDgBAAAAAAAAA8fFA9fn1"

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_Sanjeev_70122(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAN9JNQIAAAACEABCjBArAgRzAAEAtn5kgQjgoDYAAAAAAAAAIAAAAAAAAAADWAAAAAADWBEVCAJVYQAAAFUAUQMCVREVJgkRFVkSCEABAFkGSUMA8fPx+fD49fXw8PXx8Pfx+fLy8fLw8PD38/fx+PTy8/D4+PPZyePFQMHJxEBgQMXD1tRAQEBAQEBAQMXj48XZ4kBAQEBAQEDXweTiCEAIQAr08vDw8PH38/H5BUEAAAAHEcAAAAAAAAAA5gOBMZKJdReRBYAAAAACD2YADMAF+fT19vbPA/T19g=="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_Sanjeev_70110(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAS1JNQgAAAACEABCjBARAgaVAAEANn5mgQjwojYAAAAAAAAAATAAAAAAATARExkTKGEAAAAmVoYREwAREyYJERRZEghABRAAAAAGSUMA8fPx9/H58vb19vj2+fn5+fn5+fnw8PD19Pnw9/lAQEBAQEDZyePFQMHJxEDw9fn09EBAQEBAQEBAQNfTxcHiwdXj1tVAQEDDweTiBUBAQEDyCEAIQFkBAFafMwPg+MiVBYCACIAAnzcE5ZH1wJ8QBwYBEgOgsACfJgiYbQkoJ2kigp82AgAWggIYAJwBAJ8aAghAmgMhEROfAgYAAAAAATBfKgIIQIQHoAAAAJgIQArw9vDw8Pn09fj4BHUAABAWwAAQAAAAAADFA4ExdpIIFwQgAAkAAAWAAAAAAg=="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_Sanjeev_70111(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAS1JNRcAAAACEABCjBARAgZPAAEANn5mgQjwojYAAAAAAAAACYkAAAAACYkRExkhRmEAAABYhpERIQAREyYJERRZEghABRAAAAAGSUMA8fPx9/H59fj49vnx+fn5+fn5+fnw8PD19Pnw9/lAQEBAQEDZyePFQMHJxEDw9fn09EBAQEBAQEBAQNfTxcHiwdXj1tVAQEDDweTiBUBAQEDyCEAIQFkBAFafMwPg+MiVBYCACIAAnzcEqM8vW58QBwYBEgOgsACfJgjyJ+vwin+AFZ82AgAXggIYAJwBAJ8aAghAmgMhEROfAgYAAAAACYlfKgIIQIQHoAAAAJgIQArw9vDw8Pn09fj4BHUAABAWwAAQAAAAAADFA4ExdpcGUyUgAAkAAAWAAAAAAg=="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_Sanjeev_70079(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAUBJNRAAAAABEABCjBAJAboMAAEANmZmgQjwpjYAAAAAAAAAAzgAAAAAAzgRBhQCSWEAAAABd0gmCREHVBEIQAUQAAACBkE3RvHz8fDx9PDx9/f0+dfw9fbw8ff29PT09fD58PX1+Pfw8EBA1+TC08nnQHv19vBAQEBAQEBAQEBAQEDV1tnD2dbi4kBAQEBAx8Hk4gVAQEBA8ghACEAU8PD04vj08MPw8PDw8PDw8PDy8fBZAQBWnzMD4PjIlQWAAAiAAJ83BCAzrSefEAcGARIDoLAAnyYIokJges9J8hufNgIADoICGACcAQCfGgIIQJoDIREGnwIGAAAAAFEAXyoCCECEB6AAAACYCEAK8fPw8PDz8PD58gZFAAAQAAEX0AAQAAAAAADFAwExBQVpR1XUIAADAAgFgAAAAAI="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_Bruce_50038(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAMtJNQgAAAACEABCjBAcAgJQAAEANmZkgQjwoDYAAAAAAAAAAQYAAAAAAQYQECNWCWEAAAACVwImEhARWRIIQJAQAAZERQDx8vjz8/Xw8vX38PPw9vHw8PD09/T09PX58PDw8PX28fBAQObUQOLk18XZw8XV48XZQHv19vHwQEBA4sHVQNnB1NbVQEBAQMPB5OIFQEBAQPIIQAhACvD28PDw+fT1+PMCRSAX0AAQAAAAAADFA4EoOGFpWUXVIAACAAAFgAAAAAI="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_Sanjeev_70081(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAUBJNRAAAAACEABCjBAWArnNAAEANmZmgQjwpjYAAAAAAAAAAhQAAAAAAhQRBiBIJGEAAAADEyImCREHVBEIQAUQAAAABkRFcfHz8fDy+PDz8fPy8/D29fPw8PDx9PT09fD58fDx8/Tw+UBA0tnWx8XZQHv29fNAQEBAQEBAQEBAQEDB09fIxdnF4+PBQEBAx8Hk4gVAQEBA8ghACEAU8PD04vj08MPw8PDw8PDw8PDy8fRZAQBWnzMD4PjIlQWAAAiAAJ83BKhCLb+fEAcGARIDoLAAnyYIa3ms5yFN/k2fNgIAFIICGACcAQCfGgIIQJoDIREGnwIGAAAAAAIUXyoCCECEB6AAAACYCEAK8fPw8PDz8PDy8gZFAAAQAAEX0AAQAAAAAADFA4ExB0kEckjUIAATAAQFgAAAAAI="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_Bruce_70092(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAUBJNSAAAAACEABCjBAQAlQ0AAEANmZmgQjwpjYAAAAAAAAABjEAAAAABjERCBhIOGEAAAADBYkmEhEJWRIIQAUQAAAABkRFAPHz8fLx9vDz8PX58PDw+fDw8PLy9PT09fHw8PH28Pnw80BA5sHTx9nFxdXiQHvx9vD58EBAQEBAQEDiwdVA2cHU1tVAQEBAw8Hk4gVAQEBA8ghACEAU8PD04vj08MPw8PDw8PDw8PD08fJZAQBWnzMD4EjIlQWAAAiAAJ83BHVKIfqfEAcGARIDoCAAnyYIUCx6r1fCVjSfNgIASoICGACcAQCfGgIIQJoDIREInwIGAAAAAAYxXyoCCECEB6AAAACYCEAK8Pbw8PD59PX48gZFAAAQAAEX0AAQAAAAAADFA4ExJncYAnfUIAAYAAUFgAAAAAI="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_70162(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAThJNQQAAAADEABCjBAUAwboAAEANn5mgQjwojYAAAAAAAAAEAAAAAAAEAARJCA5MmEAAAA5d1ASOQARJCYRESVUEQhABRAAAAAGQjFo8fPy+PLw8/n39/Xw+fn5+fn5+fnw8PD09fnz+flAQEBAQEDiwcbF5sHoQHvw+fH5QEBAQEBAQEBAQOLB2cHj1sfBQEBAQEDDweTiBUBAQEDyCEAIQGIBAF+fMwPg+MiVBYAACIAAnzcEhz0v/J8QBwYBEgOgoACfJgjmfk8hb3odnJ82AgAfggIYAJwBAJ8aAghAmgMhESSfAgYAAAAAEABfKgIIQJ8DBgAAAAAAAIQHoAAAAJgIQArw9vDw8Pn18PfwBnUAABAAARbAABAAAAAAAMUEYTKHQ3IBgSAAFgAABYAAAAAC"

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_80006(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECASpJNQgAAAABEABCjBAfAZGKAAEANmZmgQjwojYAAAAAAAAACZkAAAAACZkSAQIJImEAAAABWXEmERIBVBEIQAUQAAAABkE3RvHz8/X1+fDx9fn38vLy8/D48PDy9PT09fDw8Pby8vPw+EBA5sjW08XGxOJA4uPDQPHw8vb3QEBAQEDD5NfF2ePJ1dZAQEBAw8Hk4gVAQEBA8ghACEBZAQBWnzMD4PjIlQWAAAiAAJ83BBIp8W+fEAcGARIDoKAAnyYI8CoL4uN5lVmfNgIAJIICGACcAQCfGgIIQJoDIREwnwIGAAAAAAmZXyoCCECEB6AAAACYCEAK8Pbw8PD59fDx9AZFAAAQAAEWwAAQAAAAAADFAwEzUHdiF4UgAVUAJQWAAAAAAg=="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_80005(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAThJNRUAAAADEABCjBAeAwYwAAEANn5mgQjwojYAAAAAAAAACVIAAAAACVISAQFQQGEAAACJMhIXUAARMCYREgFUEQhABRAAAAAGQjFo8fPz9PDx+Pnz8vHy+fn5+fn5+fnw8PD09fn29/NAQEBAQEDiwcbF5sHoQHvz8vXxQEBAQEBAQEBAQMPk18XZ48nV1kBAQEDDweTiBUBAQEDyCEAIQGIBAF+fMwPg+MiVBYAACIAAnzcE3Eik7J8QBwYBEgOgoACfJgguqPSitunq1J82AgAjggIYAJwBAJ8aAghAmgMhETCfAgYAAAAACVJfKgIIQJ8DBgAAAAAAAIQHoAAAAJgIQArw9vDw8Pn18PH0BnUAABAAARbAABAAAAAAAMUEYTNQZkBgNiAAFgAABYAAAAAC"

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_80003(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAOlJNQMAAAADEABCjBAeAwZkAAEANn5kgQjwpDYAAAAAAAAACJkAAAAACJkSAQE4JWEAAAAxcpAXOAARMCYREgFZEghAkBAABklDAPHz8/Tw8fPx9/L58Pn5+fn5+fn58PDw9fT58Pn2QEBAQEBA2cnjxUDBycRA8PX59vdAQEBAQEBAQEDD5NfF2ePJ1dZAQEBAw8Hk4gVAQEBA8ghACEAU8PD04vj08MPw8PDw8PDw8PD19PUK8Pbw8PD59fDx9AZ1EAAAAAEX0AAQAAAAAADFBGEzUFkFUgTUIAAJAAAFgAAAAAI="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_80002(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAThJNQQAAAADEABCjBAdAwY2AAEANn5mgQjwojYAAAAAAAAAFTIAAAAAFTISAQEERGEAAACEZhAXBAARMCYREgFUEQhABRAAAAAGQjFo8fPz9PDx+PT29vHw+fn5+fn5+fnw8PD09fnz+flAQEBAQEDiwcbF5sHoQHvw+fH5QEBAQEBAQEBAQOLB2cHj1sfBQEBAQEDDweTiBUBAQEDyCEAIQGIBAF+fMwPg+MiVBYAACIAAnzcEYOY5Z58QBwYBEgOgoACfJgiCNLjI26aVF582AgAgggIYAJwBAJ8aAghAmgMhETCfAgYAAAAAFTJfKgIIQJ8DBgAAAAAAAIQHoAAAAJgIQArw9vDw8Pn18PfwBnUAABAAARbAABAAAAAAAMUEYTNQOISVgiAAFgAABYAAAAAC"

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_90003(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECARNJNQgAAAAGEABCjBAOBgsCAAEAtmZkgQjwpDYAAAAAAAAAJAAAAAAAAAAYUgAAAAAYUhIHF1JUYQAAAAFECCYREghZEghAARBZBkRFAPHz9PHx9PDx9PTw+fDy9Pbw8PHy9PT09fHw8PH18vT28kBA5sHTx9nFxdXiQHvx9fL09kBAQEBAQEDiwdXjwUDD08HZwUBAw8Hk4gtAQEBAQEBAQEBA1AhACEAU8PD04vj08MPw8PDw8PDw8PD39PAK8Pbw8PD59fD18AZJAAAABwEX0AAQAAAAAADmBYE0FkN0hzfUIAAYAAUFgAAAAAIRZgAOwAX59fD38M8F8vHy8fUOAEAAAAAAAADx8UD1+fU="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_90004(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAO9JNQQAAAAGEABCjBAPBllbAAQAtmZkgQzwoDYAAABAAAAAAAAAAAAAAAAYUgAAAAAYUhIHGAIGYQAAAAJZKCYREghZEghAARBZBkRFAPHz9PHx9PDx9PTw+fD0+PT49vDy9Pbw8PHy9PT09fHw8PH18vT28kBA5sHTx9nFxdXiQHvx9fL09kBAQEBAQEDiwdXjwUDD08HZwUBAw8Hk4gH3CEAIQArw9vDw8Pn18PXwBUkAAAAHF9AAEAAAAAAA5gWBNBZDdIc31CAAGAAFB6AAAAACJQEBAAFECBIHF1JUAAAERFAAAAAAAAA="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaISO_100017(t *testing.T) {
	fmt.Println("Welcome to zenda ISO test.")
	isoFromG := "FgECAUdJNQgAAAABEABCjBAfAYILAAEANmZmgQjwpjYAAAAAAAAAGAQAAAAAGAQBBQIYSGEAAAAAmEInAQEFWRIIQAcQAAAABkRFAPLw8PXx9vDw+fj08/Dw+fDw8Pfx9PT09fHw8PH28Pnw80BA5sHTx9nFxdXiQHvx9vD58EBAQEBAQEDiwdVA2cHU1tVAQEBAw8Hk4gVAQEBA8ghACEAU8PD04vj08MPw8PDw8PDw8PDw8PBgAQBdn24EIHAAAJ8zA+D4yJUFAAAAAACfNwQlQMTgnxAHBgESA6AAAJ8mCOrqv9ZSL3F2nzYCAAGCAgAAnAEAnxoCCECaAyIBBZ8CBgAAAAAYBF8qAghAhAegAAAAAxAQCvD28PDw+fT1+PIGRQAAEAABF9AAEAAAAAAAxQMCAFCDKANZ1CAAGAAFBYAAAAAC"

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Visa)
	err = message.Unpack(rawISO[22:])
	if err != nil {
		fmt.Println(err)
	}
	f62bytes, _ := message.GetField(62).Bytes()
	fmt.Println("Let's unpack Field 62 now!")
	f62 := field62.NewField62(specs.Spec87VisaField62)
	err = f62.Unpack(f62bytes)
	if err != nil {
		fmt.Println(err)
	}

	msdi, err := f62.GetString(4)
	if err != nil {
		fmt.Println("MSDI ErR: ", err.Error())
	}

	fmt.Println("MSDI Indicator: ", msdi)

	f54bytes, _ := message.GetField(54).Bytes()
	fmt.Println("Let's unpack Field 54 now!")
	f54 := field54.NewField54(specs.Spec87VisaField54)
	err = f54.Unpack(f54bytes)
	if err != nil {
		fmt.Println(err)
	}

	qhp, ok := f54.Amounts["4S"]
	if !ok {
		fmt.Println("QHP Not Found")
	}
	fmt.Println("QHP Amount is: ", qhp)
}

func TestZendaMCISO_601(t *testing.T) {
	fmt.Println("Welcome to Zenda MC ISO test.")
	isoFromG := "8PLw8L77RgGI4eYKAAAAAAAAABDw8PDw8PDw8PDw8PDw8PDy8PLw8PDw8PDw8PDy8PLw8PDw8PDw8PDy8PLx8vD58PTx8vH39vHw8PDw8PD28fDw8PDw8Pbw+fT48vLw8fLw8PHy8Pjx8vD58fLw+PX58fLw9fHw8PDw+fj08fL29/fy9/Hw+fDw8PDw8PPw8fHz9PL19vHw+fT48vL09fbx8PDx8PDw8PDw8PDw8vT19vHw5sHTYNTB2eNAe/X28fBAQEBAQEBAQEDiwdVA2cHU1tVAQEBAQEDDwfDx8Nn28fD18PDx8PD49PD49PD49PDw8vDw8PHw+PTwxPDw8PDw8PDw8PLw8vHz8Z8DBgAAAAAAAJ8mCDge3WS/gVxcggIYAJ82AgBSnzQDQgAAnwIGAAAAAAZ2nycBgIQHoAAAAJgIQJ8QBwYBEgOgoACfCQIAjZ8zA+DYyJ8aAghAnx4IU0MwMTAwNjaaAyESCJ81ASKVBYCABIAAXyoCCECfQQMDkICcAQCfNwS31eLu8PL28PDw8PDx8PDw8PXw8fj08Pn09fjzQEBAQEDw8fLU4vD2+PLy8/Py+fPw8vHx9/Xw8PHw8PDw8PDw8PDy9PX28fA="

	rawISO, err := base64.StdEncoding.DecodeString(isoFromG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ISO Message from Galileo is: ", rawISO)
	// fmt.Println("Raw ISO without header is: ", rawISO[22:])

	message := iso8583.NewMessage(specs.Spec87Maestro)
	err = message.Unpack(rawISO)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Let's unpack DE 48")
	f48bytes, _ := message.GetField(48).Bytes()
	fmt.Println("Let's unpack Field 48 now!")
	f48ebcdic, err := encoding.EBCDIC.Encode(f48bytes)
	if err != nil {
		fmt.Println("ECBDIC Encode Err: ", err)
	}
	f48 := field48.NewField48((*field48.Field48Spec)(specs.Spec87MCField48))
	err = f48.Unpack(f48ebcdic)
	if err != nil {
		fmt.Println(err)
	}
	f48_61_str, _ := f48.GetString(61)
	msdi, _ := strconv.Atoi(string(f48_61_str[2]))
	fmt.Println("f48.GetString(61):", msdi)
	fmt.Println("Let's unpack Field 54 now!")
	f54bytes, _ := message.GetField(54).Bytes()
	f54ebcdic, err := encoding.EBCDIC.Encode(f54bytes)
	if err != nil {
		fmt.Println("ECBDIC Encode Err: ", err)
	}

	f54 := mcf54.NewField54(specs.Spec87MCField54)
	err = f54.Unpack(f54ebcdic)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("f54.Amounts: ", f54.Amounts["10"])
}

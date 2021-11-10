package iso8583_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/moov-io/iso8583"
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

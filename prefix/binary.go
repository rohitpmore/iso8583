package prefix

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

type digitType int

const (
	Dtype_BCD digitType = iota + 1
	Dtype_Binary
)

type binaryVarPrefixer struct {
	Digits int
	DType  digitType
}

var Binary = Prefixers{
	Fixed: &binaryFixedPrefixer{},
	L:     &binaryVarPrefixer{1, Dtype_Binary},
	LL:    &binaryVarPrefixer{2, Dtype_Binary},
	LLL:   &binaryVarPrefixer{3, Dtype_Binary},
	LLLL:  &binaryVarPrefixer{4, Dtype_Binary},
}

func (p *binaryVarPrefixer) EncodeLength(maxLen, dataLen int) ([]byte, error) {
	if dataLen > maxLen {
		return nil, fmt.Errorf("field length: %d is larger than maximum: %d", dataLen, maxLen)
	}

	// if len(strconv.Itoa(dataLen)) > p.Digits {
	// 	return nil, fmt.Errorf("number of digits in length: %d exceeds: %d", dataLen, p.Digits)
	// }

	res := new(bytes.Buffer)
	err := binary.Write(res, binary.BigEndian, uint32(dataLen))
	if err != nil {
		return nil, err
	}

	return res.Bytes()[len(res.Bytes())-p.Digits:], nil
}

func (p *binaryVarPrefixer) DecodeLength(maxLen int, data []byte) (int, int, error) {
	if len(data) < p.Digits {
		return 0, 0, fmt.Errorf("length mismatch: want to read %d bytes, get only %d", p.Digits, len(data))
	}

	bLen := data[:p.Digits]

	for len(bLen) != 4 {
		bLen = append([]byte{0b0}, bLen...)
	}
	buf := bytes.NewBuffer(bLen)
	var res int32
	binary.Read(buf, binary.BigEndian, &res)

	if int(res) > maxLen {
		return 0, 0, fmt.Errorf("data length %d is larger than maximum %d", int(res), maxLen)
	}

	return int(res), p.Digits, nil
}

func (p *binaryVarPrefixer) Inspect() string {
	return fmt.Sprintf("Binary.%s", strings.Repeat("L", p.Digits))
}

type binaryFixedPrefixer struct {
}

func (p *binaryFixedPrefixer) EncodeLength(fixLen, dataLen int) ([]byte, error) {
	if dataLen != fixLen {
		return nil, fmt.Errorf("field length: %d should be fixed: %d", dataLen, fixLen)
	}

	return []byte{}, nil
}

func (p *binaryFixedPrefixer) DecodeLength(fixLen int, data []byte) (int, int, error) {
	return fixLen, 0, nil
}

func (p *binaryFixedPrefixer) Inspect() string {
	return "Binary.Fixed"
}

package bencodius

import (
	"fmt"
	"strconv"
	"strings"
)

type BencodeValue interface {
	isBencoded()
}

type BencodeInt int

func (value BencodeInt) isBencoded() {}

type BencodeString string

func (value BencodeString) isBencoded() {}

type BencodeList []BencodeValue

func (values BencodeList) isBencoded() {}

type BencodeDict struct {
	keys []BencodeString
	dict map[BencodeString]BencodeValue
}

func (o *BencodeDict) insert(key BencodeString, value BencodeValue) {
	o.keys = append(o.keys, key)
	o.dict[key] = value
}

func (o *BencodeDict) get(key string) BencodeValue {
	return o.dict[BencodeString(key)]
}

func (o *BencodeDict) wtf() string {
	return "wtf"
}

func (values BencodeDict) isBencoded() {}

func Decode(data string) BencodeValue {
	var value, _ = decodeValue(data, 0)
	return value
}

func decodeValue(data string, index int) (result BencodeValue, next int) {
	if data[index] == 'i' {
		result, next = decodeInt(data, index+1)
	} else if data[index] == 'l' {
		result, next = decodeList(data, index+1)
	} else if data[index] == 'd' {
		result, next = decodeDict(data, index+1)
	} else {
		result, next = decodeString(data, index) //no terminal beginning to clear
	}
	return
}

// str -> int -> (BInt, int)
// def decode-int data, index
// 	next = data.index(after: index, 'e')
//     ? to-int(data.slice(index, to: next), next + 1)

func decodeInt(data string, index int) (BencodeInt, int) {
	// decode numbere -> number
	// 4e -> 4 22e -> 22

	var i int
	for i = index; i < len(data); i++ {
		if data[i] == 'e' {
			break
		}
	}
	var result, _ = strconv.Atoi(data[index:i])
	return BencodeInt(result), i + 1
}

func decodeString(data string, index int) (BencodeString, int) {
	// decode number:value -> "value"
	// 4:edfa -> "edfa" 4:moni -> "moni"

	var i int
	for i = index; i < len(data); i++ {
		if data[i] == ':' {
			break
		}
	}
	var string_start = i + 1
	var string_length, _ = strconv.Atoi(data[index:i])
	return BencodeString(data[string_start : string_start+string_length]), string_start + string_length
}

func decodeList(data string, index int) (BencodeList, int) {
	// decode lvaluese -> [values]
	// l2:eee -> ["ee"] li2ee -> [2]

	var (
		new_index int
		values    = []BencodeValue{}
		next      BencodeValue
	)
	for new_index = index; new_index < len(data); {
		if data[new_index] == 'e' {
			break
		}
		next, new_index = decodeValue(data, new_index)
		values = append(values, next)
	}
	return BencodeList(values), new_index + 1
}

func decodeDict(data string, index int) (BencodeDict, int) {
	// decode dkeyvaluee -> {key : value}
	// d2:eei2ee -> {"ee" : 2} d4:moni1:2e -> {"moni" : "2"}

	var (
		next          int
		internal_dict = make(map[BencodeString]BencodeValue)
		dict          = BencodeDict{keys: make([]BencodeString, 0, 1024), dict: internal_dict}
		key           BencodeString
		value         BencodeValue
	)

	for next = index; next < len(data); {
		if data[next] == 'e' {
			break
		}
		key, next = decodeString(data, next)
		value, next = decodeValue(data, next)
		dict.insert(key, value)
	}
	return dict, next + 1
}

func Encode(data BencodeValue) string {
	var out string
	switch v := data.(type) {
	case BencodeInt:
		out = fmt.Sprintf("i%de", v)
	case BencodeString:
		out = fmt.Sprintf("%d:%s", len(v), v)
	case BencodeList:
		result := []string{}
		for _, a := range v {
			result = append(result, Encode(a))
		}
		out = fmt.Sprintf("l%se", strings.Join(result, ""))
	case BencodeDict:
		result := []string{}
		for _, y := range v.keys {
			z := v.dict[y]
			result = append(append(result, Encode(y)), Encode(z))
		}
		out = fmt.Sprintf("d%se", strings.Join(result, ""))
	}
	return out
}

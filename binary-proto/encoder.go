package binary_proto

import (
	"framework-go/utils/bytes"
	"reflect"
)

/**
 * @Author: imuge
 * @Date: 2020/5/21 下午
 */

// 编码头信息
func encodeHeader(obj interface{}) []byte {
	contract := obj.(DataContract)
	buf := bytes.Int32ToBytes(contract.Code())
	buf = append(buf, bytes.Int64ToBytes(contract.Version())...)
	return buf
}

func encodePrimitiveType(v reflect.Value, primitiveType string, numberMask bytes.NumberMask) []byte {
	switch primitiveType {
	case PRIMITIVETYPE_NIL:
		return []byte{}
	case PRIMITIVETYPE_BOOLEAN:
		return []byte{encodeBool(v.Bool())}
	case PRIMITIVETYPE_INT8:
		if numberMask.Equals(bytes.NUMBERMASK_NONE) {
			return []byte{encodeInt8(int8(v.Int()))}
		} else {
			return encodeNumberMask(v.Int(), numberMask)
		}
	case PRIMITIVETYPE_INT16:
		if numberMask.Equals(bytes.NUMBERMASK_NONE) {
			return encodeInt16(int16(v.Int()))
		} else {
			return encodeNumberMask(v.Int(), numberMask)
		}
	case PRIMITIVETYPE_INT32:
		if numberMask.Equals(bytes.NUMBERMASK_NONE) {
			return encodeInt32(int32(v.Int()))
		} else {
			return encodeNumberMask(v.Int(), numberMask)
		}
	case PRIMITIVETYPE_INT64:
		if numberMask.Equals(bytes.NUMBERMASK_NONE) {
			return encodeInt64(v.Int())
		} else {
			return encodeNumberMask(v.Int(), numberMask)
		}
	case PRIMITIVETYPE_TEXT:
		return encodeString(v.String())
	default:
		panic("un support primitive type")
	}
}

func encodeBytes(data []byte) []byte {
	buf := encodeSize(len(data))
	return append(buf, data...)
}

func encodeBool(data bool) byte {
	return bytes.BoolToBytes(data)
}

func encodeInt8(data int8) byte {
	return bytes.Int8ToBytes(data)
}

func encodeInt16(data int16) []byte {
	return bytes.Int16ToBytes(data)
}

func encodeInt32(data int32) []byte {
	return bytes.Int32ToBytes(data)
}

func encodeInt64(data int64) []byte {
	return bytes.Int64ToBytes(data)
}

func encodeNumberMask(data int64, mask bytes.NumberMask) []byte {
	return mask.WriteMask(data)
}

func encodeSize(data int) []byte {
	return bytes.NUMBERMASK_NORMAL.WriteMask(int64(data))
}

func encodeString(data string) []byte {
	return encodeBytes(bytes.StringToBytes(data))
}

func encodeEnum(c *Codec, value int64, refEnum int) []byte {
	contract := (c.enumMap[int32(refEnum)]).(EnumContract)
	switch contract.Type() {
	case PRIMITIVETYPE_INT8:
		return []byte{encodeInt8(int8(value))}
	case PRIMITIVETYPE_INT16:
		return encodeInt16(int16(value))
	case PRIMITIVETYPE_INT32:
		return encodeInt32(int32(value))
	default:
		panic("un support enum value type, int8,int16,int32 only")
	}
}

func encodeArrayHeader(count int) []byte {
	return bytes.NUMBERMASK_NORMAL.WriteMask(int64(count))
}

func encodeGeneric(c *Codec, v interface{}) []byte {
	// 与非泛型引用无差别
	return encodeContract(c, v)
}

func encodeContract(c *Codec, v interface{}) []byte {
	buf, err := c.Encode(v.(DataContract))
	if err != nil {
		panic(err)
	}
	buf = append(encodeSize(len(buf)), buf...)
	return buf
}

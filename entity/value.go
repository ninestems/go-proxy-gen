package entity

func define(in string) ValueType {
	switch in {
	case "bool":
		return ValueTypeBool
	case "int":
		return ValueTypeInt
	case "int8":
		return ValueTypeInt8
	case "int16":
		return ValueTypeInt16
	case "int32":
		return ValueTypeInt32
	case "int64":
		return ValueTypeInt64
	case "uint":
		return ValueTypeUint
	case "uint8":
		return ValueTypeUint8
	case "uint16":
		return ValueTypeUint16
	case "uint32":
		return ValueTypeUint32
	case "uint64":
		return ValueTypeUint64
	case "uintptr":
		return ValueTypeUintptr
	case "float32":
		return ValueTypeFloat32
	case "float64":
		return ValueTypeFloat64
	case "complex64":
		return ValueTypeComplex64
	case "complex128":
		return ValueTypeComplex128
	case "string":
		return ValueTypeString
	case "byte":
		return ValueTypeByte
	case "rune":
		return ValueTypeRune
	default:
		return ValueTypeStruct
	}
}

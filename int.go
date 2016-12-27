package cast

import (
	"strconv"
)

// Inter is the interface that wraps the Int method.
// Int method return int64
type Inter interface {
	Int() int64
}

//Int cast input to int64
func Int(input interface{}) (output int64, err error) {

	switch castValue := input.(type) {
	case Inter:
		output = castValue.Int()
		return
	case string:
		output, err = strconv.ParseInt(castValue, 10, 64)
		return
	case []byte:
		output, err = strconv.ParseInt(string(castValue), 10, 64)
		return
	case int:
		output = int64(castValue)
		return
	case int8:
		output = int64(castValue)
		return
	case int16:
		output = int64(castValue)
		return
	case int32:
		output = int64(castValue)
		return
	case int64:
		output = int64(castValue)
		return
	case uint:
		output = int64(castValue)
		return
	case uint8:
		output = int64(castValue)
		return
	case uint16:
		output = int64(castValue)
		return
	case uint32:
		output = int64(castValue)
		return
	case uint64:
		output = int64(castValue)
		return
	case float32:
		output = int64(castValue)
		return
	case float64:
		output = int64(castValue)
		return
	case bool:
		output = int64(0)
		if castValue {
			output = int64(1)
		}
		return
	case nil:
		output = int64(0)
		return
	default:
		err = NewCastError("Could not convert to int")
	}
	return
}

//MustInt cast input to int64 and panic if error
func MustInt(input interface{}) int64 {
	output, err := Int(input)
	if err != nil {
		panic(err)
	}
	return output
}

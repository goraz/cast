package cast

import (
	"strconv"
)

// Uinter is the interface that wraps the Uint method.
// Uint method return uint64
type Uinter interface {
	Uint() uint64
}

//Uint Cast input to uint64
func Uint(input interface{}) (output uint64, err error) {

	switch castValue := input.(type) {
	case Uinter:
		output = castValue.Uint()
		return
	case string:
		output, err = strconv.ParseUint(castValue, 10, 64)
		return
	case []byte:
		output, err = strconv.ParseUint(string(castValue), 10, 64)
		return
	case int:
		if castValue < 0 {
			err = NewCastError("Could not convert negative value to uint")
			return
		}
		output = uint64(castValue)
		return
	case int8:
		if castValue < 0 {
			err = NewCastError("Could not convert negative value to uint")
			return
		}
		output = uint64(castValue)
		return
	case int16:
		if castValue < 0 {
			err = NewCastError("Could not convert negative value to uint")
			return
		}
		output = uint64(castValue)
		return
	case int32:
		if castValue < 0 {
			err = NewCastError("Could not convert negative value to uint")
			return
		}
		output = uint64(castValue)
		return
	case int64:
		if castValue < 0 {
			err = NewCastError("Could not convert negative value to uint")
			return
		}
		output = uint64(castValue)
		return
	case uint:
		output = uint64(castValue)
		return
	case uint8:
		output = uint64(castValue)
		return
	case uint16:
		output = uint64(castValue)
		return
	case uint32:
		output = uint64(castValue)
		return
	case uint64:
		output = uint64(castValue)
		return
	case float32:
		if castValue < 0 {
			err = NewCastError("Could not convert negative value to uint")
			return
		}
		output = uint64(castValue)
		return
	case float64:
		if castValue < 0 {
			err = NewCastError("Could not convert negative value to uint")
			return
		}
		output = uint64(castValue)
		return
	case bool:
		output = uint64(0)
		if castValue {
			output = uint64(1)
		}
		return
	case nil:
		output = uint64(0)
		return
	default:
		err = NewCastError("Could not convert to uint")
	}
	return
}

//MustUint cast input to uint64 and panic if error
func MustUint(input interface{}) uint64 {
	output, err := Uint(input)
	if err != nil {
		panic(err)
	}
	return output
}

package cast

import (
	"strconv"
)

// Booler is the interface that wraps the Bool method.
// Bool method return bool
type Booler interface {
	Bool() bool
}

//Bool cast input to bool
func Bool(input interface{}) (output bool, err error) {
	switch castValue := input.(type) {
	case Booler:
		output = castValue.Bool()
		return
	case string:
		var errParse error
		output, errParse = strconv.ParseBool(castValue)
		if errParse != nil && len(castValue) > 0 {
			output = true
		}
		return
	case int:
		if castValue != 0 {
			output = true
		}
		return
	case int8:
		if castValue != 0 {
			output = true
		}
		return
	case int16:
		if castValue != 0 {
			output = true
		}
		return
	case int32:
		if castValue != 0 {
			output = true
		}
		return
	case int64:
		if castValue != 0 {
			output = true
		}
		return
	case uint:
		if castValue != 0 {
			output = true
		}
		return
	case uint8:
		if castValue != 0 {
			output = true
		}
		return
	case uint16:
		if castValue != 0 {
			output = true
		}
		return
	case uint32:
		if castValue != 0 {
			output = true
		}
		return
	case uint64:
		if castValue != 0 {
			output = true
		}
		return
	case float32:
		if castValue != 0 {
			output = true
		}
		return
	case float64:
		if castValue != 0 {
			output = true
		}
		return
	case bool:
		output = castValue
		return
	case nil:
		output = false
		return
	default:
		err = NewCastError("Could not convert to bool")
	}
	return
}

//MustBool cast input to bool and panic if error
func MustBool(input interface{}) bool {
	output, err := Bool(input)
	if err != nil {
		panic(err)
	}
	return output
}

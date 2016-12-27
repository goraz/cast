package cast

import (
	"strconv"
)

// Floater is the interface that wraps the Float method.
// Float method return int64
type Floater interface {
	Float() float64
}

//Float cast input to int64
func Float(input interface{}) (output float64, err error) {

	switch castValue := input.(type) {
	case Floater:
		output = castValue.Float()
		return
	case string:
		output, err = strconv.ParseFloat(castValue, 64)
	case []byte:
		output, err = strconv.ParseFloat(string(castValue), 64)
		return
	case int:
		output = float64(castValue)
		return
	case int8:
		output = float64(castValue)
		return
	case int16:
		output = float64(castValue)
		return
	case int32:
		output = float64(castValue)
		return
	case int64:
		output = float64(castValue)
		return
	case uint:
		output = float64(castValue)
		return
	case uint8:
		output = float64(castValue)
		return
	case uint16:
		output = float64(castValue)
		return
	case uint32:
		output = float64(castValue)
		return
	case uint64:
		output = float64(castValue)
		return
	case float32:
		output = float64(castValue)
		return
	case float64:
		output = float64(castValue)
		return
	case bool:
		output = float64(0)
		if castValue {
			output = float64(1)
		}
		return
	case nil:
		output = float64(0)
		return
	default:
		err = NewCastError("Could not convert to float64")
	}
	return
}

//MustFloat cast input to int64 and panic if error
func MustFloat(input interface{}) float64 {
	output, err := Float(input)
	if err != nil {
		panic(err)
	}
	return output
}

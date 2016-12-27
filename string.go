package cast

import (
	"strconv"
)

// Stringer is the interface that wraps the String method.
// String method return string
type Stringer interface {
	String() string
}

//String cast input to string
func String(input interface{}) (output string, err error) {

	switch castValue := input.(type) {
	case Stringer:
		output = castValue.String()
		return
	case string:
		output = string(castValue)
		return
	case []byte:
		output = string(castValue)
		return
	case int:
		output = strconv.Itoa(castValue)
		return
	case int8:
		output = strconv.Itoa(int(castValue))
		return
	case int16:
		output = strconv.Itoa(int(castValue))
		return
	case int32:
		output = strconv.Itoa(int(castValue))
		return
	case int64:
		output = strconv.FormatInt(castValue, 10)
		return
	case uint:
		output = strconv.FormatUint(uint64(castValue), 10)
		return
	case uint8:
		output = strconv.FormatUint(uint64(castValue), 10)
		return
	case uint16:
		output = strconv.FormatUint(uint64(castValue), 10)
		return
	case uint32:
		output = strconv.FormatUint(uint64(castValue), 10)
		return
	case uint64:
		output = strconv.FormatUint(uint64(castValue), 10)
		return
	case float32:
		// What exact value for perc value
		output = strconv.FormatFloat(float64(castValue), 'g', -1, 64)
		return
	case float64:
		// What exact value for perc value
		output = strconv.FormatFloat(castValue, 'g', -1, 64)
		return
	case nil:
		output = ""
		return
	case bool:
		output = strconv.FormatBool(castValue)
		return
	default:
		err = NewCastError("Could not convert to string")
	}
	return
}

//MustString cast input to string and panic if error
func MustString(input interface{}) string {
	output, err := String(input)
	if err != nil {
		panic(err)
	}
	return output
}

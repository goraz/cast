package cast

// StringSlicer is the interface that wraps the StringSlice method.
// StringSlicer method return []string
type StringSlicer interface {
	StringSlice() []string
}

//StringSlice cast slice input to string slice
func StringSlice(input interface{}) (output []string, err error) {
	var castError error
	switch castValue := input.(type) {
	case StringSlicer:
		output = castValue.StringSlice()
		return
	case []string:
		output = castValue
		return
	case []interface{}:
		output = make([]string, len(castValue))
		for index := range castValue {
			if output[index], castError = String(castValue[index]); castError != nil {
				//todo. return better error
				err = NewCastError("Could not convert to string")
			}
		}
		return
	case []int:
		output = make([]string, len(castValue))
		for index := range castValue {
			output[index], _ = String(castValue[index])
		}
		return
	case []int8:
		output = make([]string, len(castValue))
		for index := range castValue {
			output[index], _ = String(castValue[index])
		}
		return
	case []int16:
		output = make([]string, len(castValue))
		for index := range castValue {
			output[index], _ = String(castValue[index])
		}
		return
	case []int32:
		output = make([]string, len(castValue))
		for index := range castValue {
			output[index], _ = String(castValue[index])
		}
		return
	case []int64:
		output = make([]string, len(castValue))
		for index := range castValue {
			output[index], _ = String(castValue[index])
		}
		return
	case []uint:
		output = make([]string, len(castValue))
		for index := range castValue {
			output[index], _ = String(castValue[index])
		}
		return
	case []uint8:
		output = make([]string, len(castValue))
		for index := range castValue {
			output[index], _ = String(castValue[index])
		}
		return
	case []uint16:
		output = make([]string, len(castValue))
		for index := range castValue {
			output[index], _ = String(castValue[index])
		}
		return
	case []uint32:
		output = make([]string, len(castValue))
		for index := range castValue {
			output[index], _ = String(castValue[index])
		}
		return
	case []uint64:
		output = make([]string, len(castValue))
		for index := range castValue {
			output[index], _ = String(castValue[index])
		}
		return
	case []float32:
		output = make([]string, len(castValue))
		for index := range castValue {
			output[index], _ = String(castValue[index])
		}
		return
	case []float64:
		output = make([]string, len(castValue))
		for index := range castValue {
			output[index], _ = String(castValue[index])
		}
		return
	case []bool:
		output = make([]string, len(castValue))
		for index := range castValue {
			output[index], _ = String(castValue[index])
		}
		return

	default:
		err = NewCastError("Could not convert to string")
	}
	return
}

//MustStringSlice cast input slice to string slice and panic if error
func MustStringSlice(input interface{}) []string {
	output, err := StringSlice(input)
	if err != nil {
		panic(err)
	}
	return output
}

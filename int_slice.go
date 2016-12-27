package cast

// IntSlicer is the interface that wraps the IntSlice method.
// IntSlice method return []int64
type IntSlicer interface {
	IntSlice() []int64
}

//IntSlice cast slice input to int64 slice
func IntSlice(input interface{}) (output []int64, err error) {
	var castError error
	switch castValue := input.(type) {
	case IntSlicer:
		output = castValue.IntSlice()
		return
	case []string:
		output = make([]int64, len(castValue))
		for index := range castValue {
			if output[index], castError = Int(castValue[index]); castError != nil {
				//todo. return better error
				err = NewCastError("Could not convert to int64")
			}

		}
		return
	case []interface{}:
		output = make([]int64, len(castValue))
		for index := range castValue {
			if output[index], castError = Int(castValue[index]); castError != nil {
				//todo. return better error
				err = NewCastError("Could not convert to int64")
			}
		}
		return
	case []int:
		output = make([]int64, len(castValue))
		for index := range castValue {
			output[index], _ = Int(castValue[index])
		}
		return
	case []int8:
		output = make([]int64, len(castValue))
		for index := range castValue {
			output[index], _ = Int(castValue[index])
		}
		return
	case []int16:
		output = make([]int64, len(castValue))
		for index := range castValue {
			output[index], _ = Int(castValue[index])
		}
		return
	case []int32:
		output = make([]int64, len(castValue))
		for index := range castValue {
			output[index], _ = Int(castValue[index])
		}
		return
	case []int64:
		output = castValue
		return
	case []uint:
		output = make([]int64, len(castValue))
		for index := range castValue {
			output[index], _ = Int(castValue[index])
		}
		return
	case []uint8:
		output = make([]int64, len(castValue))
		for index := range castValue {
			output[index], _ = Int(castValue[index])
		}
		return
	case []uint16:
		output = make([]int64, len(castValue))
		for index := range castValue {
			output[index], _ = Int(castValue[index])
		}
		return
	case []uint32:
		output = make([]int64, len(castValue))
		for index := range castValue {
			output[index], _ = Int(castValue[index])
		}
		return
	case []uint64:
		output = make([]int64, len(castValue))
		for index := range castValue {
			output[index], _ = Int(castValue[index])
		}
		return
	case []float32:
		output = make([]int64, len(castValue))
		for index := range castValue {
			output[index], _ = Int(castValue[index])
		}
		return
	case []float64:
		output = make([]int64, len(castValue))
		for index := range castValue {
			output[index], _ = Int(castValue[index])
		}
		return

	case []bool:
		output = make([]int64, len(castValue))
		for index := range castValue {
			output[index], _ = Int(castValue[index])
		}
		return

	default:
		err = NewCastError("Could not convert to string")
	}
	return
}

//MustIntSlice cast input slice to int64 slice and panic if error
func MustIntSlice(input interface{}) []int64 {
	output, err := IntSlice(input)
	if err != nil {
		panic(err)
	}
	return output
}

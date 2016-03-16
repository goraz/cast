package cast

// FloatSlicer is the interface that wraps the FloatSlice method.
// FloatSlice method return []float64
type FloatSlicer interface {
	FloatSlice() []float64
}

//FloatSlice cast slice input to float64 slice
func FloatSlice(input interface{}) (output []float64, err error) {
	var castError error
	switch castValue := input.(type) {
	case FloatSlicer:
		output = castValue.FloatSlice()
		return
	case []string:
		output = make([]float64, len(castValue))
		for index := range castValue {
			if output[index], castError = Float(castValue[index]); castError != nil {
				//todo. return better error
				err = NewCastError("Could not convert to string")
			}
		}
		return
	case []interface{}:
		output = make([]float64, len(castValue))
		for index := range castValue {
			if output[index], castError = Float(castValue[index]); castError != nil {
				//todo. return better error
				err = NewCastError("Could not convert to string")
			}
		}
		return
	case []int:
		output = make([]float64, len(castValue))
		for index := range castValue {
			output[index], _ = Float(castValue[index])
		}
		return
	case []int8:
		output = make([]float64, len(castValue))
		for index := range castValue {
			output[index], _ = Float(castValue[index])
		}
		return
	case []int16:
		output = make([]float64, len(castValue))
		for index := range castValue {
			output[index], _ = Float(castValue[index])
		}
		return
	case []int32:
		output = make([]float64, len(castValue))
		for index := range castValue {
			output[index], _ = Float(castValue[index])
		}
		return
	case []int64:
		output = make([]float64, len(castValue))
		for index := range castValue {
			output[index], _ = Float(castValue[index])
		}
		return
	case []uint:
		output = make([]float64, len(castValue))
		for index := range castValue {
			output[index], _ = Float(castValue[index])
		}
		return
	case []uint8:
		output = make([]float64, len(castValue))
		for index := range castValue {
			output[index], _ = Float(castValue[index])
		}
		return
	case []uint16:
		output = make([]float64, len(castValue))
		for index := range castValue {
			output[index], _ = Float(castValue[index])
		}
		return
	case []uint32:
		output = make([]float64, len(castValue))
		for index := range castValue {
			output[index], _ = Float(castValue[index])
		}
		return
	case []uint64:
		output = make([]float64, len(castValue))
		for index := range castValue {
			output[index], _ = Float(castValue[index])
		}
		return
	case []float32:
		output = make([]float64, len(castValue))
		for index := range castValue {
			output[index], _ = Float(castValue[index])
		}
		return
	case []float64:
		output = castValue
		return
	case []bool:
		output = make([]float64, len(castValue))
		for index := range castValue {
			output[index], _ = Float(castValue[index])
		}
		return
	default:
		err = NewCastError("Could not convert to string")
	}
	return
}

//MustFloatSlice cast input slice to float64 slice and panic if error
func MustFloatSlice(input interface{}) []float64 {
	output, err := FloatSlice(input)
	if err != nil {
		panic(err)
	}
	return output
}

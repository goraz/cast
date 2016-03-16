package cast

import "time"

//Duration cast input to time.Duration
func Duration(input interface{}) (output time.Duration, err error) {
	switch castValue := input.(type) {
	case string:
		output, err = time.ParseDuration(castValue)
		if err != nil {
			outputInt, errCast := Int(castValue)
			if errCast != nil {
				err = NewCastError("Could not convert to time.Duration")
				return
			}
			output = time.Duration(int(outputInt))
		}
		return
	case time.Duration:
		output = castValue
		return
	default:
		outputInt, errCast := Int(castValue)
		if errCast == nil {
			output = time.Duration(int(outputInt))
			return
		}
		err = NewCastError("Could not convert to time.Duration")
	}
	return
}

//MustDuration cast input to time.Duration and panic if error
func MustDuration(input interface{}) time.Duration {
	output, err := Duration(input)
	if err != nil {
		panic(err)
	}
	return output
}

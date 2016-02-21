package cast

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type BoolCaster struct{}

func (c BoolCaster) Bool() bool {
	return true
}

type NotBoolCaster struct{}

func TestBoolCast(t *testing.T) {
	Convey("Cast int to bool", t, func() {
		val := 12
		castVal, _ := Bool(val)
		So(castVal, ShouldEqual, true)

		zeroVal := 0
		castZeroVal, _ := Bool(zeroVal)
		So(castZeroVal, ShouldEqual, false)

		negativeVal := -12
		castNegativeVal, _ := Bool(negativeVal)
		So(castNegativeVal, ShouldEqual, true)
	})

	Convey("Cast int8 to bool", t, func() {
		val := int8(12)
		castVal, _ := Bool(val)
		So(castVal, ShouldEqual, true)

		zeroVal := int8(0)
		castZeroVal, _ := Bool(zeroVal)
		So(castZeroVal, ShouldEqual, false)

		negativeVal := int8(-12)
		castNegativeVal, _ := Bool(negativeVal)
		So(castNegativeVal, ShouldEqual, true)
	})

	Convey("Cast int16 to bool", t, func() {
		val := int16(12)
		castVal, _ := Bool(val)
		So(castVal, ShouldEqual, true)

		zeroVal := int16(0)
		castZeroVal, _ := Bool(zeroVal)
		So(castZeroVal, ShouldEqual, false)

		negativeVal := int16(-12)
		castNegativeVal, _ := Bool(negativeVal)
		So(castNegativeVal, ShouldEqual, true)
	})

	Convey("Cast int32 to bool", t, func() {
		val := int32(12)
		castVal, _ := Bool(val)
		So(castVal, ShouldEqual, true)

		zeroVal := int32(0)
		castZeroVal, _ := Bool(zeroVal)
		So(castZeroVal, ShouldEqual, false)

		negativeVal := int32(-12)
		castNegativeVal, _ := Bool(negativeVal)
		So(castNegativeVal, ShouldEqual, true)
	})

	Convey("Cast int64 to bool", t, func() {
		val := int64(12)
		castVal, _ := Bool(val)
		So(castVal, ShouldEqual, true)

		zeroVal := int64(0)
		castZeroVal, _ := Bool(zeroVal)
		So(castZeroVal, ShouldEqual, false)

		negativeVal := int64(-12)
		castNegativeVal, _ := Bool(negativeVal)
		So(castNegativeVal, ShouldEqual, true)
	})

	Convey("Cast uint to bool", t, func() {
		val := uint(12)
		castVal, _ := Bool(val)
		So(castVal, ShouldEqual, true)

		zeroVal := uint(0)
		castZeroVal, _ := Bool(zeroVal)
		So(castZeroVal, ShouldEqual, false)
	})

	Convey("Cast uint8 to bool", t, func() {
		val := uint8(12)
		castVal, _ := Bool(val)
		So(castVal, ShouldEqual, true)

		zeroVal := uint8(0)
		castZeroVal, _ := Bool(zeroVal)
		So(castZeroVal, ShouldEqual, false)
	})

	Convey("Cast uint16 to bool", t, func() {
		val := uint16(12)
		castVal, _ := Bool(val)
		So(castVal, ShouldEqual, true)

		zeroVal := uint16(0)
		castZeroVal, _ := Bool(zeroVal)
		So(castZeroVal, ShouldEqual, false)
	})

	Convey("Cast uint32 to bool", t, func() {
		val := uint32(12)
		castVal, _ := Bool(val)
		So(castVal, ShouldEqual, true)

		zeroVal := uint32(0)
		castZeroVal, _ := Bool(zeroVal)
		So(castZeroVal, ShouldEqual, false)
	})

	Convey("Cast uint64 to bool", t, func() {
		val := uint64(12)
		castVal, _ := Bool(val)
		So(castVal, ShouldEqual, true)

		zeroVal := uint64(0)
		castZeroVal, _ := Bool(zeroVal)
		So(castZeroVal, ShouldEqual, false)
	})

	Convey("Cast float32 to bool", t, func() {
		val := float32(1.2)
		castVal, _ := Bool(val)
		So(castVal, ShouldEqual, true)

		zeroVal := float32(0)
		castZeroVal, _ := Bool(zeroVal)
		So(castZeroVal, ShouldEqual, false)

		negativeVal := float32(-1.2)
		castNegativeVal, _ := Bool(negativeVal)
		So(castNegativeVal, ShouldEqual, true)
	})

	Convey("Cast float64 to bool", t, func() {
		val := float64(1.2)
		castVal, _ := Bool(val)
		So(castVal, ShouldEqual, true)

		zeroVal := float64(0)
		castZeroVal, _ := Bool(zeroVal)
		So(castZeroVal, ShouldEqual, false)

		negativeVal := float64(-1.2)
		castNegativeVal, _ := Bool(negativeVal)
		So(castNegativeVal, ShouldEqual, true)
	})

	Convey("Cast bool to bool", t, func() {
		val := true
		castVal, _ := Bool(val)
		So(castVal, ShouldEqual, true)

		falseVal := false
		castFalseVal, _ := Bool(falseVal)
		So(castFalseVal, ShouldEqual, false)
	})

	Convey("Cast nil to bool", t, func() {
		castVal, _ := Bool(nil)
		So(castVal, ShouldEqual, false)
	})

	Convey("Cast string to bool", t, func() {
		val := "true"
		castVal, _ := Bool(val)
		So(castVal, ShouldEqual, true)

		val = "false"
		castVal, _ = Bool(val)
		So(castVal, ShouldEqual, false)

		val = "abc"
		castVal, _ = Bool(val)
		So(castVal, ShouldEqual, true)
		val = ""
		castVal, _ = Bool(val)
		So(castVal, ShouldEqual, false)
	})

	Convey("Cast NotBoolCaster to bool and return Error", t, func() {
		val := NotBoolCaster{}
		castVal, err := Bool(val)
		So(castVal, ShouldEqual, false)
		So(err, ShouldNotEqual, nil)
	})

	Convey("Cast BoolCaster to bool", t, func() {
		val := BoolCaster{}
		castVal, err := Bool(val)
		So(castVal, ShouldEqual, true)
		So(err, ShouldEqual, nil)
	})

	Convey("Must cast without panic", t, func() {

		So(func() {
			val := BoolCaster{}
			MustBool(val)
		}, ShouldNotPanic)
	})

	Convey("Must cast with panic", t, func() {

		So(func() {
			val := NotBoolCaster{}
			MustBool(val)
		}, ShouldPanic)
	})
}

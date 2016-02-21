package cast

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type FloatCaster struct{}

func (c FloatCaster) Float() float64 {
	return 1.2
}

type NotFloatCaster struct{}

func TestFloatCast(t *testing.T) {
	Convey("Cast int to float", t, func() {
		val := 12
		castVal, _ := Float(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := 0
		castZeroVal, _ := Float(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := -12
		castNegativeVal, _ := Float(negativeVal)
		So(castNegativeVal, ShouldEqual, -12)
	})

	Convey("Cast int8 to float", t, func() {
		val := int8(12)
		castVal, _ := Float(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := int8(0)
		castZeroVal, _ := Float(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := int8(-12)
		castNegativeVal, _ := Float(negativeVal)
		So(castNegativeVal, ShouldEqual, -12)
	})

	Convey("Cast int16 to float", t, func() {
		val := int16(12)
		castVal, _ := Float(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := int16(0)
		castZeroVal, _ := Float(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := int16(-12)
		castNegativeVal, _ := Float(negativeVal)
		So(castNegativeVal, ShouldEqual, -12)
	})

	Convey("Cast int32 to float", t, func() {
		val := int32(12)
		castVal, _ := Float(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := int32(0)
		castZeroVal, _ := Float(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := int32(-12)
		castNegativeVal, _ := Float(negativeVal)
		So(castNegativeVal, ShouldEqual, -12)
	})

	Convey("Cast int64 to float", t, func() {
		val := int64(12)
		castVal, _ := Float(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := int64(0)
		castZeroVal, _ := Float(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := int64(-12)
		castNegativeVal, _ := Float(negativeVal)
		So(castNegativeVal, ShouldEqual, -12)
	})

	Convey("Cast uint to float", t, func() {
		val := uint(12)
		castVal, _ := Float(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := uint(0)
		castZeroVal, _ := Float(zeroVal)
		So(castZeroVal, ShouldEqual, 0)
	})

	Convey("Cast uint8 to float", t, func() {
		val := uint8(12)
		castVal, _ := Float(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := uint8(0)
		castZeroVal, _ := Float(zeroVal)
		So(castZeroVal, ShouldEqual, 0)
	})

	Convey("Cast uint16 to float", t, func() {
		val := uint16(12)
		castVal, _ := Float(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := uint16(0)
		castZeroVal, _ := Float(zeroVal)
		So(castZeroVal, ShouldEqual, 0)
	})

	Convey("Cast uint32 to float", t, func() {
		val := uint32(12)
		castVal, _ := Float(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := uint32(0)
		castZeroVal, _ := Float(zeroVal)
		So(castZeroVal, ShouldEqual, 0)
	})

	Convey("Cast uint64 to float", t, func() {
		val := uint64(12)
		castVal, _ := Float(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := uint64(0)
		castZeroVal, _ := Float(zeroVal)
		So(castZeroVal, ShouldEqual, 0)
	})

	Convey("Cast float32 to float", t, func() {
		val := float32(1.2)
		castVal, _ := Float(val)
		So(castVal, ShouldEqual, 1.2000000476837158)

		zeroVal := float32(0)
		castZeroVal, _ := Float(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := float32(-1.2)
		castNegativeVal, _ := Float(negativeVal)
		So(castNegativeVal, ShouldEqual, -1.2000000476837158)
	})

	Convey("Cast float64 to float", t, func() {
		val := float64(1.2)
		castVal, _ := Float(val)
		So(castVal, ShouldEqual, 1.2)

		zeroVal := float64(0)
		castZeroVal, _ := Float(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := float64(-1.2)
		castNegativeVal, _ := Float(negativeVal)
		So(castNegativeVal, ShouldEqual, -1.2)
	})

	Convey("Cast bool to float", t, func() {
		val := true
		castVal, _ := Float(val)
		So(castVal, ShouldEqual, 1)

		falseVal := false
		castFalseVal, _ := Float(falseVal)
		So(castFalseVal, ShouldEqual, 0)
	})

	Convey("Cast nil to float", t, func() {
		castVal, _ := Float(nil)
		So(castVal, ShouldEqual, 0)
	})

	Convey("Cast string to float", t, func() {
		val := "12"
		castVal, _ := Float(val)
		So(castVal, ShouldEqual, 12)

		val = "0"
		castVal, _ = Float(val)
		So(castVal, ShouldEqual, 0)

		val = "-12"
		castVal, _ = Float(val)
		So(castVal, ShouldEqual, -12)

		val = "12.246"
		castVal, _ = Float(val)
		So(castVal, ShouldEqual, 12.246)
	})

	Convey("Cast NotFloatCaster to float and return Error", t, func() {
		val := NotFloatCaster{}
		castVal, err := Float(val)
		So(castVal, ShouldEqual, 0)
		So(err, ShouldNotEqual, nil)
	})

	Convey("Cast FloatCaster to float", t, func() {
		val := FloatCaster{}
		castVal, err := Float(val)
		So(castVal, ShouldEqual, 1.2)
		So(err, ShouldEqual, nil)
	})

	Convey("Must cast without panic", t, func() {

		So(func() {
			val := FloatCaster{}
			MustFloat(val)
		}, ShouldNotPanic)
	})

	Convey("Must cast with panic", t, func() {

		So(func() {
			val := NotFloatCaster{}
			MustFloat(val)
		}, ShouldPanic)
	})
}

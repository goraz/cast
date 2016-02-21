package cast

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type IntCaster struct{}

func (c IntCaster) Int() int64 {
	return 12
}

type NotIntCaster struct{}

func TestIntCast(t *testing.T) {
	Convey("Cast int to int", t, func() {
		val := 12
		castVal, _ := Int(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := 0
		castZeroVal, _ := Int(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := -12
		castNegativeVal, _ := Int(negativeVal)
		So(castNegativeVal, ShouldEqual, -12)
	})

	Convey("Cast int8 to int", t, func() {
		val := int8(12)
		castVal, _ := Int(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := int8(0)
		castZeroVal, _ := Int(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := int8(-12)
		castNegativeVal, _ := Int(negativeVal)
		So(castNegativeVal, ShouldEqual, -12)
	})

	Convey("Cast int16 to int", t, func() {
		val := int16(12)
		castVal, _ := Int(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := int16(0)
		castZeroVal, _ := Int(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := int16(-12)
		castNegativeVal, _ := Int(negativeVal)
		So(castNegativeVal, ShouldEqual, -12)
	})

	Convey("Cast int32 to int", t, func() {
		val := int32(12)
		castVal, _ := Int(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := int32(0)
		castZeroVal, _ := Int(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := int32(-12)
		castNegativeVal, _ := Int(negativeVal)
		So(castNegativeVal, ShouldEqual, -12)
	})

	Convey("Cast int64 to int", t, func() {
		val := int64(12)
		castVal, _ := Int(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := int64(0)
		castZeroVal, _ := Int(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := int64(-12)
		castNegativeVal, _ := Int(negativeVal)
		So(castNegativeVal, ShouldEqual, -12)
	})

	Convey("Cast uint to int", t, func() {
		val := uint(12)
		castVal, _ := Int(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := uint(0)
		castZeroVal, _ := Int(zeroVal)
		So(castZeroVal, ShouldEqual, 0)
	})

	Convey("Cast uint8 to int", t, func() {
		val := uint8(12)
		castVal, _ := Int(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := uint8(0)
		castZeroVal, _ := Int(zeroVal)
		So(castZeroVal, ShouldEqual, 0)
	})

	Convey("Cast uint16 to int", t, func() {
		val := uint16(12)
		castVal, _ := Int(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := uint16(0)
		castZeroVal, _ := Int(zeroVal)
		So(castZeroVal, ShouldEqual, 0)
	})

	Convey("Cast uint32 to int", t, func() {
		val := uint32(12)
		castVal, _ := Int(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := uint32(0)
		castZeroVal, _ := Int(zeroVal)
		So(castZeroVal, ShouldEqual, 0)
	})

	Convey("Cast uint64 to int", t, func() {
		val := uint64(12)
		castVal, _ := Int(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := uint64(0)
		castZeroVal, _ := Int(zeroVal)
		So(castZeroVal, ShouldEqual, 0)
	})

	Convey("Cast float32 to int", t, func() {
		val := float32(1.2)
		castVal, _ := Int(val)
		So(castVal, ShouldEqual, 1)

		zeroVal := float32(0)
		castZeroVal, _ := Int(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := float32(-1.2)
		castNegativeVal, _ := Int(negativeVal)
		So(castNegativeVal, ShouldEqual, -1)
	})

	Convey("Cast float64 to int", t, func() {
		val := float64(1.2)
		castVal, _ := Int(val)
		So(castVal, ShouldEqual, 1)

		zeroVal := float64(0)
		castZeroVal, _ := Int(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := float64(-1.2)
		castNegativeVal, _ := Int(negativeVal)
		So(castNegativeVal, ShouldEqual, -1)
	})

	Convey("Cast bool to int", t, func() {
		val := true
		castVal, _ := Int(val)
		So(castVal, ShouldEqual, 1)

		falseVal := false
		castFalseVal, _ := Int(falseVal)
		So(castFalseVal, ShouldEqual, 0)
	})

	Convey("Cast nil to int", t, func() {
		castVal, _ := Int(nil)
		So(castVal, ShouldEqual, 0)
	})

	Convey("Cast string to int", t, func() {
		val := "12"
		castVal, _ := Int(val)
		So(castVal, ShouldEqual, 12)

		val = "0"
		castVal, _ = Int(val)
		So(castVal, ShouldEqual, 0)

		val = "-12"
		castVal, _ = Int(val)
		So(castVal, ShouldEqual, -12)

		val = "12.246"
		castVal, _ = Int(val)
		So(castVal, ShouldEqual, 0)
	})

	Convey("Cast NotIntCaster to int and return Error", t, func() {
		val := NotIntCaster{}
		castVal, err := Int(val)
		So(castVal, ShouldEqual, 0)
		So(err, ShouldNotEqual, nil)
	})

	Convey("Cast IntCaster to int", t, func() {
		val := IntCaster{}
		castVal, err := Int(val)
		So(castVal, ShouldEqual, 12)
		So(err, ShouldEqual, nil)
	})

	Convey("Must cast without panic", t, func() {

		So(func() {
			val := IntCaster{}
			MustInt(val)
		}, ShouldNotPanic)
	})

	Convey("Must cast with panic", t, func() {

		So(func() {
			val := NotIntCaster{}
			MustInt(val)
		}, ShouldPanic)
	})
}

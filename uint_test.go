package cast

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type UintCaster struct{}

func (c UintCaster) Uint() uint64 {
	return 12
}

type NotUintCaster struct{}

func TestUintCast(t *testing.T) {
	Convey("Cast int to uint", t, func() {
		val := 12
		castVal, _ := Uint(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := 0
		castZeroVal, _ := Uint(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := -12
		castNegativeVal, err := Uint(negativeVal)
		So(castNegativeVal, ShouldEqual, 0)
		So(err, ShouldNotEqual, nil)
	})

	Convey("Cast int8 to uint", t, func() {
		val := int8(12)
		castVal, _ := Uint(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := int8(0)
		castZeroVal, _ := Uint(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := int8(-12)
		castNegativeVal, err := Uint(negativeVal)
		So(castNegativeVal, ShouldEqual, 0)
		So(err, ShouldNotEqual, nil)
	})

	Convey("Cast int16 to uint", t, func() {
		val := int16(12)
		castVal, _ := Uint(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := int16(0)
		castZeroVal, _ := Uint(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := int16(-12)
		castNegativeVal, err := Uint(negativeVal)
		So(castNegativeVal, ShouldEqual, 0)
		So(err, ShouldNotEqual, nil)
	})

	Convey("Cast int32 to uint", t, func() {
		val := int32(12)
		castVal, _ := Uint(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := int32(0)
		castZeroVal, _ := Uint(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := int32(-12)
		castNegativeVal, err := Uint(negativeVal)
		So(castNegativeVal, ShouldEqual, 0)
		So(err, ShouldNotEqual, nil)
	})

	Convey("Cast int64 to uint", t, func() {
		val := int64(12)
		castVal, _ := Uint(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := int64(0)
		castZeroVal, _ := Uint(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := int64(-12)
		castNegativeVal, err := Uint(negativeVal)
		So(castNegativeVal, ShouldEqual, 0)
		So(err, ShouldNotEqual, nil)
	})

	Convey("Cast uint to uint", t, func() {
		val := uint(12)
		castVal, _ := Uint(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := uint(0)
		castZeroVal, _ := Uint(zeroVal)
		So(castZeroVal, ShouldEqual, 0)
	})

	Convey("Cast uint8 to uint", t, func() {
		val := uint8(12)
		castVal, _ := Uint(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := uint8(0)
		castZeroVal, _ := Uint(zeroVal)
		So(castZeroVal, ShouldEqual, 0)
	})

	Convey("Cast uint16 to uint", t, func() {
		val := uint16(12)
		castVal, _ := Uint(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := uint16(0)
		castZeroVal, _ := Uint(zeroVal)
		So(castZeroVal, ShouldEqual, 0)
	})

	Convey("Cast uint32 to uint", t, func() {
		val := uint32(12)
		castVal, _ := Uint(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := uint32(0)
		castZeroVal, _ := Uint(zeroVal)
		So(castZeroVal, ShouldEqual, 0)
	})

	Convey("Cast uint64 to uint", t, func() {
		val := uint64(12)
		castVal, _ := Uint(val)
		So(castVal, ShouldEqual, 12)

		zeroVal := uint64(0)
		castZeroVal, _ := Uint(zeroVal)
		So(castZeroVal, ShouldEqual, 0)
	})

	Convey("Cast float32 to uint", t, func() {
		val := float32(1.2)
		castVal, _ := Uint(val)
		So(castVal, ShouldEqual, 1)

		zeroVal := float32(0)
		castZeroVal, _ := Uint(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := float32(-1.2)
		castNegativeVal, err := Uint(negativeVal)
		So(castNegativeVal, ShouldEqual, 0)
		So(err, ShouldNotEqual, nil)
	})

	Convey("Cast float64 to uint", t, func() {
		val := float64(1.2)
		castVal, _ := Uint(val)
		So(castVal, ShouldEqual, 1)

		zeroVal := float64(0)
		castZeroVal, _ := Uint(zeroVal)
		So(castZeroVal, ShouldEqual, 0)

		negativeVal := float64(-1.2)
		castNegativeVal, err := Uint(negativeVal)
		So(castNegativeVal, ShouldEqual, 0)
		So(err, ShouldNotEqual, nil)
	})

	Convey("Cast bool to uint", t, func() {
		val := true
		castVal, _ := Uint(val)
		So(castVal, ShouldEqual, 1)

		falseVal := false
		castFalseVal, _ := Uint(falseVal)
		So(castFalseVal, ShouldEqual, 0)
	})

	Convey("Cast nil to uint", t, func() {
		castVal, _ := Uint(nil)
		So(castVal, ShouldEqual, 0)
	})

	Convey("Cast string to uint", t, func() {
		val := "12"
		castVal, _ := Uint(val)
		So(castVal, ShouldEqual, 12)

		val = "0"
		castVal, _ = Uint(val)
		So(castVal, ShouldEqual, 0)

		val = "-12"
		castVal, _ = Uint(val)
		So(castVal, ShouldEqual, 0)

		val = "12.246"
		castVal, _ = Uint(val)
		So(castVal, ShouldEqual, 0)
	})

	Convey("Cast NotUintCaster to uint and return Error", t, func() {
		val := NotUintCaster{}
		castVal, err := Uint(val)
		So(castVal, ShouldEqual, 0)
		So(err, ShouldNotEqual, nil)
	})

	Convey("Cast UintCaster to uint", t, func() {
		val := UintCaster{}
		castVal, err := Uint(val)
		So(castVal, ShouldEqual, 12)
		So(err, ShouldEqual, nil)
	})

	Convey("Must cast without panic", t, func() {

		So(func() {
			val := UintCaster{}
			MustUint(val)
		}, ShouldNotPanic)
	})

	Convey("Must cast with panic", t, func() {

		So(func() {
			val := NotUintCaster{}
			MustUint(val)
		}, ShouldPanic)
	})
}

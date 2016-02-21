package cast

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type StringCaster struct{}

func (c StringCaster) String() string {
	return "abc"
}

type NotStringCaster struct{}

func TestStringCast(t *testing.T) {
	Convey("Cast int to string", t, func() {
		val := 12
		castVal, _ := String(val)
		So(castVal, ShouldEqual, "12")

		zeroVal := 0
		castZeroVal, _ := String(zeroVal)
		So(castZeroVal, ShouldEqual, "0")

		negativeVal := -12
		castNegativeVal, _ := String(negativeVal)
		So(castNegativeVal, ShouldEqual, "-12")
	})

	Convey("Cast int8 to string", t, func() {
		val := int8(12)
		castVal, _ := String(val)
		So(castVal, ShouldEqual, "12")

		zeroVal := int8(0)
		castZeroVal, _ := String(zeroVal)
		So(castZeroVal, ShouldEqual, "0")

		negativeVal := int8(-12)
		castNegativeVal, _ := String(negativeVal)
		So(castNegativeVal, ShouldEqual, "-12")
	})

	Convey("Cast int16 to string", t, func() {
		val := int16(12)
		castVal, _ := String(val)
		So(castVal, ShouldEqual, "12")

		zeroVal := int16(0)
		castZeroVal, _ := String(zeroVal)
		So(castZeroVal, ShouldEqual, "0")

		negativeVal := int16(-12)
		castNegativeVal, _ := String(negativeVal)
		So(castNegativeVal, ShouldEqual, "-12")
	})

	Convey("Cast int32 to string", t, func() {
		val := int32(12)
		castVal, _ := String(val)
		So(castVal, ShouldEqual, "12")

		zeroVal := int32(0)
		castZeroVal, _ := String(zeroVal)
		So(castZeroVal, ShouldEqual, "0")

		negativeVal := int32(-12)
		castNegativeVal, _ := String(negativeVal)
		So(castNegativeVal, ShouldEqual, "-12")
	})

	Convey("Cast int64 to string", t, func() {
		val := int64(12)
		castVal, _ := String(val)
		So(castVal, ShouldEqual, "12")

		zeroVal := int64(0)
		castZeroVal, _ := String(zeroVal)
		So(castZeroVal, ShouldEqual, "0")

		negativeVal := int64(-12)
		castNegativeVal, _ := String(negativeVal)
		So(castNegativeVal, ShouldEqual, "-12")
	})

	Convey("Cast uint to string", t, func() {
		val := uint(12)
		castVal, _ := String(val)
		So(castVal, ShouldEqual, "12")

		zeroVal := uint(0)
		castZeroVal, _ := String(zeroVal)
		So(castZeroVal, ShouldEqual, "0")
	})

	Convey("Cast uint8 to string", t, func() {
		val := uint8(12)
		castVal, _ := String(val)
		So(castVal, ShouldEqual, "12")

		zeroVal := uint8(0)
		castZeroVal, _ := String(zeroVal)
		So(castZeroVal, ShouldEqual, "0")
	})

	Convey("Cast uint16 to string", t, func() {
		val := uint16(12)
		castVal, _ := String(val)
		So(castVal, ShouldEqual, "12")

		zeroVal := uint16(0)
		castZeroVal, _ := String(zeroVal)
		So(castZeroVal, ShouldEqual, "0")
	})

	Convey("Cast uint32 to string", t, func() {
		val := uint32(12)
		castVal, _ := String(val)
		So(castVal, ShouldEqual, "12")

		zeroVal := uint32(0)
		castZeroVal, _ := String(zeroVal)
		So(castZeroVal, ShouldEqual, "0")
	})

	Convey("Cast uint64 to string", t, func() {
		val := uint64(12)
		castVal, _ := String(val)
		So(castVal, ShouldEqual, "12")

		zeroVal := uint64(0)
		castZeroVal, _ := String(zeroVal)
		So(castZeroVal, ShouldEqual, "0")
	})

	Convey("Cast float32 to string", t, func() {
		val := float32(1.2)
		castVal, _ := String(val)
		//:))
		So(castVal, ShouldEqual, "1.2000000476837158")

		zeroVal := float32(0)
		castZeroVal, _ := String(zeroVal)
		So(castZeroVal, ShouldEqual, "0")

		negativeVal := float32(-1.2)
		castNegativeVal, _ := String(negativeVal)
		So(castNegativeVal, ShouldEqual, "-1.2000000476837158")
	})

	Convey("Cast float64 to string", t, func() {
		val := float64(1.2)
		castVal, _ := String(val)
		So(castVal, ShouldEqual, "1.2")

		zeroVal := float64(0)
		castZeroVal, _ := String(zeroVal)
		So(castZeroVal, ShouldEqual, "0")

		negativeVal := float64(-1.2)
		castNegativeVal, _ := String(negativeVal)
		So(castNegativeVal, ShouldEqual, "-1.2")
	})

	Convey("Cast bool to string", t, func() {
		val := true
		castVal, _ := String(val)
		So(castVal, ShouldEqual, "true")

		falseVal := false
		castFalseVal, _ := String(falseVal)
		So(castFalseVal, ShouldEqual, "false")
	})

	Convey("Cast nil to string", t, func() {
		castVal, _ := String(nil)
		So(castVal, ShouldEqual, "")
	})

	Convey("Cast string to string", t, func() {

		val := "abc"
		castVal, _ := String(val)
		So(castVal, ShouldEqual, "abc")
	})

	Convey("Cast NotStringCaster to string and return Error", t, func() {
		val := NotStringCaster{}
		castVal, err := String(val)
		So(castVal, ShouldEqual, "")
		So(err, ShouldNotEqual, nil)
	})

	Convey("Cast StringCaster to string", t, func() {
		val := StringCaster{}
		castVal, err := String(val)
		So(castVal, ShouldEqual, "abc")
		So(err, ShouldEqual, nil)
	})

	Convey("Must cast without panic", t, func() {

		So(func() {
			val := StringCaster{}
			MustString(val)
		}, ShouldNotPanic)
	})

	Convey("Must cast with panic", t, func() {

		So(func() {
			val := NotStringCaster{}
			MustString(val)
		}, ShouldPanic)
	})
}

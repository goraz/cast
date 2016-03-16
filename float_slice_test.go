package cast

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

type FloatStringCaster struct{}

func (c FloatStringCaster) FloatSlice() []float64 {
	return []float64{1, 2, 3}
}

type NotFloatStringCaster struct{}

func TestFloatSliceCast(t *testing.T) {
	Convey("Cast []int to []float6", t, func() {
		val := []int{1, 2, 3}
		castVal, _ := FloatSlice(val)

		So(reflect.DeepEqual(castVal, []float64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []int8 to []float6", t, func() {
		val := []int8{1, 2, 3}
		castVal, _ := FloatSlice(val)

		So(reflect.DeepEqual(castVal, []float64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []int16 to []float6", t, func() {
		val := []int16{1, 2, 3}
		castVal, _ := FloatSlice(val)

		So(reflect.DeepEqual(castVal, []float64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []int32 to []float6", t, func() {
		val := []int32{1, 2, 3}
		castVal, _ := FloatSlice(val)

		So(reflect.DeepEqual(castVal, []float64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []int64 to []float6", t, func() {
		val := []int64{1, 2, 3}
		castVal, _ := FloatSlice(val)

		So(reflect.DeepEqual(castVal, []float64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []uint to []float6", t, func() {
		val := []uint{1, 2, 3}
		castVal, _ := FloatSlice(val)

		So(reflect.DeepEqual(castVal, []float64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []uint8 to []float6", t, func() {
		val := []uint8{1, 2, 3}
		castVal, _ := FloatSlice(val)

		So(reflect.DeepEqual(castVal, []float64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []uint16 to []float6", t, func() {
		val := []uint16{1, 2, 3}
		castVal, _ := FloatSlice(val)

		So(reflect.DeepEqual(castVal, []float64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []uint32 to []float6", t, func() {
		val := []uint32{1, 2, 3}
		castVal, _ := FloatSlice(val)

		So(reflect.DeepEqual(castVal, []float64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []uint64 to []float6", t, func() {
		val := []uint64{1, 2, 3}
		castVal, _ := FloatSlice(val)

		So(reflect.DeepEqual(castVal, []float64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast []float32 to []float6", t, func() {
		val := []float32{1.1, 2.1, 3.1}
		castVal, _ := FloatSlice(val)
		So(reflect.DeepEqual(castVal, []float64{1.100000023841858, 2.0999999046325684, 3.0999999046325684}), ShouldBeTrue)
	})

	Convey("Cast []float64 to []float6", t, func() {
		val := []float64{1.1, 2.1, 3.1}
		castVal, _ := FloatSlice(val)
		So(reflect.DeepEqual(castVal, []float64{1.1, 2.1, 3.1}), ShouldBeTrue)
	})

	Convey("Cast []bool to []float6", t, func() {
		val := []bool{true, false, true}
		castVal, _ := FloatSlice(val)
		So(reflect.DeepEqual(castVal, []float64{1, 0, 1}), ShouldBeTrue)
	})

	Convey("Cast []string to []float6", t, func() {

		val := []string{"1", "2", "3"}
		castVal, _ := FloatSlice(val)

		So(reflect.DeepEqual(castVal, []float64{1, 2, 3}), ShouldBeTrue)
	})

	Convey("Cast invalid []string to []float6", t, func() {

		val := []string{"1", "a", "3"}
		castVal, err := FloatSlice(val)

		So(reflect.DeepEqual(castVal, []float64{1, 0, 3}), ShouldBeTrue)
		So(err, ShouldNotEqual, nil)
	})

	Convey("Cast []NotFloatStringCaster to []float6 and return Error", t, func() {
		val := NotFloatStringCaster{}
		castVal, err := FloatSlice(val)
		So(len(castVal), ShouldEqual, 0)
		So(err, ShouldNotEqual, nil)
	})

	Convey("Cast []FloatStringCaster to []float6", t, func() {
		val := FloatStringCaster{}
		castVal, err := FloatSlice(val)
		So(reflect.DeepEqual(castVal, []float64{1, 2, 3}), ShouldBeTrue)
		So(err, ShouldEqual, nil)
	})

	Convey("Cast []interface{} to []float6", t, func() {
		val := []interface{}{"1", 2, true}
		castVal, err := FloatSlice(val)
		So(reflect.DeepEqual(castVal, []float64{1, 2, 1}), ShouldBeTrue)
		So(err, ShouldEqual, nil)
	})

	Convey("Cast []interface{} to []float6 with error", t, func() {
		val := []interface{}{"1", 1, true, NotFloatStringCaster{}, uint8(2)}
		castVal, err := FloatSlice(val)
		So(reflect.DeepEqual(castVal, []float64{1, 1, 1, 0, 2}), ShouldBeTrue)
		So(err, ShouldNotEqual, nil)
	})

	Convey("Must cast without panic", t, func() {

		So(func() {
			val := FloatStringCaster{}
			MustFloatSlice(val)
		}, ShouldNotPanic)
	})

	Convey("Must cast with panic", t, func() {

		So(func() {
			val := NotFloatStringCaster{}
			MustFloatSlice(val)
		}, ShouldPanic)
	})
}
